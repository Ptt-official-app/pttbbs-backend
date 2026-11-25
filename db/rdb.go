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

// rdbIncrWithExpireScript runs INCR and the TTL assignment as one atomic step.
//
// The TTL is set whenever the key has none, not only when the counter is created.
// A separate INCR then EXPIRE can be interrupted (timeout, crash) between the two
// calls and leave a counter without TTL, which would never expire.
// Keys that already have a TTL are not extended, so the window stays fixed.
var rdbIncrWithExpireScript = redis.NewScript(`
local value = redis.call('INCR', KEYS[1])
if tonumber(ARGV[1]) > 0 and redis.call('PTTL', KEYS[1]) < 0 then
	redis.call('PEXPIRE', KEYS[1], ARGV[1])
end
return value
`)

// RDBIncrWithExpire increments key and makes sure it has a TTL.
//
// Subsequent increments do not extend the window.
func RDBIncrWithExpire(rdb *redis.Client, key string, expireTSDuration time.Duration) (value int64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), REDIS_TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	value, err = rdbIncrWithExpireScript.Run(ctx, rdb, []string{key}, expireTSDuration.Milliseconds()).Int64()
	if err != nil {
		return 0, err
	}

	return value, nil
}
