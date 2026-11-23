# Repository Guidelines

## Project Overview
`github.com/Ptt-official-app/pttbbs-backend`: Go middleware layer between web/app clients and legacy PTT BBS. Gin HTTP/JSON API plus an embedded OIDC provider. Talks to `go-pttbbs` (HTTP), `boardd`/`mand` (gRPC), MongoDB (cache/store), Redis (OIDC/email state), and Lemmy. Go 1.26.x. Repo was renamed from `go-openbbsmiddleware`.

## Architecture & Data Flow
Client -> Gin (CORS, binder/`process*`) -> `XWrapper(c)` -> `X(...)` handler -> one of:
- `schema.*` (MongoDB documents, `*_c` collections)
- `boardd` / `mand` gRPC clients (global `Cli`)
- `utils.BackendGet/BackendPost` (HTTP to go-pttbbs, uses `pttbbsapi.*` param/result structs)
- `lemmyutils` (e.g. `TryCreateBoard`)

Result -> `processResult`: 200 returns JSON, 303 redirects, otherwise `{Msg, tokenuser}` JSON.
- Background: `cron/Retry*` loops refresh Mongo from boardd/mand; `queue` (golang-queue, single attempt, no retries) processes comments asynchronously (`QueueCommentDBCS` -> `ProcessCommentQueue`). Queue tasks must not call `Start`/`Close`.
- Auth: JWT access token + `X-CSRFToken` header vs cookie (skipped for `CLIENT_TYPE_APP` and `types.DEV`). Over-18 via cookie.
- `main.go`: `initGin()` selects `initGinAllGuest()` or `initGinCore()` by `types.IS_ALL_GUEST`. Many routes are intentionally commented out (read-mostly guest build). Shutdown via `appleboy/graceful` (cron loops as jobs, `queue.Close()` after HTTP stops).
- Startup: `main_init.go` `initAllConfig` (viper ini -> each package `InitConfig()`: types, db, schema, pttbbsapi, pttbbstypes, queue, boardd, mand, dbcs) then `initMain` (`schema.Init`, `boardd.Init(false)`, `mand.Init(false)`).

## Key Directories
- `api/` one file per endpoint (+ helpers: `api.go` binders, `api_login_required*.go`, `api_utils.go` JWT/CSRF, `errors.go`, `*_perm.go`)
- `schema/` Mongo document structs, indexes (`init.go`), Redis prefixes (`const.go`), OIDC client impl
- `apitypes/` response DTOs with `New*` constructors from schema structs
- `db/` thin Mongo/Redis wrappers; `boardd/`, `mand/` gRPC clients (generated `*.pb.go`, do not hand-edit)
- `queue/`, `cron/`, `dbcs/` (Big5/DBCS article+comment parsing), `fav/`, `oidcop/`, `zk/`, `lemmyutils/`, `utils/`, `configutil/`, `types/` (global config vars, Big5 tables)
- `mock/`, `mockhttp/`, `testcase/` test fixtures; `docs/` (PIPs, nginx, config templates); `scripts/`; `apidoc/` (Python swagger gen); `docker/`
- `pttbbs/`, `lemmy/`, `lemmy-js-client/`, `lemmy-ui/` are git submodules: do not edit. `mandloadmanarticles/`, `boarddreadcontent/` are extra binaries.

## Development Commands
```
make / make build     # bin/pttbbs-backend (-tags $(TAGS))
make test TAGS=noqueue   # TAGS 預設為空，不帶會讓 -tags 吃掉 -coverprofile；go test -p 1 -v -cover -tags $(TAGS) -coverprofile coverage.txt ./...
make lint             # golangci-lint run -v --timeout=3m
make fmt              # golangci-lint fmt (gofmt + goimports)
go test -p 1 -v ./... -run TestQueueCommentDBCS -tags queue   # queue test, separate
./pttbbs-backend -ini docs/examples/etc/pttbbs-backend/production.ini
./scripts/start-getting-started.sh   # docker compose dev stack
./scripts/swagger.sh <host>          # openapi via apidoc/ + swagger-ui
scripts/protoc_board.sh | protoc_man.sh   # regenerate gRPC stubs
```
Version injected via `-ldflags -X types.VERSION / types.GIT_VERSION`. Production build uses `-tags production`.

