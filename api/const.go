package api

import (
	"time"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

const (
	DEFAULT_MAX_ARTICLE_BLOCK_LIST = 1
	DEFAULT_MAX_LIST               = 200
	DEFAULT_DESCENDING             = true
	DEFAULT_ASCENDING              = true

	ARTICLE_LOCK_TS          = 10
	ARTICLE_LOCK_TS_DURATION = time.Duration(ARTICLE_LOCK_TS) * time.Second
	ARTICLE_LOCK_NANO_TS     = types.NanoTS(ARTICLE_LOCK_TS) * types.TS_TO_NANO_TS

	HTML_CACHE_CONTROL_TS = 3600
	JS_CACHE_CONTROL_TS   = 86400 * 365

	OFFSET_MTIME_NANO_TS = 1000000000 // 1 second

	REFRESH_IS_OVER18_NANO_TS = 86400 * 1000000000
)

var (
	GET_ARTICLE_CONTENT_INFO_TOO_SOON_NANO_TS = ARTICLE_LOCK_NANO_TS + types.NanoTS(1)*types.TS_TO_NANO_TS // 10 + 1 seconds.

	MIME_TYPE_MAP = map[string]string{
		".js":   "text/javascript",
		".html": "text/html",
		".map":  "application/json",
	}

	BOARD_ID_WHITE_LIST_MAP = map[bbs.BBoardID]bool{}
)

const (
	MAX_USER_FAVORITES_BUF_SIZE = 65535
)

const (
	DEFAULT_FAV_TITLE = "新的目錄"
)

const (
	IS_OVER_18_SUBSTRING = "document.cookie.indexOf('over18=1') === -1)"
)
