package schema

import (
	"strings"
	"sync"
	"testing"
	"time"
)

func TestSetEmailVerification(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		email            string
		expireTSDuration time.Duration
		wantErr          bool
		wantErr2         bool
	}{
		// TODO: Add test cases.
		{
			email:            "test@email.dev",
			expireTSDuration: time.Duration(1) * time.Second,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			got, gotErr := SetEmailVerification(tt.email, tt.expireTSDuration)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("SetEmailVerification() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Errorf("SetEmailVerification() succeeded unexpectedly")
			}

			if strings.HasSuffix(got, "=") {
				t.Errorf("SetEmailVerification: unexpected = in got: %v", got)
			}

			emailSalt, gotErr2 := GetEmailVerification(got)
			if gotErr2 != nil {
				if !tt.wantErr2 {
					t.Errorf("GetEmailVerification() failed: e: %v", gotErr2)
				}
				return
			}
			if emailSalt.Email != tt.email {
				t.Errorf("GetEmailVerification() = %v, want %v", emailSalt.Email, tt.email)
			}

			isValid, gotErr3 := VerifyEmailVerification(got, emailSalt)
			if gotErr3 != nil {
				t.Errorf("VerifyEmailVerification failed e: %v", gotErr3)
			}
			if !isValid {
				t.Errorf("VerifyEmailVerification invalid: got: %v emailSalt: %v", got, emailSalt)
			}

			got4, gotErr4 := SetEmailVerification(tt.email, tt.expireTSDuration)
			if gotErr4 != nil {
				if !tt.wantErr {
					t.Errorf("SetEmailVerification(2) failed: %v", gotErr4)
				}
				return
			}
			if strings.HasSuffix(got4, "=") {
				t.Errorf("SetEmailVerification(2): unexpected = in got: %v", got4)
			}

			if got == got4 {
				t.Errorf("SetEmailVerification(2): same token: got: %v got4: %v", got, got4)
			}

			emailSalt5, gotErr5 := GetEmailVerification(got4)
			if gotErr5 != nil {
				if !tt.wantErr2 {
					t.Errorf("GetEmailVerification(5) failed: e: %v", gotErr5)
				}
				return
			}
			if emailSalt5.Email != tt.email {
				t.Errorf("GetEmailVerification(5) = %v, want %v", emailSalt5.Email, tt.email)
			}

			isValid6, gotErr6 := VerifyEmailVerification(got4, emailSalt5)
			if gotErr6 != nil {
				t.Errorf("VerifyEmailVerification(6) failed e: %v", gotErr6)
			}
			if !isValid6 {
				t.Errorf("VerifyEmailVerification(6) invalid: got4: %v emailSalt5: %v", got4, emailSalt5)
			}

			time.Sleep(tt.expireTSDuration)

			emailSalt7, gotErr7 := GetEmailVerification(got)
			if gotErr7 == nil {
				t.Errorf("GetEmailVerification(7) succeed unexpectedly: emailSalt7: %v", emailSalt7)
			}

			emailSalt8, gotErr8 := GetEmailVerification(got4)
			if gotErr8 == nil {
				t.Errorf("GetEmailVerification(8) succeed unexpectedly: emailSalt8: %v", emailSalt8)
			}
		})
		wg.Wait()
	}
}

func Test_b64WithPadding(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		theStr string
		want   string
	}{
		// TODO: Add test cases.
		{
			theStr: "",
			want:   "",
		},
		{
			theStr: "a",
			want:   "a===",
		},
		{
			theStr: "aa",
			want:   "aa==",
		},
		{
			theStr: "aaa",
			want:   "aaa=",
		},
		{
			theStr: "aaaa",
			want:   "aaaa",
		},
		{
			theStr: "aaaaa",
			want:   "aaaaa===",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := b64WithPadding(tt.theStr)
			if got != tt.want {
				t.Errorf("b64WithPadding() = %v, want %v", got, tt.want)
			}
		})
	}
}
