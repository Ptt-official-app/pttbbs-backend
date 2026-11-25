package db

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func RDBGet(rdb *redis.Client, key string) (value string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), REDIS_TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	value, err = rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return value, nil
}

func RDBSAdd(rdb *redis.Client, key string, value string, expireTSDuration time.Duration) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), REDIS_TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	val, err := rdb.SAdd(ctx, key, value).Result()
	if err != nil {
		return err
	}
	if expireTSDuration > 0 {
		err = rdb.Expire(ctx, key, expireTSDuration).Err()
		if err != nil {
			return err
		}
	}
	if val == 0 {
		return ErrRDBAlreadyExists
	}
	return nil
}

func RDBSMembers(rdb *redis.Client, key string) (values []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), REDIS_TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	values, err = rdb.SMembers(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if values == nil {
		return []string{}, nil
	}

	return values, nil
}

func RDBSRem(rdb *redis.Client, key string, value string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), REDIS_TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	_, err = rdb.SRem(ctx, key, value).Result()
	if err != nil {
		return err
	}

	return nil
}

func RDBSet(rdb *redis.Client, key string, value string, expireTSDuration time.Duration) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), REDIS_TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	theExpireTSDuration := expireTSDuration
	if theExpireTSDuration == 0 {
		theExpireTSDuration = redis.KeepTTL
	}

	return rdb.Set(ctx, key, value, theExpireTSDuration).Err()
}

func RDBSetNX(rdb *redis.Client, key string, value string, expireTSDuration time.Duration) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), REDIS_TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	theExpireTSDuration := expireTSDuration
	if theExpireTSDuration == 0 {
		theExpireTSDuration = redis.KeepTTL
	}

	val, err := rdb.SetNX(ctx, key, value, theExpireTSDuration).Result()
	if err != nil {
		return err
	}
	if !val {
		return ErrRDBAlreadyExists
	}

	return nil
}

func RDBDel(rdb *redis.Client, key string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), REDIS_TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	_, err = rdb.Del(ctx, key).Result()
	if err != nil {
		return err
	}

	return nil
}

// RDBIncrWithExpire increments key and sets its TTL when the key is newly created.
//
// The TTL is set only when the counter goes from 0 to 1, so subsequent increments
// do not extend the window.
// RDBDelExisting deletes the key and reports whether this call was the one that removed it.
// Redis executes DEL atomically, so concurrent callers see isDeleted = true exactly once.
func RDBDelExisting(rdb *redis.Client, key string) (isDeleted bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), REDIS_TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	n, err := rdb.Del(ctx, key).Result()
	if err != nil {
		return false, err
	}

	return n == 1, nil
}

func RDBIncrWithExpire(rdb *redis.Client, key string, expireTSDuration time.Duration) (value int64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), REDIS_TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	value, err = rdb.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	if value == 1 && expireTSDuration > 0 {
		err = rdb.Expire(ctx, key, expireTSDuration).Err()
		if err != nil {
			return 0, err
		}
	}

	return value, nil
}
