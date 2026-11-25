package api

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
)

func Test_createCSRFToken(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			got, err := createCSRFToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("createCSRFToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !isValidCSRFToken(got) {
				t.Errorf("createCSRFToken: not valid csrf-token: got: %v", got)
			}
		})
	}
	wg.Wait()
}

func Test_isValidCSRFToken(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name     string
		args     args
		expected bool
	}{
		// TODO: Add test cases.
		{
			args:     args{raw: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGkiOiIiLCJleHAiOjEwNTk3MjIyMzQwLCJzdWIiOiIifQ.4Wr9FuxAve-kiHgt51-u4ewtI5CIkRk9tQQXbE5C8HU"},
			expected: true,
		},
		{
			args: args{raw: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGkiOiIiLCJleHAiOjk0OTM0MjI4MTIsInN1YiI6IlNZU09QMiJ9.VbixNBxg4h5FCyTmvhtVzBJ4HsE5_va-MPZzR8TLaY8"},
		},
		{
			args: args{raw: "invalid-raw"},
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if got := isValidCSRFToken(tt.args.raw); got != tt.expected {
				t.Errorf("isValidCSRFToken() = %v, want %v", got, tt.expected)
			}
		})
	}
	wg.Wait()
}

func Test_processCSRFContent(t *testing.T) {
	setupTest()
	defer teardownTest()

	filename := filepath.Join(t.TempDir(), "test.html")
	if err := os.WriteFile(filename, []byte("token=__CSRFTOKEN__"), 0o600); err != nil {
		t.Fatalf("unable to write file: e: %v", err)
	}

	validToken, _ := createCSRFToken()

	tests := []struct {
		name          string
		cookie        string
		wantReuse     bool
		wantNotInBody string
	}{
		{name: "no cookie"},
		{name: "valid cookie is reused", cookie: validToken, wantReuse: true},
		{name: "invalid cookie is not reflected", cookie: "<script>alert(1)</script>", wantNotInBody: "<script>"},
		{name: "expired cookie is replaced", cookie: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGkiOiIiLCJleHAiOjk0OTM0MjI4MTIsInN1YiI6IlNZU09QMiJ9.VbixNBxg4h5FCyTmvhtVzBJ4HsE5_va-MPZzR8TLaY8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			req, _ := http.NewRequest("GET", "/", nil)
			if tt.cookie != "" {
				req.AddCookie(&http.Cookie{Name: types.CSRF_TOKEN, Value: tt.cookie})
			}
			c.Request = req

			processCSRFContent(filename, 60, c)

			if w.Code != http.StatusOK {
				t.Fatalf("code: %v", w.Code)
			}
			if got := w.Header().Get("Cache-Control"); got != "private, max-age=60" {
				t.Errorf("Cache-Control: %v", got)
			}

			body := w.Body.String()
			token := strings.TrimPrefix(body, "token=")
			if !isValidCSRFToken(token) {
				t.Errorf("token in body is not valid: %v", body)
			}
			if tt.wantReuse && token != tt.cookie {
				t.Errorf("token not reused: got: %v want: %v", token, tt.cookie)
			}
			if !tt.wantReuse && tt.cookie != "" && token == tt.cookie {
				t.Errorf("invalid cookie reused: %v", token)
			}
			if tt.wantNotInBody != "" && strings.Contains(body, tt.wantNotInBody) {
				t.Errorf("body contains %v: %v", tt.wantNotInBody, body)
			}
		})
	}
}

func Test_processCSRFContentFileNotFound(t *testing.T) {
	setupTest()
	defer teardownTest()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/", nil)

	processCSRFContent(filepath.Join(t.TempDir(), "not-exists.html"), 60, c)

	if w.Code != http.StatusNotFound {
		t.Errorf("code: %v, want 404", w.Code)
	}
}
