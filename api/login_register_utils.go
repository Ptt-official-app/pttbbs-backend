package api

import (
	"fmt"
	"strings"
	"time"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

func setTokenToCookie(c *gin.Context, accessToken string) {
	setCookie(c, types.ACCESS_TOKEN_NAME, accessToken, types.ACCESS_TOKEN_EXPIRE_TS_DURATION, true)
}

func removeTokenFromCookie(c *gin.Context) {
	removeCookie(c, types.ACCESS_TOKEN_NAME, true)
}

func gen2FATokenAndSendEmail(userID bbs.UUserID, username string, email string, title string, template string, expireTS time.Duration) (err error) {
	token := gen2FAToken()

	err = schema.Set2FA(userID, email, token, expireTS)
	if err != nil {
		return err
	}

	logrus.Infof("gen2FATokenAndSendEmail: userID: %v username %v email: %v", userID, username, email)

	content := strings.ReplaceAll(
		strings.ReplaceAll(
			template, types.TEMPLATE_USER, username,
		), types.TEMPLATE_TOKEN, token,
	)

	return utils.SendEmail([]string{email}, title, content)
}

func gen2FAToken() string {
	randInt := utils.GenRandomInt64(types.MAX_2FA_TOKEN)
	return fmt.Sprintf(types.MAX_2FA_TOKEN_STR_PROMPT, randInt)
}

func check2FAToken(userID bbs.UUserID, token string) (err error) {
	token_db, err := schema.Get2FA(userID)
	if err != nil {
		return err
	}

	if token != token_db {
		return ErrInvalidToken
	}

	return nil
}

func genEmailVerificationTokenAndSendEmail(email string, title string, url string, template string) (err error) {
	token, err := schema.SetEmailVerification(email, types.EXPIRE_ATTEMPT_REGISTER_USER_EMAIL_TS_DURATION)
	if err != nil {
		return err
	}

	theUrl := url + "?token=" + token

	content := replaceEmailVerificationTemplate(template, "", email, theUrl)

	return utils.SendEmail([]string{email}, title, content)
}

func getEmailFromEmailVerificationToken(token string) (email string, err error) {
	emailtoken_db, err := schema.GetEmailVerification(token)
	if err != nil {
		return "", err
	}

	return emailtoken_db.Email, nil
}

func replaceEmailVerificationTemplate(template string, user string, email string, url string) string {
	return strings.ReplaceAll(strings.ReplaceAll(
		strings.ReplaceAll(
			template, types.TEMPLATE_EMAIL, email,
		), types.TEMPLATE_URL, url,
	), types.TEMPLATE_USER, user)
}

func genUserID(email string) (userID bbs.UUserID, err error) {
	updateNanoTS := types.NowNanoTS()

	userIDStr := ""
	for theLoop := range 3 {
		for range 3 {
			userIDStr = ""
			for range theLoop + 1 {
				userIDStr += utils.GenRandomString()
			}
			userID = bbs.UUserID(userIDStr)
			err = schema.CreateUserEmail(userID, email, true, updateNanoTS)
			if err == nil {
				return userID, nil
			}
		}
	}
	return "", ErrNoUserID
}

func genUsername(userID bbs.UUserID) (username string, err error) {
	updateNanoTS := types.NowNanoTS()

	theRandom := ""
	for theLoop := range 3 {
		for range 3 {
			theRandom = ""
			for range theLoop + 1 {
				theRandom += utils.GenRandomString()[:INIT_USERNAME_RANDOM_LEN+theLoop]
			}
			username = INIT_USERNAME_PREFIX + "." + theRandom
			err = schema.CreateUsername(userID, username, updateNanoTS)
			if err == nil {
				return username, nil
			}
		}
	}

	return "", ErrNoUsername
}

func genAccessToken(userID bbs.UUserID, clientInfo string) (token string, tokenExpireTS types.Time8, err error) {
	return genToken(userID, clientInfo, types.ACCESS_TOKEN_EXPIRE_TS, types.ACCESS_TOKEN_SECRET)
}

func genRefreshToken(userID bbs.UUserID, clientInfo string) (token string, tokenExpireTS types.Time8, err error) {
	return genToken(userID, clientInfo, types.REFRESH_TOKEN_EXPIRE_TS, types.REFRESH_TOKEN_SECRET)
}

func genToken(userID bbs.UUserID, clientInfo string, expireTS int, secret []byte) (token string, tokenExpireTS types.Time8, err error) {
	defer func() {
		err2 := recover()
		if err2 == nil {
			return
		}

		err = types.ErrRecover(err2)
	}()

	tokenExpireTS = types.NowTS() + types.Time8(expireTS)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"cli": clientInfo,
		"sub": userID,
		"exp": int(tokenExpireTS),
	})

	token, err = jwtToken.SignedString(secret)
	if err != nil {
		return "", 0, err
	}

	return token, tokenExpireTS, nil
}

