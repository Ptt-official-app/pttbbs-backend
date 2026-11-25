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

var rdbDelIfEqualScript = redis.NewScript(`
if redis.call("GET", KEYS[1]) == ARGV[1] then
	return redis.call("DEL", KEYS[1])
end
return 0
`)

// RDBDelIfEqual deletes key only if its current value equals value (compare-and-delete).
func RDBDelIfEqual(rdb *redis.Client, key string, value string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), REDIS_TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	return rdbDelIfEqualScript.Run(ctx, rdb, []string{key}, value).Err()
}

// RDBIncrWithExpire increments key and sets its TTL when the key is newly created.
//
// The TTL is set only when the counter goes from 0 to 1, so subsequent increments
// do not extend the window.
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
