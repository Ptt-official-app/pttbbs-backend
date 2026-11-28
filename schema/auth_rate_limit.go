package schema

import (
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/db"
)

// IncrAuthRateLimit counts a login / register request from clientIP and returns
// the number of requests in the current fixed window.
func IncrAuthRateLimit(clientIP string, windowDuration time.Duration) (count int64, err error) {
	return db.RDBIncrWithExpire(rdb, RDB_PREFIX_AUTH_RATE_LIMIT+clientIP, windowDuration)
}
