package api

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func verifyIsOver18(c *gin.Context) bool {
	return utils.GetCookie(c, types.IS_OVER_18_NAME) == types.IS_OVER_18_VALUE
}

// verifyJwt verifies the access token (and the csrf-token for web clients).
//
//   - no access token: GUEST, no error.
//   - invalid / expired access token: ErrInvalidToken.
//   - valid access token but invalid csrf-token: ErrInvalidCSRFToken.
func verifyJwt(c *gin.Context) (userID bbs.UUserID, isOver18 bool, err error) {
	token := utils.GetAccessToken(c)
	if token == "" {
		return bbs.UUserID(pttbbsapi.GUEST), false, nil
	}

	userID, _, _, clientType, isOver18, err := verifyAccessToken(token)
	if err != nil {
		return "", false, err
	}

	if clientType == types.CLIENT_TYPE_APP {
		return userID, isOver18, nil
	}

	if types.SERVICE_MODE == types.DEV { // no checking X-CSRFToken in dev mode.
		return userID, isOver18, nil
	}

	csrfToken := c.GetHeader("X-CSRFToken")
	if len(csrfToken) == 0 {
		return "", false, ErrInvalidCSRFToken
	}

	cookieCSRFToken := utils.GetCookie(c, types.CSRF_TOKEN)
	if cookieCSRFToken == "" {
		return "", false, ErrInvalidCSRFToken
	}

	if csrfToken != cookieCSRFToken {
		return "", false, ErrInvalidCSRFToken
	}

	if !isValidCSRFToken(csrfToken) {
		return "", false, ErrInvalidCSRFToken
	}

	return userID, isOver18, nil
}

// getLoginUser returns the user of the request for login-required apis.
//
// The request is processed as GUEST if there is no access token, or the csrf-token is invalid.
// It returns 401 if the access token is invalid or expired,
// so the client knows that the user is no longer logged-in,
// instead of silently being processed as GUEST.
func getLoginUser(c *gin.Context) (userID bbs.UUserID, isOver18 bool, statusCode int, err error) {
	userID, isOver18, err = verifyJwt(c)
	switch {
	case err == nil:
		return userID, isOver18, 200, nil
	case errors.Is(err, ErrInvalidCSRFToken):
		return bbs.UUserID(pttbbsapi.GUEST), false, 200, nil
	case errors.Is(err, ErrInvalidToken):
		return "", false, 401, err
	default:
		return "", false, 500, err
	}
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
		return
	}

	ext := filepath.Ext(filename)
	mimeType := MIME_TYPE_MAP[ext]

	content := string(contentBytes)

	// The cookie is client-controlled: reuse it only if it is a valid (signed and not expired) token.
	// Otherwise an expired token would be embedded in the page, and an arbitrary cookie value
	// would be reflected into the HTML.
	csrfToken := utils.GetCookie(c, types.CSRF_TOKEN)
	if !isValidCSRFToken(csrfToken) {
		csrfToken, err = createCSRFToken()
		if err != nil {
			processResult(c, nil, 500, err, "")
			return
		}
		setCookie(c, types.CSRF_TOKEN, csrfToken, types.CSRF_TOKEN_TS_DURATION, true, types.CSRF_COOKIE_DOMAIN)
	}
	content = strings.Replace(content, "__CSRFTOKEN__", csrfToken, 1)

	// The content contains the per-user csrf-token, it must not be stored by shared caches.
	c.Header("Cache-Control", "private, max-age="+strconv.Itoa(cacheControlMaxAge))

	processStringResult(c, content, mimeType)
}

func setCookie(c *gin.Context, name string, value string, expireDuration time.Duration, isHTTPOnly bool, cookieDomain string) {
	if c == nil || IsTest {
		return
	}

	if cookieDomain == "" {
		cookieDomain = types.COOKIE_DOMAIN
	}

	// Lax: browsers do not send the session cookies on cross-site POSTs (e.g. forged logout).
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(name, value, int(expireDuration.Seconds()), "/", cookieDomain, true, isHTTPOnly)
}

func removeCookie(c *gin.Context, name string, isHTTPOnly bool, cookieDomain string) {
	if c == nil {
		return
	}

	if cookieDomain == "" {
		cookieDomain = types.COOKIE_DOMAIN
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(name, "", -1, "/", cookieDomain, true, isHTTPOnly)
}
