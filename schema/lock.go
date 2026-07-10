package schema

import (
	"context"
	"strconv"
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

// TryLock
func TryLock(key string, expireTSDuration time.Duration) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), REDIS_TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	updateNanoTS := int64(types.NowNanoTS())
	updateNanoTSStr := strconv.FormatInt(updateNanoTS, 10)

	err = db.RDBSetNX(rdb, RDB_PREFIX_LOCK+key, updateNanoTSStr, expireTSDuration)
	if err != nil {
		return err
	}

	return nil
}

// Unlock
func Unlock(key string) (err error) {
	err = db.RDBDel(rdb, RDB_PREFIX_LOCK+key)
	if err != nil {
		return err
	}

	return nil
}
