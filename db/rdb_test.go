package db

import (
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
