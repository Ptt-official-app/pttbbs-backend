package lemmyutils

import (
	"net/http"

	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/gin-gonic/gin"
)

type LoginParams struct {
	UsernameOrEmail string `json:"username_or_email"`
	Password        string `json:"password"`
	StayLoggedIn    bool   `json:"stay_logged_in"`
}

type LoginResult struct {
	JWT                 string `json:"jwt"`
	RegistrationCreated bool   `json:"registration_created"`
	VerifyEmailSent     bool   `json:"verify_email_sent"`
}

func TryLemmyLogin(c *gin.Context) (err error) {
	nowNS := types.NowNanoTS()
	if nowNS-LEMMY_ADMIN_TOKEN_NS < LEMMY_ADMIN_TOKEN_REFRESH_NS {
		return nil
	}

	return LemmyLogin(c)
}

func LemmyLogin(c *gin.Context) (err error) {
	url := types.LEMMY_PREFIX + "/account/auth/login"

	params := &LoginParams{
		UsernameOrEmail: types.LEMMY_ADMIN,
		Password:        types.LEMMY_ADMIN_PASSWORD,
		StayLoggedIn:    true,
	}

	var ret *LoginResult
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	statusCode, err := utils.BackendPost(c, url, params, headers, &ret)
	if err != nil {
		return err
	}
	if statusCode != http.StatusOK {
		return ErrInvalidBackendStatusCode
	}

	LEMMY_ADMIN_TOKEN = ret.JWT
	LEMMY_ADMIN_TOKEN_NS = types.NowNanoTS()

	return nil
}
