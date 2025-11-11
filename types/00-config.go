package types

import (
	"time"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

var (
	SERVICE_MODE = DEV // can be DEV, PRODUCTION, INFO, DEBUG

	HTTP_HOST       = "localhost:3457"            // serving http-host
	URL_PREFIX      = "http://localhost:3457/bbs" // advertising url-prefix
	BACKEND_PREFIX  = "http://localhost:3456/v1"  // backend url-prefix
	FRONTEND_PREFIX = "http://localhost:3457"     // frontend-prefix, email
	API_PREFIX      = "/api"                      // api-prefix

	PTTSYSOP = bbs.UUserID("SYSOP")

	BBSNAME  = "新批踢踢" /* 中文站名 */
	BBSENAME = "PTT2" /* 英文站名 */

	// web
	STATIC_DIR           = "docs/examples"
	ALLOW_ORIGINS        = []string{}
	ALLOW_ORIGINS_MAP    = map[string]bool{}
	BLOCKED_REFERERS     = []string{}
	BLOCKED_REFERERS_MAP = map[string]bool{}
	IS_ALLOW_CROSSDOMAIN = true

	COOKIE_DOMAIN       = "localhost"
	TOKEN_COOKIE_SUFFIX = "Secure;"

	CSRF_SECRET            = []byte("test_csrf_secret")
	CSRF_TOKEN             = "csrftoken"
	CSRF_TOKEN_TS          = 3600 // csrf-token expires in 1 hour.
	CSRF_TOKEN_TS_DURATION = time.Duration(CSRF_TOKEN_TS) * time.Second

	ACCESS_TOKEN_NAME               = "token" // access-token-name in cookie
	ACCESS_TOKEN_EXPIRE_TS          = 86400
	ACCESS_TOKEN_EXPIRE_TS_DURATION = time.Duration(ACCESS_TOKEN_EXPIRE_TS) * time.Second

	IS_OVER_18_NAME  = "over18"
	IS_OVER_18_VALUE = "1"

	// email
	EMAIL_TOKEN_NAME = "token" // email-token in email-url

	EMAIL_FROM   = "noreply@localhost"
	EMAIL_SERVER = "localhost:25"

	EMAILTOKEN_TITLE            = "更換 " + BBSNAME + " 的聯絡信箱 (Updating " + BBSENAME + " Contact Email)"
	IDEMAILTOKEN_TITLE          = "更換 " + BBSNAME + " 的認證信箱 (Updating " + BBSENAME + " Identity Email)"
	ATTEMPT_REGISTER_USER_TITLE = "註冊 " + BBSNAME + " 的確認碼 (Registering " + BBSENAME + " Confirmation Code)"

	EMAILTOKEN_TEMPLATE                    = "docs/etc/emailtoken.template"
	IDEMAILTOKEN_TEMPLATE                  = "docs/etc/idemailtoken.template"
	ATTEMPT_REGISTER_USER_TEMPLATE         = "docs/etc/attemptregister.template"
	ATTEMPT_REGISTER_USER_TEMPLATE_CONTENT = "__USER__, __TOKEN__"

	EMAILTOKEN_TEMPLATE_CONTENT   = "__EMAIL__, __USER__, __URL__"
	IDEMAILTOKEN_TEMPLATE_CONTENT = "__EMAIL__, __USER__, __URL__"

	EXPIRE_USER_ID_EMAIL_IS_SET_NANO_TS = NanoTS(100 * 86400 * 1000000000) // 100 days
	EXPIRE_USER_EMAIL_IS_SET_NANO_TS    = NanoTS(1 * 86400 * 1000000000)   // 1 day

	EXPIRE_USER_ID_EMAIL_IS_NOT_SET_NANO_TS = NanoTS(300 * 1000000000) // 5 mins
	EXPIRE_USER_EMAIL_IS_NOT_SET_NANO_TS    = NanoTS(300 * 1000000000) // 5 mins

	EXPIRE_ATTEMPT_REGISTER_USER_EMAIL_TS          = 300
	EXPIRE_ATTEMPT_REGISTER_USER_EMAIL_TS_DURATION = time.Duration(EXPIRE_ATTEMPT_REGISTER_USER_EMAIL_TS) * time.Second // 5 mins

	// 2fa
	IS_2FA              = true
	MAX_2FA_TOKEN int64 = 10000

	// big5
	BIG5_TO_UTF8 = "types/uao250-b2u.big5.txt"
	UTF8_TO_BIG5 = "types/uao250-u2b.big5.txt"
	AMBCJK       = "types/ambcjk.big5.txt"

	// time-location
	TIME_LOCATION = "Asia/Taipei"

	// carriage-return
	IS_CARRIAGE_RETURN = true

	// is-all-guest
	IS_ALL_GUEST = false

	// pttweb-hotboard-url
	PTTWEB_HOTBOARD_URL = "http://localhost:3457/static/ptt_cc_websites/HotBoards.html"

	// pttweb-base-url
	PTTWEB_BASE_URL = "http://localhost:3457/static/ptt_cc_websites"

	// expire-http-request-ts
	EXPIRE_HTTP_REQUEST_TS = 10

	// max-popular-boards
	MAX_POPULAR_BOARDS = 128
)
