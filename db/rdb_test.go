package db

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func TestRDBGet(t *testing.T) {
	setupTest()
	defer teardownTest()

	REDIS_HOST := "localhost:6379"

	rdb := redis.NewClient(&redis.Options{
		Addr:     REDIS_HOST,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		rdb      *redis.Client
		key      string
		want     string
		wantErr  bool
		wantErr2 bool
	}{
		// TODO: Add test cases.
		{
			rdb:  rdb,
			key:  "test-key",
			want: "test-value",
		},
		{
			rdb:  rdb,
			key:  "test-key2",
			want: "test-value2",
		},
		{
			rdb:      rdb,
			key:      "test-key2",
			want:     "test-value2",
			wantErr2: true,
		},
	}
	expireTSDuration := time.Duration(1) * time.Second
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotErr2 := RDBSetNX(tt.rdb, tt.key, tt.want, expireTSDuration)
			if gotErr2 != nil {
				if !tt.wantErr2 {
					t.Errorf("RDBSetNX() failed: %v", gotErr2)
				}
				return
			}
			if tt.wantErr2 {
				t.Errorf("RDBSetNX() succeeded unexpectedly")
			}

			got, gotErr := RDBGet(tt.rdb, tt.key)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("RDBGet() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Errorf("RDBGet() succeeded unexpectedly")
			}

			if got != tt.want {
				t.Errorf("RDBGet() = %v, want %v", got, tt.want)
			}
		})
		wg.Wait()
	}
}

func TestRDBDel(t *testing.T) {
	setupTest()
	defer teardownTest()

	REDIS_HOST := "localhost:6379"

	rdb := redis.NewClient(&redis.Options{
		Addr:     REDIS_HOST,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		rdb      *redis.Client
		key      string
		want     string
		wantErr  bool
		wantErr2 bool
		wantErr3 bool
	}{
		// TODO: Add test cases.
		{
			rdb:  rdb,
			key:  "test-del",
			want: "test-del-value",
		},
	}
	expireTSDuration := time.Duration(1) * time.Second
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()

			gotErr2 := RDBSetNX(tt.rdb, tt.key, tt.want, expireTSDuration)
			if gotErr2 != nil {
				if !tt.wantErr2 {
					t.Errorf("RDBSetNX() failed: %v", gotErr2)
				}
				return
			}
			if tt.wantErr2 {
				t.Errorf("RDBSetNX() succeeded unexpectedly")
			}

			got3, gotErr3 := RDBGet(tt.rdb, tt.key)
			if gotErr3 != nil {
				if !tt.wantErr3 {
					t.Errorf("RDBGet() failed: %v", gotErr3)
				}
				return
			}
			if tt.wantErr {
				t.Errorf("RDBGet() succeeded unexpectedly")
			}

			if got3 != tt.want {
				t.Errorf("RDBGet() = %v, want %v", got3, tt.want)
			}

			gotErr := RDBDel(tt.rdb, tt.key)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("RDBDel() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Errorf("RDBDel() succeeded unexpectedly")
			}

			got4, gotErr4 := RDBGet(tt.rdb, tt.key)
			if gotErr4 == nil {
				t.Errorf("RDBGet(2) succeeded unexpectedly: %v", got4)
			}
		})
		wg.Wait()
	}
}

func TestRDBSAddExpire(t *testing.T) {
	setupTest()
	defer teardownTest()

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()

	key := "test-sadd-expire"
	defer func() { _ = RDBDel(rdb, key) }()

	err := RDBSAdd(rdb, key, "member", 10*time.Second)
	if err != nil {
		t.Fatalf("RDBSAdd() failed: %v", err)
	}

	ttl, err := rdb.TTL(context.Background(), key).Result()
	if err != nil {
		t.Fatalf("TTL() failed: %v", err)
	}
	if ttl <= 0 || ttl > 10*time.Second {
		t.Errorf("TTL() = %v, want in (0, 10s]", ttl)
	}

	err = RDBSAdd(rdb, key, "member", 10*time.Second)
	if err != ErrRDBAlreadyExists {
		t.Errorf("RDBSAdd(dup) err = %v, want %v", err, ErrRDBAlreadyExists)
	}
}
  

func TestRDBIncrWithExpire(t *testing.T) {
	setupTest()
	defer teardownTest()

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	key := "test-incr"
	defer func() { _ = RDBDel(rdb, key) }()

	for i := int64(1); i <= 3; i++ {
		got, err := RDBIncrWithExpire(rdb, key, 1*time.Second)
		if err != nil {
			t.Fatalf("RDBIncrWithExpire() failed: %v", err)
		}
		if got != i {
			t.Errorf("RDBIncrWithExpire() = %v, want %v", got, i)
		}
	}

	time.Sleep(1500 * time.Millisecond)

	got, err := RDBIncrWithExpire(rdb, key, 1*time.Second)
	if err != nil {
		t.Fatalf("RDBIncrWithExpire() after expire failed: %v", err)
	}
	if got != 1 {
		t.Errorf("RDBIncrWithExpire() after expire = %v, want 1", got)
	}
}

func TestRDBIncrWithExpire_restoreMissingTTL(t *testing.T) {
	setupTest()
	defer teardownTest()

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	key := "test-incr-no-ttl"
	defer func() { _ = RDBDel(rdb, key) }()

	// a counter left without TTL (e.g. EXPIRE failed after INCR).
	ctx := context.Background()
	if err := rdb.Set(ctx, key, "2", 0).Err(); err != nil {
		t.Fatalf("Set() failed: %v", err)
	}

	got, err := RDBIncrWithExpire(rdb, key, 10*time.Second)
	if err != nil {
		t.Fatalf("RDBIncrWithExpire() failed: %v", err)
	}
	if got != 3 {
		t.Errorf("RDBIncrWithExpire() = %v, want 3", got)
	}

	ttl, err := rdb.PTTL(ctx, key).Result()
	if err != nil {
		t.Fatalf("PTTL() failed: %v", err)
	}
	if ttl <= 0 {
		t.Errorf("PTTL() = %v, want a positive TTL", ttl)
	}

	// an existing TTL is not extended.
	time.Sleep(200 * time.Millisecond)
	if _, err := RDBIncrWithExpire(rdb, key, 10*time.Second); err != nil {
		t.Fatalf("RDBIncrWithExpire() failed: %v", err)
	}
	ttl2, _ := rdb.PTTL(ctx, key).Result()
	if ttl2 >= ttl {
		t.Errorf("PTTL() = %v after incr, want < %v", ttl2, ttl)
	}
}
