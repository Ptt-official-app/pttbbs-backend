package api

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/gin-gonic/gin"
)

const GET_USERNAME_R = "/username"

type GetUsernameResult struct {
	Username string `json:"username"`

	TokenUser bbs.UUserID `json:"tokenuser"`
}

func GetUsernameWrapper(c *gin.Context) {
	LoginRequiredQuery(GetUsername, nil, c)
}

func GetUsername(remoteAddr string, user *UserInfo, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	userID := user.UserID

	userNameInfo, err := schema.GetUserNameInfo(userID)
	if err != nil {
		return nil, 500, err
	}

	// username default to pttbbsapi.GUEST
	username := pttbbsapi.GUEST
	if userNameInfo != nil {
		username = userNameInfo.Username
	}

	return &GetUsernameResult{
		Username: username,
	}, 200, nil
}
