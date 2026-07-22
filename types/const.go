package types

import "time"

var (
	TIMEZONE, _ = time.LoadLocation(TIME_LOCATION)
	VERSION     = ""
	GIT_VERSION = ""

	LEN_DATE_YEAR_TIME_STR = len("01/02/2006 15:04:05")
	LEN_DATE_MIN_STR       = len("2006/01/02 15:04")

	COLOR_PREFIX_BYTES  = []byte("\x1b[")
	COLOR_POSTFIX_BYTES = []byte("m")
	COLOR_RESET_BYTES   = []byte("\x1b[m")

	DEFAULT_LEN_COLOR_BYTES = 20 // \x1b[0;1;5;37;40m

	POSTTIME_REJECT = 15

	TEMPLATE_EMAIL      = "__EMAIL__"
	TEMPLATE_TOKEN      = "__TOKEN__"
	TEMPLATE_USER       = "__USER__"
	TEMPLATE_URL        = "__URL__"
	TEMPLATE_BBSNAME    = "__BBSNAME__"
	TEMPLATE_BBSNAME_EN = "__BBSNAME_EN__"

	ZK_USER_ID_KEY = "PTTBBS_USER_ID"

	WEB_CLIENT_ID = "web"
)
