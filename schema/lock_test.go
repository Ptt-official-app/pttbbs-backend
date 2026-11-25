package schema

import (
	"sync"
	"testing"
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTryLock(t *testing.T) {
	setupTest()
	defer teardownTest()

	expireNanoTS := time.Duration(2) * time.Second
	type args struct {
		key    string
		expire time.Duration
	}
	tests := []struct {
		name        string
		args        args
		sleepTS     int
		isUnlock    bool
		wantErr     bool
		expectedErr error
	}{
		// TODO: Add test cases.
		{
			args: args{key: "test1", expire: expireNanoTS},
		},
		{
			args:        args{key: "test1", expire: expireNanoTS},
			wantErr:     true,
			expectedErr: db.ErrRDBAlreadyExists,
		},
		{
			args:     args{key: "test1", expire: expireNanoTS},
			isUnlock: true,
		},
		{
			args: args{key: "test2", expire: expireNanoTS},
		},
		{
			args:    args{key: "test1", expire: expireNanoTS},
			sleepTS: 4,
		},
	}

	tokens := map[string]string{}
	var wg sync.WaitGroup // to sequentially exec the test.
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if tt.sleepTS > 0 {
				time.Sleep(time.Duration(tt.sleepTS) * time.Second)
			}
			if tt.isUnlock {
				_ = Unlock(tt.args.key, tokens[tt.args.key])
			}
			token, err := TryLock(tt.args.key, tt.args.expire)
			if err == nil {
				tokens[tt.args.key] = token
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("TryLock() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != tt.expectedErr {
				t.Errorf("TryLock: e: %v expected: %v", err, tt.expectedErr)
			}
		})
		wg.Wait()
	}
}

func TestUnlock_notOwner(t *testing.T) {
	setupTest()
	defer teardownTest()

	key := "test-unlock-not-owner"

	// A holds the lock, it expires, then B acquires it.
	tokenA, err := TryLock(key, 100*time.Millisecond)
	require.NoError(t, err)
	time.Sleep(200 * time.Millisecond)
	tokenB, err := TryLock(key, 10*time.Second)
	require.NoError(t, err)
	require.NotEqual(t, tokenA, tokenB)

	// A's late Unlock must not release B's lock.
	err = Unlock(key, tokenA)
	require.NoError(t, err)
	_, err = TryLock(key, 10*time.Second)
	assert.ErrorIs(t, err, db.ErrRDBAlreadyExists)

	// the owner can release it.
	err = Unlock(key, tokenB)
	require.NoError(t, err)
	_, err = TryLock(key, 10*time.Second)
	assert.NoError(t, err)
}
