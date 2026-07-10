package schema

import (
	"sync"
	"testing"
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/db"
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

	var wg sync.WaitGroup // to sequentially exec the test.
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if tt.sleepTS > 0 {
				time.Sleep(time.Duration(tt.sleepTS) * time.Second)
			}
			if tt.isUnlock {
				Unlock(tt.args.key)
			}
			err := TryLock(tt.args.key, tt.args.expire)

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
