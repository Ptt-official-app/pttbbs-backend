package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
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
