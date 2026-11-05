package lemmyutils

import "github.com/Ptt-official-app/pttbbs-backend/types"

var (
	LEMMY_ADMIN_TOKEN                         = ""
	LEMMY_ADMIN_TOKEN_NS         types.NanoTS = 0
	LEMMY_ADMIN_TOKEN_REFRESH_NS              = 3600 * types.TS_TO_NANO_TS
)
