package schema

import (
	"context"
	"strconv"
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

// TryLock
//
// The returned token identifies the holder and must be passed to Unlock.
func TryLock(key string, expireTSDuration time.Duration) (token string, err error) {
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
		return "", err
	}

	return updateNanoTSStr, nil
}

// Unlock releases the lock only if it is still held with the given token.
//
// The lock may have expired and been re-acquired by someone else; deleting it
// unconditionally would release that holder's lock.
func Unlock(key string, token string) (err error) {
	return db.RDBDelIfEqual(rdb, RDB_PREFIX_LOCK+key, token)
}
