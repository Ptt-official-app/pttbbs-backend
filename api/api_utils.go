package api

import (
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

func verifyIsOver18(c *gin.Context) bool {
	return utils.GetCookie(c, types.IS_OVER_18_NAME) == types.IS_OVER_18_VALUE
}

func verifyJwt(c *gin.Context) (userID bbs.UUserID, err error) {
	token := utils.GetCookie(c, types.ACCESS_TOKEN_NAME)

	userID, _, _, err = verifyAccessToken(token)
	logrus.Infof("api.verifyJwt: after verifyAccessToken: userID: %v e: %v", userID, err)
	if err != nil {
		return "", err
	}

	// XXX skip client-info for now.
	clientInfo := &ClientInfo{
		ClientType: types.CLIENT_TYPE_APP,
	}
	/*
		err = json.Unmarshal([]byte(clientInfoStr), clientInfo)
		if err != nil {
			return "", err
		}
	*/

	if clientInfo.ClientType == types.CLIENT_TYPE_APP {
		return userID, nil
	}

	csrfToken := c.GetHeader("X-CSRFToken")
	if len(csrfToken) == 0 {
		return "", ErrInvalidToken
	}

	cookieCSRFToken := utils.GetCookie(c, types.CSRF_TOKEN)
	if cookieCSRFToken == "" {
		return "", ErrInvalidToken
	}

	if csrfToken != cookieCSRFToken {
		return "", ErrInvalidToken
	}

	if !isValidCSRFToken(csrfToken) {
		return "", ErrInvalidToken
	}

	return userID, nil
}

func createCSRFToken() (raw string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": int(types.NowTS()) + types.CSRF_TOKEN_TS,
	})

	raw, err = token.SignedString(types.CSRF_SECRET)
	if err != nil {
		return "", err
	}

	return raw, nil
}

func isValidCSRFToken(raw string) bool {
	tok, err := ParseJwt(raw, types.CSRF_SECRET)
	if err != nil {
		return false
	}

	claim, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return false
	}

	exp, err := ParseClaimInt(claim, "exp")
	if err != nil {
		return false
	}

	nowTS := int(types.NowTS())

	return nowTS <= exp
}

func processCSRFContent(filename string, cacheControlMaxAge int, c *gin.Context) {
	file, err := os.Open(filename)
	if err != nil {
		processResult(c, nil, 404, ErrFileNotFound, "")
		return
	}
	defer file.Close()

	reader := io.Reader(file)
	contentBytes, err := io.ReadAll(reader)
	if err != nil {
		processResult(c, nil, 500, ErrInvalidPath, "")
	}

	ext := filepath.Ext(filename)
	mimeType := MIME_TYPE_MAP[ext]

	content := string(contentBytes)

	csrfToken := utils.GetCookie(c, types.CSRF_TOKEN)
	if csrfToken == "" {
		csrfToken, _ = createCSRFToken()
		setCookie(c, types.CSRF_TOKEN, csrfToken, types.CSRF_TOKEN_TS_DURATION, true)
	}
	content = strings.Replace(content, "__CSRFTOKEN__", csrfToken, 1)

	c.Header("Cache-Control", "max-age="+strconv.Itoa(cacheControlMaxAge))

	processStringResult(c, content, mimeType)
}

func setCookie(c *gin.Context, name string, value string, expireDuration time.Duration, isHTTPOnly bool) {
	if c == nil || IsTest {
		return
	}

	setCookie := name + "=" + value + ";Domain=" + types.COOKIE_DOMAIN + ";Path=/;"
	if expireDuration != 0 {
		expires := time.Now().Add(expireDuration)
		expiresStr := expires.Format("Mon, Jan 2 2006 15:04:05 MST")
		setCookie += "Expires=" + expiresStr + ";"
	}
	if isHTTPOnly {
		setCookie += "HttpOnly;"
	}

	setCookie += "SameSite=Lax;" + types.TOKEN_COOKIE_SUFFIX
	c.Header("Set-Cookie", setCookie)
}

func removeCookie(c *gin.Context, name string, isHTTPOnly bool) {
	setCookie := name + "=" + ";Domain=" + types.COOKIE_DOMAIN + ";Path=/;"
	if isHTTPOnly {
		setCookie += "HttpOnly;"
	}
	c.Header("Set-Cookie", setCookie)
}
