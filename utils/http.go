package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/mockhttp"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/google/go-querystring/query"
)

// BackendPost
//
// Params
//
//	postData: http-post data
//	result: resp-data, requires pointer of pointer to malloc.
//
// Ex:
//
//	url := backend.LOGIN_R
//	postData := &backend.LoginParams{}
//	result := &backend.LoginResult{}
//	BackendPost(c, url, postData, nil, &result)
func BackendPost(c *gin.Context, url string, postData interface{}, headers map[string]string, result interface{}) (statusCode int, err error) {
	if isTest {
		return mockhttp.HTTPPost(url, postData, result)
	}

	url = withBackendPrefix(url)

	if headers == nil {
		headers = make(map[string]string)
	}

	httpUpdateHeaders(headers, c)

	jsonBytes, err := json.Marshal(postData)
	if err != nil {
		return 500, err
	}

	buf := bytes.NewBuffer(jsonBytes)

	// req
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return 500, err
	}

	return httpProcess(req, headers, result)
}

func BackendGet(c *gin.Context, url string, params interface{}, headers map[string]string, result interface{}) (statusCode int, err error) {
	if isTest {
		return mockhttp.HTTPPost(url, params, result)
	}

	url = withBackendPrefix(url)

	if headers == nil {
		headers = make(map[string]string)
	}

	httpUpdateHeaders(headers, c)

	v, _ := query.Values(params)
	url = url + "?" + v.Encode()

	// req
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 500, err
	}

	return httpProcess(req, headers, result)
}

func httpUpdateHeaders(headers map[string]string, c *gin.Context) {
	if c == nil {
		headers["Content-Type"] = "application/json"
		headers["Host"] = types.HTTP_HOST
		headers["X-Forwarded-For"] = "127.0.0.1"
		return
	}

	remoteAddr := c.ClientIP()

	headers["Content-Type"] = "application/json"
	headers["Host"] = types.HTTP_HOST
	headers["X-Forwarded-For"] = remoteAddr

	authorization := c.GetHeader("Authorization")
	if authorization == "" {
		access_token := GetCookie(c, types.ACCESS_TOKEN_NAME)
		if access_token != "" {
			authorization = "Bearer " + access_token
		}
	}

	if authorization != "" {
		headers["Authorization"] = authorization
	}
}

func httpProcess(req *http.Request, headers map[string]string, result interface{}) (statusCode int, err error) {
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	ctx, cancel := context.WithTimeout(req.Context(), time.Duration(types.EXPIRE_HTTP_REQUEST_TS)*time.Second)
	defer cancel()

	req = req.WithContext(ctx)

	// send http
	//nolint:gosec // G704: httpProcess is from BackendGet and BackendPost
	resp, err := httpClient.Do(req)
	if err != nil {
		return 500, err
	}
	defer resp.Body.Close()

	statusCode = resp.StatusCode

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 501, err
	}

	if statusCode != 200 {
		errResult := &httpErrResult{}
		err = json.Unmarshal(body, errResult)
		if err != nil {
			return statusCode, errors.New(string(body))
		}
		return statusCode, errors.New(errResult.Msg)
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		return 501, err
	}

	return 200, nil
}

func withBackendPrefix(url string) string {
	return types.BACKEND_PREFIX + url
}

func MergeURL(urlMap map[string]string, url string) string {
	urlList := strings.Split(url, "/")

	newURLList := make([]string, len(urlList))
	for idx, each := range urlList {
		if len(each) == 0 || each[0] != ':' {
			newURLList[idx] = each
			continue
		}

		theKey := each[1:]
		theVal := urlMap[theKey]

		newURLList[idx] = theVal
	}

	return strings.Join(newURLList, "/")
}

func GetCookie(c *gin.Context, name string) string {
	cookie, err := c.Request.Cookie(name)
	if err != nil {
		return ""
	}
	if cookie == nil {
		return ""
	}

	return cookie.Value
}
