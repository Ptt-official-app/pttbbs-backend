package api

import (
	"context"
	"fmt"
	"strings"
	"time"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/oidcop"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zitadel/oidc/v3/pkg/oidc"
	"github.com/zitadel/oidc/v3/pkg/op"
	"go.mongodb.org/mongo-driver/mongo"
)

func setTokenToCookie(c *gin.Context, accessToken string, refreshToken string) {
	setCookie(c, types.ACCESS_TOKEN_NAME, accessToken, types.ACCESS_TOKEN_EXPIRE_TS_DURATION, true, "")
	setCookie(c, types.REFRESH_TOKEN_NAME, refreshToken, types.REFRESH_TOKEN_EXPIRE_TS_DURATION, true, "")
}

func removeTokenFromCookie(c *gin.Context) {
	removeCookie(c, types.ACCESS_TOKEN_NAME, true)
	removeCookie(c, types.REFRESH_TOKEN_NAME, true)
}

func gen2FATokenAndSendEmail(userID bbs.UUserID, username string, email string, oidcID string, title string, template string, expireTS time.Duration) (err error) {
	token := gen2FAToken()

	err = schema.Set2FA(userID, email, token, oidcID, expireTS)
	if err != nil {
		return err
	}

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

func check2FAToken(userID bbs.UUserID, token string) (oidcID string, err error) {
	twoFA_db, err := schema.Get2FA(userID)
	if err != nil {
		return "", err
	}

	if token != twoFA_db.Token {
		return "", ErrInvalidToken
	}

	return twoFA_db.OIDCID, nil
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

func genAccessAndRefreshTokens(ctx context.Context, username string, clientID string, refreshToken string) (accessToken string, newRefreshToken string, validity time.Duration, err error) {
	client, err := schema.GetClient(clientID)
	if err != nil {
		logrus.Warnf("api.genAccessAndRefreshTokens: unable to GetClient: clientID: %v e: %v", clientID, err)
		return "", "", 0, err
	}
	if client == nil {
		logrus.Warnf("api.genAccessAndRefreshTokens: no client: clientID: %v", clientID)
		return "", "", 0, mongo.ErrNoDocuments
	}

	authRequest := schema.NewAuthRequestByUsername(username, clientID)
	ctx = op.ContextWithIssuer(ctx, types.OIDC_OP_ISSUER)

	return op.CreateAccessToken(ctx, authRequest, op.AccessTokenTypeJWT, oidcop.PROVIDER, client, refreshToken)
}

func verifyAccessToken(token string) (userID bbs.UUserID, expireTS int, newClientID string, clientType types.ClientType, isOver18 bool, err error) {
	defer func() {
		err2 := recover()
		if err2 == nil {
			return
		}

		err = types.ErrRecover(err2)
	}()

	if token == "" {
		return bbs.UUserID(pttbbsapi.GUEST), 0, "", types.CLIENT_TYPE_WEB, false, nil
	}

	ctxTodo := context.TODO()
	ctx := op.ContextWithIssuer(ctxTodo, types.OIDC_OP_ISSUER)

	claim, err := op.VerifyAccessToken[*oidc.AccessTokenClaims](ctx, token, oidcop.PROVIDER.AccessTokenVerifier(ctx))
	if err != nil {
		return "", 0, "", types.CLIENT_TYPE_WEB, false, ErrInvalidToken
	}

	username := claim.Subject
	userUsernameOver18, err := schema.GetUserUsernameOver18ByUsername(username)
	if err != nil {
		return "", 0, "", types.CLIENT_TYPE_WEB, false, err
	}

	expireTS = int(claim.Expiration.AsTime().Unix())

	clientType = types.CLIENT_TYPE_APP
	if claim.ClientID == types.WEB_CLIENT_ID {
		clientType = types.CLIENT_TYPE_WEB
	}

	return userUsernameOver18.UserID, expireTS, claim.ClientID, clientType, userUsernameOver18.Over18, nil
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
		logrus.Errorf("loginGetEmailByUsername: unable to get userUsername: username: %v e: %v", username, err)
		return "", nil, err
	}
	if userUsername == nil {
		logrus.Errorf("loginGetEmailByUsername: unable to get userUsername (nil): username: %v", username)
		return "", nil, ErrNoUsername
	}
	userEmail, err = schema.GetUserEmailByUserID(userUsername.UserID)
	if err != nil {
		logrus.Errorf("loginGetEmailByUsername: unable to get userEmail: username: %v userID: %v e: %v", username, userUsername.UserID, err)
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
