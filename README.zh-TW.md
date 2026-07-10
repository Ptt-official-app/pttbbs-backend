# pttbbs-backend

[![Go](https://github.com/Ptt-official-app/pttbbs-backend/actions/workflows/go.yml/badge.svg)](https://github.com/Ptt-official-app/pttbbs-backend/actions/workflows/go.yml)
[![GoDoc](https://pkg.go.dev/badge/github.com/Ptt-official-app/pttbbs-backend?status.svg)](https://pkg.go.dev/github.com/Ptt-official-app/pttbbs-backend?tab=doc)
[![codecov](https://codecov.io/gh/Ptt-official-app/pttbbs-backend/branch/main/graph/badge.svg)](https://codecov.io/gh/Ptt-official-app/pttbbs-backend)

## README Translation

* [English](https://github.com/Ptt-official-app/pttbbs-backend/blob/main/README.en.md)
* [正體中文](https://github.com/Ptt-official-app/pttbbs-backend/blob/main/README.zh-TW.md)

## 概觀

這裡是使用 golang 來達成 [PTT 的 中介應用層](https://hackmd.io/@twbbs/Root#%E6%9E%B6%E6%A7%8B%E5%9C%96).

與 [Ptt-official-app pttbbs](https://github.com/ptt-official-app/go-pttbbs) 一起成為現代的 BBS 架構.

## Demo Site

* [dev](https://www.devptt.dev)
* [term (PttChrome)](https://term.devptt.dev)

## Getting Started

您可以用以下方式快速開始:

1. 安裝 [docker](https://www.docker.com/)
2. `./scripts/start-getting-started.sh`
3. 到 [http://localhost:8080](http://localhost:8080) 使用 API.

### [舊版 pttbbs](https://github.com/ptt/pttbbs)

您可以使用以下方式連到 舊版 pttbbs:

4. `telnet localhost 8888`
5. 使用 SYSOP/123123 登入.

## API

您可以到 [https://doc.devptt.dev](https://doc.devptt.dev) 試著使用 api.

如果您在 api 網頁裡遇到 CORS 的問題. 您可以在網頁裡 copy `curl` 指令測試.

## Coding Convention

我們使用以下 library 幫助 coding convention:

* [gotests](https://github.com/cweill/gotests) for test-generation.
* [gofumpt](https://github.com/mvdan/gofumpt) for formatting.

## Discussing / Reviewing / Questioning the code

除了開 issues 以外, 您還可以做以下的事情來對於 code 做討論 / review / 提出問題.

* `git clone` 這個 repo.
* 開一個 review-\[topic\] 的 branch.
* 對於想要討論的部分在 code 裡寫 comments.
* pull-request
* 對於 PR 進行討論.

## Develop

您可以使用 fork 來一起開發.

## Unit-Test

你可以做以下的事情來進行 unit-test:

* `./scripts/test.sh`

您可以做以下的事情來進行 coverage-check:

* `./scripts/coverage.sh`

## Swagger

You can run swagger with:

* 設定 python virtualenv.
* `cd apidoc; pip install .; cd ..`
* `./scripts/swagger.sh [host]`
* go to [http://localhost:8080](http://localhost:8080)

## Schema definition

* [https://github.com/Ptt-official-app/pttbbs-backend/tree/main/schema](https://github.com/Ptt-official-app/pttbbs-backend/tree/main/schema)

## Repository Naming

* 2024-06-26: 這個 repo 主要是以 pttbbs 的機制為主, 所以改成命名為 pttbbs-backend.
* 2022-12-16: 這個 repo 之所以會被稱為 go-openbbsmiddleware, 是因為古早的 [.NET ASP](https://github.com/Ptt-official-app/AspCoreOpenBBSMiddleware) 的開發者希望 [中介應用層](https://hackmd.io/@twbbs/Root#%E6%9E%B6%E6%A7%8B%E5%9C%96) 可以擴及其他版本的 bbs (Maple/中山之島). 這個 repo 就 follow 當時的 naming convention 命名為 go-openbbsmiddleware.
