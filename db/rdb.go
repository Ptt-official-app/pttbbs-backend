package db

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func RDBGetAllKeys(rdb *redis.Client, pattern string) (keys []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), REDIS_TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	cmd := rdb.Keys(ctx, pattern)
	return cmd.Result()
}

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

func RDBSetNX(rdb *redis.Client, key string, value string, expireTSDuration time.Duration) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), REDIS_TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		ctxErr := ctx.Err()
		cancel()
		if err == nil {
			err = ctxErr
		}
	}()

	val, err := rdb.SetNX(ctx, key, value, expireTSDuration).Result()
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
