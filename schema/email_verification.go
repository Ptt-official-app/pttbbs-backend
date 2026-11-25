package schema

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/db"
)

// EmailVerification is the value stored under the verification token.
//
// The stored format is the base64 of the JSON. Entries written before the
// salt was removed still carry a "Salt" field, which is ignored when reading.
type EmailVerification struct {
	Email string
}

// SetEmailVerification stores email under a random token and returns the token.
//
// The token is the only secret: it is 32 random bytes, so the Redis key can not be guessed.
func SetEmailVerification(email string, expireTSDuration time.Duration) (token string, err error) {
	tokenBytes := make([]byte, EMAIL_VERIFICATION_TOKEN_LEN)
	_, err = rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}
	token = strings.TrimRight(base64.URLEncoding.EncodeToString(tokenBytes), "=")

	emailVerification, err := serializeEmailVerification(email)
	if err != nil {
		return "", err
	}

	err = db.RDBSetNX(rdb, RDB_PREFIX_EMAIL+token, emailVerification, expireTSDuration)
	if err != nil {
		return "", err
	}

	return token, nil
}

func serializeEmailVerification(email string) (emailVerificationStr string, err error) {
	emailVerificationBytes, err := json.Marshal(&EmailVerification{Email: email})
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(emailVerificationBytes), nil
}

func deserializeEmailVerification(emailVerificationStr string) (emailVerification *EmailVerification, err error) {
	emailVerificationBytes, err := base64.URLEncoding.DecodeString(emailVerificationStr)
	if err != nil {
		return nil, err
	}

	emailVerification = &EmailVerification{}

	err = json.Unmarshal(emailVerificationBytes, emailVerification)
	if err != nil {
		return nil, err
	}

	return emailVerification, nil
}

func GetEmailVerification(token string) (emailVerification *EmailVerification, err error) {
	emailVerificationStr, err := db.RDBGet(rdb, RDB_PREFIX_EMAIL+token)
	if err != nil {
		return nil, err
	}

	emailVerification, err = deserializeEmailVerification(emailVerificationStr)
	if err != nil {
		return nil, err
	}

	return emailVerification, nil
}
