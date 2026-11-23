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
			args: args{userID: "SYSOP", token: "123456", expireTSDuration: time.Duration(1) * time.Second},
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
			if err := Set2FA(tt.args.userID, tt.args.token, tt.args.email, "", tt.args.expireTSDuration); (err != nil) != tt.wantErr {
				t.Errorf("Set2FA() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		wg.Wait()
	}
}

func TestGet2FA(t *testing.T) {
	setupTest()
	defer teardownTest()

	_ = Set2FA("SYSOP", "test@ptt.test", "123456", "", time.Duration(1)*time.Second)

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
			if tt.wantErr == true {
				return
			}

			if gotToken.Token != tt.expectedToken {
				t.Errorf("Get2FA() = %v, want %v", gotToken, tt.expectedToken)
			}
		})
	}
}

func TestDelete2FA(t *testing.T) {
	setupTest()
	defer teardownTest()

	_ = Set2FA("SYSOP", "test@ptt.test", "123456", "", time.Duration(10)*time.Second)

	if _, err := Get2FA("SYSOP"); err != nil {
		t.Fatalf("Get2FA() before delete: e: %v", err)
	}

	if err := Delete2FA("SYSOP"); err != nil {
		t.Fatalf("Delete2FA() e: %v", err)
	}

	if got, err := Get2FA("SYSOP"); err == nil {
		t.Errorf("Get2FA() after delete succeeded unexpectedly: %v", got)
	}
}

func TestIncr2FAAttempt(t *testing.T) {
	setupTest()
	defer teardownTest()

	for i := int64(1); i <= 3; i++ {
		got, err := Incr2FAAttempt("SYSOP", time.Duration(10)*time.Second)
		if err != nil {
			t.Fatalf("Incr2FAAttempt() e: %v", err)
		}
		if got != i {
			t.Errorf("Incr2FAAttempt() = %v, want %v", got, i)
		}
	}

	// Set2FA does not reset the attempt counter.
	_ = Set2FA("SYSOP", "test@ptt.test", "123456", "", time.Duration(10)*time.Second)
	got, _ := Incr2FAAttempt("SYSOP", time.Duration(10)*time.Second)
	if got != 4 {
		t.Errorf("Incr2FAAttempt() after Set2FA = %v, want 4", got)
	}

	if err := Reset2FAAttempt("SYSOP"); err != nil {
		t.Fatalf("Reset2FAAttempt() e: %v", err)
	}
	got, _ = Incr2FAAttempt("SYSOP", time.Duration(10)*time.Second)
	if got != 1 {
		t.Errorf("Incr2FAAttempt() after reset = %v, want 1", got)
	}
}