## Code Conventions & Common Patterns
- **Handler pattern** (copy `api/get_board_detail.go`): `const X_R = "/path/:id"`; types `XParams` (tags `json`/`form`/`uri`), optional `XPath`, `XResult`; `func XWrapper(c *gin.Context)` calls a binder, e.g. `LoginRequiredPathQuery(X, params, path, c)`; `func X(remoteAddr string, user *UserInfo, params any, [path any,] c *gin.Context) (result any, statusCode int, err error)` type-asserts params, returning `(nil, 400, ErrInvalidParams)` on mismatch. Return explicit status: 400 / 403 / 500 / 200. `XAllGuest` variants set `user.UserID = GUEST` and delegate. Register route in `main.go` as `router.GET(withPrefix(api.X_R), api.XWrapper)` (both init functions if needed).
- **Naming**: snake_case files; ALL_CAPS for constants/config vars/routes (revive underscore rules are disabled); `Board_c` collections; bson field consts `BOARD_BBOARD_ID_b`; `_db` suffix = schema-side values, `_b` = backend values; `theX` for asserted params.
- **Errors**: sentinel `errors.New` in per-package `errors.go` (`ErrXxx`); log with logrus `"Func: msg: e: %v"`; `%w` wrapping in queue.
- **Config**: per-package `config.go` (`setXConfig("KEY", KEY)`), `config_util.go` (`configPrefix`, e.g. `pttbbs-backend:types`), defaults as package-level vars. INI sections `[pttbbs-backend:<pkg>]`; `[go-pttbbs:*]` sections must match go-pttbbs. Keys ALL_CAPS. `SetDurationConfig` takes nanoseconds.
- **State/DI**: package-level globals (`boardd.Cli`, `schema.client`, `queue.client`), mutex-guarded; no DI framework. Cron jobs: `func(ctx context.Context) error` looping on `select ctx.Done()`.
- Production code branches on `SERVICE_MODE == TEST` / `IsTest` to use mocks.

## Important Files
`main.go`, `main_init.go`, `api/api.go`, `api/types.go`, `schema/init.go`, `types/00-config.go`, `configutil/config_util.go`, `utils/http.go`, `.golangci.yaml`, `docs/config/*.ini` (config templates), `testcase/test.ini`, `docker/Dockerfile`.

## Runtime/Tooling Preferences
- Go per `go.mod` (1.26.x); module deps vendored by `go mod`, no vendor dir. `go.mod` has `ignore docs`.
- Needs MongoDB (replica set, 27017) and Redis (6379) for tests and run.
- Lint: golangci-lint v2 (gosec, misspell US, nolintlint, revive); tests are not linted; `testinit.go` and `00-config.go` excluded. README suggests gofumpt and gotests.
- `gotest` (rakyll) is used by `scripts/test.sh` only.

## Testing & QA
- testing + testify (`assert`/`require`); `testutil.TDeepEqual` from go-pttbbs; gotests-style table tests (`TestXxx`, `Test_xxx`), `<source>_test.go` beside source. No gomock/sqlmock.
- Each package has `testinit.go` (`SetIsTest/UnsetIsTest`, mutex-guarded) and `testinit_test.go` (`setupTest()`/`teardownTest()`); every test starts `setupTest(); defer teardownTest()`; schema tests also `defer X_c.Drop()`. No `TestMain`.
- `schema.SetIsTest` switches DB to `devptt_test`; `UnsetIsTest` drops collections and flushes Redis. Never run against real data.
- go-pttbbs backend is mocked in `mockhttp/` (one file per call, numbered variants); gRPC mocked via `boardd`/`mand` `IsTest`; fixtures in `mock/`.
- `-p 1` is mandatory (shared DB). CI: `TAGS=noqueue`, queue test separately. `api/testinit_test.go` has `testSetRequest(...)` for handler tests.
- Coverage: codecov project target 50% (5% threshold); ignores `mock`, `mockhttp`, `apidoc`, `scripts`, `testcase`, `utils`, `queue`. `coverage.txt` is generated.
