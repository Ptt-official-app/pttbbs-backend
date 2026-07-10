package schema

import (
	"sync"
	"testing"
	"time"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

func TestSet2FA(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		userID           bbs.UUserID
		email            string
		token            string
		expireTSDuration time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		sleep   time.Duration
	}{
		// TODO: Add test cases.
		{
			args: args{userID: "SYSOP", token: "123456", expireTSDuration: time.Duration(1) * time.Second},
		},
		{
			args:    args{userID: "SYSOP", token: "123456", expireTSDuration: time.Duration(1) * time.Second},
			wantErr: true,
		},
		{
			sleep: time.Duration(5) * time.Second,
			args:  args{userID: "SYSOP", token: "123456", expireTSDuration: time.Duration(1) * time.Second},
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()

			time.Sleep(tt.sleep)
			if err := Set2FA(tt.args.userID, tt.args.token, tt.args.email, tt.args.expireTSDuration); (err != nil) != tt.wantErr {
				t.Errorf("Set2FA() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		wg.Wait()
	}
}

func TestGet2FA(t *testing.T) {
	setupTest()
	defer teardownTest()

	_ = Set2FA("SYSOP", "test@ptt.test", "123456", time.Duration(1)*time.Second)

	type args struct {
		userID bbs.UUserID
	}
	tests := []struct {
		name          string
		args          args
		expectedToken string
		wantErr       bool
	}{
		// TODO: Add test cases.
		{
			args:          args{userID: "SYSOP"},
			expectedToken: "123456",
		},
		{
			args:    args{userID: "SYSOP2"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotToken, err := Get2FA(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get2FA() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotToken != tt.expectedToken {
				t.Errorf("Get2FA() = %v, want %v", gotToken, tt.expectedToken)
			}
		})
	}
}