func verifyAccessToken(token string) (userID bbs.UUserID, expireTS int, clientInfo string, err error) {
	defer func() {
		err2 := recover()
		if err2 == nil {
			return
		}

		err = types.ErrRecover(err2)
	}()

	if token == "" {
		return bbs.UUserID(pttbbsapi.GUEST), 0, "", nil
	}

	claim, err := parseJwtClaim(token, types.ACCESS_TOKEN_SECRET)
	if err != nil {
		return "", 0, "", ErrInvalidToken
	}

	currentTS := int(types.NowTS())
	if currentTS > claim.Expire {
		return "", 0, "", ErrInvalidToken
	}

	return bbs.UUserID(claim.UUserID), claim.Expire, claim.ClientInfo, nil
}

func parseJwtClaim(token string, secret []byte) (cl *JwtClaim, err error) {
	tok, err := ParseJwt(token, secret)
	if err != nil {
		return nil, err
	}

	claim, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	cli, err := ParseClaimString(claim, "cli")
	if err != nil {
		return nil, err
	}
	sub, err := ParseClaimString(claim, "sub")
	if err != nil {
		return nil, err
	}
	exp, err := ParseClaimInt(claim, "exp")
	if err != nil {
		return nil, err
	}

	cl = &JwtClaim{
		ClientInfo: cli,
		UUserID:    sub,
		Expire:     exp,
	}

	return cl, nil
}

func loginInputToUsernameEmail(input string) (userID bbs.UUserID, username string, email string, err error) {
	email = input
	username = input
	var userEmail *schema.UserEmail
	if !strings.Contains(email, "@") { // username, not email
		email, userEmail, err = loginGetEmailByUsername(username)
		if err != nil {
			return "", "", "", err
		}
		if userEmail == nil {
			return "", "", "", ErrNoEmail
		}
		userID = userEmail.UserID
	} else {
		username, userEmail, err = loginGetUsernameByEmail(email)
		if err != nil {
			return "", "", "", err
		}
		if userEmail == nil {
			return "", "", "", ErrNoEmail
		}
		userID = userEmail.UserID
	}

	return userID, username, email, nil
}

func loginGetEmailByUsername(username string) (email string, userEmail *schema.UserEmail, err error) {
	userUsername, err := schema.GetUserUsernameByUsername(username)
	if err != nil {
		return "", nil, err
	}
	if userUsername == nil {
		logrus.Errorf("loginGetEmailByUsername: unable to get userUsername (nil): username: %v", username)
		return "", nil, ErrNoUsername
	}
	userEmail, err = schema.GetUserEmailByUserID(userUsername.UserID)
	if err != nil {
		return "", nil, err
	}
	if userEmail == nil {
		return "", nil, ErrNoEmail
	}

	return userEmail.Email, userEmail, nil
}

func loginGetUsernameByEmail(email string) (username string, userEmail *schema.UserEmail, err error) {
	userEmail, err = schema.GetUserEmailByEmail(email)
	if err != nil {
		return "", nil, err
	}
	if userEmail == nil {
		return "", nil, ErrNoEmail
	}
	userUsername, err := schema.GetUserUsernameByUserID(userEmail.UserID)
	if err != nil {
		return "", nil, err
	}
	if userUsername == nil {
		logrus.Errorf("loginGetUsernameByEmail: unable to get userUsername (nil): username: %v", username)
		return "", nil, ErrNoUsername
	}

	return userUsername.Username, userEmail, nil
}
