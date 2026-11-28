package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_processResultNilErr(t *testing.T) {
	gin.SetMode(gin.TestMode)

	for _, statusCode := range []int{400, 403, 500} {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("GET", "/", nil)

		processResult(c, nil, statusCode, nil, "")

		if w.Code != statusCode {
			t.Errorf("processResult(%v, nil err): code: %v", statusCode, w.Code)
		}
		var result errResult
		_ = json.Unmarshal(w.Body.Bytes(), &result)
		if result.Msg != http.StatusText(statusCode) {
			t.Errorf("processResult(%v, nil err): Msg: %v, want %v", statusCode, result.Msg, http.StatusText(statusCode))
		}
	}
}

func TestLoginInvalidParams(t *testing.T) {
	_, statusCode, err := Login("localhost", &UserInfo{}, nil, nil)
	if statusCode != 400 || err != ErrInvalidParams {
		t.Errorf("Login(nil params): statusCode: %v e: %v, want 400 ErrInvalidParams", statusCode, err)
	}
}

func TestProcessResultHidesInternalErrors(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		statusCode int
		err        error
		wantMsg    string
	}{
		{"5xx hides the error", 500, errors.New(`Get "http://localhost:3456/v1/boards": dial tcp 127.0.0.1:3456: connect: connection refused`), "Internal Server Error"},
		{"502 hides the error", 502, errors.New("upstream body"), "Bad Gateway"},
		{"4xx keeps the error", 400, ErrInvalidParams, ErrInvalidParams.Error()},
		{"nil error uses status text", 404, nil, "Not Found"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest(http.MethodGet, "/api/boards", nil)

			processResult(c, nil, tt.statusCode, tt.err, "SYSOP")

			assert.Equal(t, tt.statusCode, w.Code)
			result := &errResult{}
			require.NoError(t, json.Unmarshal(w.Body.Bytes(), result))
			assert.Equal(t, tt.wantMsg, result.Msg)
		})
	}
}
