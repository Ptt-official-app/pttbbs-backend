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
