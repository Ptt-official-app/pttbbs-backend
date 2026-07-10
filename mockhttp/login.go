package mockhttp

import (
	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/types"

	"github.com/golang-jwt/jwt/v5"
)

func Login(params *api.LoginParams) (ret *api.LoginResult) {
	userID := params.Username
	token, _ := createToken(userID)

	ret = &api.LoginResult{
		UserID:    bbs.UUserID(userID),
		Jwt:       token,
		TokenType: "bearer",
	}

	return ret
}

func createToken(userID string) (raw string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": int(types.NowTS()) + 72*3600,
	})

	raw, err = token.SignedString(api.JWT_SECRET)
	if err != nil {
		return "", err
	}

	return raw, nil
}
