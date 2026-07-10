package api

import (
	"path/filepath"

	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
)

type HTMLPath struct {
	Path string `uri:"path"`
}

const (
	ROOT_HTML_R  = "/"
	INDEX_HTML_R = "/index.html"
)

func IndexHTMLWrapper(c *gin.Context) {
	filename := filepath.Join(types.STATIC_DIR, "home.html")

	maxAge := htmlMaxAge()
	processCSRFContent(filename, maxAge, c)
}

const CLS_BOARDS_HTML_R = "/cls/:cls_id"

const BOARDS_HTML_R = "/boards"

const BOARDS_POPULAR_HTML_R = "/boards/popular"

const BOARDS_FAVORITES_HTML_R = "/user/:user_id/favorites"

const ARTICLES_HTML_R = "/board/:bid/articles"

const ARTICLE_HTML_R = "/board/:bid/article/:aid"

const CREATE_ARTICLE_HTML_R = "/board/:bid/post"

func AllHTMLWrapper(c *gin.Context) {
	filename := filepath.Join(types.STATIC_DIR, "all.html")

	maxAge := htmlMaxAge()
	processCSRFContent(filename, maxAge, c)
}

const LOGIN_HTML_R = "/login"

func LoginHTMLWrapper(c *gin.Context) {
	filename := filepath.Join(types.STATIC_DIR, "login.html")

	maxAge := htmlMaxAge()
	processCSRFContent(filename, maxAge, c)
}

const REGISTER_HTML_R = "/register"

func RegisterHTMLWrapper(c *gin.Context) {
	filename := filepath.Join(types.STATIC_DIR, "register.html")

	maxAge := htmlMaxAge()
	processCSRFContent(filename, maxAge, c)
}

const INIT_HTML_R = "/init"

func InitHTMLWrapper(c *gin.Context) {
	filename := filepath.Join(types.STATIC_DIR, "init.html")

	maxAge := htmlMaxAge()
	processCSRFContent(filename, maxAge, c)
}

const ERR_HTML_R = "/error"

func ErrorHTMLWrapper(c *gin.Context) {
	filename := filepath.Join(types.STATIC_DIR, "error.html")

	maxAge := htmlMaxAge()
	processCSRFContent(filename, maxAge, c)
}

const USER_HTML_R = "/user/:user_id"

func UserHTMLWrapper(c *gin.Context) {
	filename := filepath.Join(types.STATIC_DIR, "user-info.html")

	maxAge := htmlMaxAge()
	processCSRFContent(filename, maxAge, c)
}

const USER_CHANGE_PASSWD_HTML_R = "/user/:user_id/resetpassword" //nolint // passwd as route

func UserChangePasswdHTMLWrapper(c *gin.Context) {
	filename := filepath.Join(types.STATIC_DIR, "change-passwd.html")

	maxAge := htmlMaxAge()
	processCSRFContent(filename, maxAge, c)
}

const USER_ATTEMPT_CHANGE_EMAIL_HTML_R = "/user/:user_id/attemptchangeemail"

func UserAttemptChangeEmailHTMLWrapper(c *gin.Context) {
	filename := filepath.Join(types.STATIC_DIR, "attempt-change-email.html")

	maxAge := htmlMaxAge()
	processCSRFContent(filename, maxAge, c)
}

const USER_CHANGE_EMAIL_HTML_R = "/user/:user_id/changeemail"

func UserChangeEmailHTMLWrapper(c *gin.Context) {
	filename := filepath.Join(types.STATIC_DIR, "change-email.html")

	maxAge := htmlMaxAge()
	processCSRFContent(filename, maxAge, c)
}

const USER_ATTEMPT_SET_ID_EMAIL_HTML_R = "/user/:user_id/attemptsetidemail"

func UserAttemptSetIDEmailHTMLWrapper(c *gin.Context) {
	filename := filepath.Join(types.STATIC_DIR, "attempt-set-id-email.html")

	maxAge := htmlMaxAge()
	processCSRFContent(filename, maxAge, c)
}

const USER_SET_ID_EMAIL_HTML_R = "/user/:user_id/setidemail"

func UserSetIDEmailHTMLWrapper(c *gin.Context) {
	filename := filepath.Join(types.STATIC_DIR, "set-id-email.html")

	maxAge := htmlMaxAge()
	processCSRFContent(filename, maxAge, c)
}

func htmlMaxAge() (maxAge int) {
	maxAge = HTML_CACHE_CONTROL_TS

	return maxAge
}
