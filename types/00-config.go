package types

import (
	"time"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

var (
	SERVICE_MODE = DEV // can be DEV, PRODUCTION, STAGING, INFO, DEBUG

	HTTP_SCHEME      = "http://"
	HTTP_HOST        = "localhost:3457"            // serving http-host
	URL_PREFIX       = "http://localhost:3457/bbs" // advertising url-prefix (for constructing article url)
	GO_PTTBBS_PREFIX = "http://localhost:3456/v1"  // backend url-prefix

	FRONTEND_PREFIX    = "http://localhost:5173" // frontend-prefix, email
	FRONTEND_LOGIN_URL = "http://localhost:5173/login"
	FRONTEND_INIT_URL  = "http://localhost:5173/init"
	FRONTEND_ERR_URL   = "http://localhost:5173/error"

	API_PREFIX        = "/api" // api-prefix
	REGISTER_USER_URL = "http://localhost:3457/api/account/register"
	ZK_PREFIX         = "http://localhost:3458"
	EMAIL_URL         = "http://localhost:3457" // for email url

	OIDC_OP_KEY             = "oidc-op-key"
	OIDC_OP_KEY_ID          = "oidc-op-key-id"
	OIDC_OP_POST_LOGOUT_URL = "http://localhost:5173"
	OIDC_OP_ISSUER          = "http://localhost:3457"
	OIDC_OP_IS_ALLOW_HTTP   = true

	PTTSYSOP = bbs.UUserID("SYSOP")

	BBSNAME       = "新批踢踢"   // site-name
	BBSNAME_EN    = "NeoPTT" // English site-name
	SENDER_SUFFIX = "管理員"    // for utils.SendEmail

	// web
	STATIC_DIR = "docs/examples"

	ALLOW_ORIGINS        = []string{"*"}
	BLOCKED_REFERERS     = []string{}
	IS_ALLOW_CROSSDOMAIN = true

	COOKIE_DOMAIN       = "localhost"
	TOKEN_COOKIE_SUFFIX = ""

	CSRF_SECRET            = []byte("test_csrf_secret")
	CSRF_TOKEN             = "csrftoken"
	CSRF_TOKEN_TS          = 3600 // csrf-token expires in 1 hour.
	CSRF_TOKEN_TS_DURATION = time.Duration(CSRF_TOKEN_TS) * time.Second
	CSRF_COOKIE_DOMAIN     = "localhost"

	ACCESS_TOKEN_NAME               = "token" // access-token-name in cookie
	ACCESS_TOKEN_EXPIRE_TS          = 86400
	ACCESS_TOKEN_EXPIRE_TS_DURATION = time.Duration(ACCESS_TOKEN_EXPIRE_TS) * time.Second
	ACCESS_TOKEN_SECRET             = []byte("access_token_secret")

	REFRESH_TOKEN_NAME               = "refresh_token" // refresh-token-name in cookie
	REFRESH_TOKEN_EXPIRE_TS          = 86400
	REFRESH_TOKEN_EXPIRE_TS_DURATION = time.Duration(REFRESH_TOKEN_EXPIRE_TS) * time.Second
	REFRESH_TOKEN_SECRET             = []byte("refresh_token_secret")

	IS_OVER_18_NAME  = "over18"
	IS_OVER_18_VALUE = "1"

	// email
	EMAIL_TOKEN_NAME = "token" // email-token in email-url

	EMAIL_FROM   = "noreply@localhost"
	EMAIL_SERVER = "localhost:25"

	EMAILTOKEN_TITLE_TEMPLATE   = "增加 __BBSNAME__ 的聯絡信箱 (Adding __BBSNAME_EN__ Contact Email)"
	EMAILTOKEN_TITLE            = "增加 " + BBSNAME + " 的聯絡信箱 (Adding " + BBSNAME_EN + " Contact Email)"
	EMAILTOKEN_TEMPLATE         = "docs/etc/emailtoken.template"
	EMAILTOKEN_TEMPLATE_CONTENT = "__EMAIL__, __USER__, __URL__"

	IDEMAILTOKEN_TITLE_TEMPLATE   = "更換 __BBSNAME__ 的認證信箱 (Updating __BBSNAME_EN__ Identity Email)"
	IDEMAILTOKEN_TITLE            = "更換 " + BBSNAME + " 的認證信箱 (Updating " + BBSNAME_EN + " Identity Email)"
	IDEMAILTOKEN_TEMPLATE         = "docs/etc/idemailtoken.template"
	IDEMAILTOKEN_TEMPLATE_CONTENT = "__EMAIL__, __USER__, __URL__"

	ATTEMPT_REGISTER_USER_TITLE_TEMPLATE           = "註冊 __BBSNAME__  的驗證連結 (Registering __BBSNAME_EN__ Verification Link)"
	ATTEMPT_REGISTER_USER_TITLE                    = "註冊 " + BBSNAME + " 的驗證連結 (Registering " + BBSNAME_EN + " Verification Link)"
	ATTEMPT_REGISTER_USER_TEMPLATE                 = "docs/etc/attemptregister.template"
	ATTEMPT_REGISTER_USER_TEMPLATE_CONTENT         = "__EMAIL__, __USER__, __URL__"
	EXPIRE_ATTEMPT_REGISTER_USER_EMAIL_TS          = 300
	EXPIRE_ATTEMPT_REGISTER_USER_EMAIL_TS_DURATION = time.Duration(EXPIRE_ATTEMPT_REGISTER_USER_EMAIL_TS) * time.Second // 5 mins

	ATTEMPT_LOGIN_TITLE_TEMPLATE           = "登入 __BBSNAME__ 的驗證碼 (Login __BBSNAME_EN__ Verification Code)"
	ATTEMPT_LOGIN_TITLE                    = "登入 " + BBSNAME + " 的驗證碼 (Login " + BBSNAME_EN + " Verification Code)"
	ATTEMPT_LOGIN_TEMPLATE                 = "docs/etc/attemptlogin.template"
	ATTEMPT_LOGIN_TEMPLATE_CONTENT         = "__USER__, __TOKEN__"
	EXPIRE_ATTEMPT_LOGIN_EMAIL_TS          = 300
	EXPIRE_ATTEMPT_LOGIN_EMAIL_TS_DURATION = time.Duration(EXPIRE_ATTEMPT_LOGIN_EMAIL_TS) * time.Second // 5 mins

	EXPIRE_USER_ID_EMAIL_IS_SET_NANO_TS = NanoTS(100 * 86400 * 1000000000) // 100 days
	EXPIRE_USER_EMAIL_IS_SET_NANO_TS    = NanoTS(1 * 86400 * 1000000000)   // 1 day

	EXPIRE_USER_ID_EMAIL_IS_NOT_SET_NANO_TS = NanoTS(300 * 1000000000) // 5 mins
	EXPIRE_USER_EMAIL_IS_NOT_SET_NANO_TS    = NanoTS(300 * 1000000000) // 5 mins

	EXPIRE_OIDC_AUTH_REQUEST_TS          = 300
	EXPIRE_OIDC_AUTH_REQUEST_TS_DURATION = time.Duration(EXPIRE_OIDC_AUTH_REQUEST_TS) * time.Second

	// 2fa
	IS_2FA                         = true
	MAX_2FA_TOKEN            int64 = 1000000
	MAX_2FA_TOKEN_STR_PROMPT       = "%06d"

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
	PTTWEB_HOTBOARD_URL = "http://localhost:3457/static/ptt_cc_websites/bbs/HotBoards.html"

	// pttweb-base-url (excluding bbs/, man/, cls/)
	PTTWEB_BASE_URL = "http://localhost:3457/static/ptt_cc_websites"

	// expire-http-request-ts
	EXPIRE_HTTP_REQUEST_TS = 10

	// max-popular-boards
	MAX_POPULAR_BOARDS = 128

	BRDNAME_WHITE_LIST_MAP_FILENAME = ""

	BRDNAME_BLACK_LIST_MAP_FILENAME = ""

	SLEEP_RETRY_LOAD_POPULAR_BOARDS_TS          = 3600
	SLEEP_RETRY_LOAD_POPULAR_BOARDS_TS_DURATION = time.Duration(SLEEP_RETRY_LOAD_POPULAR_BOARDS_TS) * time.Second

	SLEEP_RETRY_LOAD_GENERAL_ARTICLES_TS          = 3600
	SLEEP_RETRY_LOAD_GENERAL_ARTICLES_TS_DURATION = time.Duration(SLEEP_RETRY_LOAD_GENERAL_ARTICLES_TS) * time.Second

	SLEEP_RETRY_LOAD_ARTICLE_DETAILS_TS          = 3600
	SLEEP_RETRY_LOAD_ARTICLE_DETAILS_TS_DURATION = time.Duration(SLEEP_RETRY_LOAD_ARTICLE_DETAILS_TS) * time.Second
)
