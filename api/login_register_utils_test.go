package api

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/pttbbs-backend/types"
)

func Test_gen2FAToken(t *testing.T) {
	tests := []struct {
		name     string
		expected int
	}{
		// TODO: Add test cases.
		{
			expected: 6,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			got := gen2FAToken()
			if len(got) != tt.expected {
				t.Errorf("gen2FAToken() = %v, want %v", got, tt.expected)
			}
		})
		wg.Wait()
	}
}

func Test_genEmailVerificationTokenAndSendEmail(t *testing.T) {
	setupTest()
	defer teardownTest()

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		email    string
		title    string
		url      string
		template string
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			email:    "test@test.dev",
			title:    types.ATTEMPT_REGISTER_USER_TITLE,
			url:      types.REGISTER_USER_URL,
			template: types.ATTEMPT_REGISTER_USER_TEMPLATE_CONTENT,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotErr := genEmailVerificationTokenAndSendEmail(tt.email, tt.title, tt.url, tt.template)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("genEmailVerificationTokenAndSendEmail() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Errorf("genEmailVerificationTokenAndSendEmail() succeeded unexpectedly")
			}
		})
		wg.Wait()
	}
}
