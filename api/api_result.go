package api

import (
	"net/http"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func processResult(c *gin.Context, result any, statusCode int, err error, userID bbs.UUserID) {
	setHeader(c)

	if statusCode == 200 {
		c.JSON(200, result)
		return
	}

	if statusCode == 303 {
		redirectURL, ok := result.(string)
		if !ok {
			c.JSON(500, &errResult{Msg: ErrInvalidRedirect.Error(), TokenUser: userID})
			return
		}

		c.Redirect(statusCode, redirectURL)
		return
	}

	c.JSON(statusCode, &errResult{Msg: errMsg(c, statusCode, err), TokenUser: userID})
}

// errMsg returns the error message, or the status text if err is nil or the status is 5xx.
//
// 5xx errors come from mongo / redis / go-pttbbs / the file system and may contain
// internal hosts, urls and file paths, so they are only logged.
func errMsg(c *gin.Context, statusCode int, err error) string {
	if err == nil {
		return http.StatusText(statusCode)
	}
	if statusCode >= 500 {
		logrus.Errorf("errMsg: %v %v: status: %v e: %v", c.Request.Method, c.Request.URL.Path, statusCode, err)
		return http.StatusText(statusCode)
	}
	return err.Error()
}

func processStringResult(c *gin.Context, content string, contentType string) {
	setHeader(c)

	c.Header("Content-Type", contentType)
	c.String(200, "%v", content)
}

func processRedirectResult(c *gin.Context, redirectPath string, statusCode int, err error, userID bbs.UUserID) {
	setHeader(c)

	if statusCode == 303 {
		c.Redirect(statusCode, redirectPath)
	} else {
		c.JSON(statusCode, &errResult{Msg: errMsg(c, statusCode, err), TokenUser: userID})
	}
}

func setHeader(c *gin.Context) {
	/*
		if !types.IS_ALLOW_CROSSDOMAIN {
			return
		}

		origin := c.GetHeader("Origin")

		if origin == "" {
			return
		}

		requestHeaders := c.GetHeader("Access-Control-Request-Headers")

		c.Header("X-Frame-Options", "SAMEORIGIN")
		c.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", origin)
		if requestHeaders != "" {
			c.Header("Access-Control-Allow-Headers", requestHeaders)
		}
	*/
}
