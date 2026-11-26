package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMergeURL(t *testing.T) {
	urlMap1 := make(map[string]string)
	urlMap1["bid"] = "test1"
	url1 := "/board/:bid/test"
	expected1 := "/board/test1/test"

	url2 := "/board/:bid"
	expected2 := "/board/test1"

	url3 := "/board/:bidtest"
	expected3 := "/board/"

	url4 := "/board/:b/:bid"
	expected4 := "/board//test1"

	type args struct {
		urlMap map[string]string
		url    string
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		// TODO: Add test cases.
		{
			args:     args{urlMap: urlMap1, url: url1},
			expected: expected1,
		},
		{
			args:     args{urlMap: urlMap1, url: url2},
			expected: expected2,
		},
		{
			args:     args{urlMap: urlMap1, url: url3},
			expected: expected3,
		},
		{
			args:     args{urlMap: urlMap1, url: url4},
			expected: expected4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeURL(tt.args.urlMap, tt.args.url); got != tt.expected {
				t.Errorf("MergeURL() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func Test_httpProcess_respectsConfiguredTimeout(t *testing.T) {
	oldTS := types.EXPIRE_HTTP_REQUEST_TS
	defer func() { types.EXPIRE_HTTP_REQUEST_TS = oldTS }()
	types.EXPIRE_HTTP_REQUEST_TS = 10

	// slower than the previous hard-coded 3s client timeout, faster than EXPIRE_HTTP_REQUEST_TS.
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		time.Sleep(3500 * time.Millisecond)
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer srv.Close()

	req, err := http.NewRequest("GET", srv.URL, nil)
	require.NoError(t, err)

	result := map[string]bool{}
	statusCode, err := httpProcess(req, map[string]string{}, &result)
	assert.NoError(t, err)
	assert.Equal(t, 200, statusCode)
	assert.True(t, result["ok"])
}
