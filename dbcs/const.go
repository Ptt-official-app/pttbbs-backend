package dbcs

import (
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/types"
)

const (
	N_FIRST_COMMENTS                = 10
	COMMENT_STEP_DURATION           = 1 * time.Millisecond
	REPLY_STEP_NANO_TS              = 100000    // 0.1 millisecond
	DELETE_STEP_NANO_TS             = 10000     // 0.01 milliseond
	COMMENT_STEP_NANO_TS            = 1000000   // 1 millisecond
	COMMENT_EXCEED_NANO_TS          = 1000      // 0.001 millisecond
	COMMENT_BACKWARD_OFFSET_NANO_TS = 900000000 // 900 millisecond
	COMMENT_DIFF_ALIGN_END_NANO_TS  = 60 * types.TS_TO_NANO_TS
	COMMENT_DIFF2_ALIGN_END_NANO_TS = 86400 * types.TS_TO_NANO_TS

	DEFAULT_LINE_BYTES = 200

	LEN_OLD_RECOMMEND_DATE = 5
	LEN_RECOMMEND_DATE     = 11
	LEN_EDIT_DATE_TIME     = 19 // MM/DD/YYYY HH:mm:ss

	ONE_YEAR_OFFSET_NANO_TS = 365 * 86400 * types.TS_TO_NANO_TS

	N_LINES_PER_CONTENT_BLOCK = 50

	MAX_COMMENT_BYTES = 81
)
