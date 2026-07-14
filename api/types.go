package api

import (
	"time"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
)

type APIFunc func(remoteAddr string, user *UserInfo, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error)

type RedirectAPIFunc func(remoteAddr string, user *UserInfo, params interface{}, c *gin.Context) (redirectPath string, statusCode int, err error)

type PathAPIFunc func(remoteAddr string, user *UserInfo, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error)

type LoginRequiredAPIFunc func(remoteAddr string, user *UserInfo, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error)

type LoginRequiredPathAPIFunc func(remoteAddr string, user *UserInfo, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error)

type LoginRequiredRedirectPathAPIFunc func(remoteAddr string, user *UserInfo, params interface{}, path interface{}, c *gin.Context) (redirectPath string, statusCode int, err error)

type errResult struct {
	Msg string

	TokenUser bbs.UUserID `json:"tokenuser"`
}

type ClientInfo struct {
	ClientID   string           `json:"c"`
	ClientType types.ClientType `json:"t"`
}

type UserInfo struct {
	UserID   bbs.UUserID
	Username string
	IsOver18 bool
}

type JwtClaim struct {
	ClientID   string           `json:"cli"`
	ClientType types.ClientType `json:"typ"`
	UUserID    string           `json:"sub"`
	Over18     bool             `json:"over18"`
	Expire     time.Time        `json:"exp"`
}
