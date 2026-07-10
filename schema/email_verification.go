package schema

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/db"
	"golang.org/x/crypto/argon2"
)

type EmailSalt struct {
	Email string
	Salt  []byte
}

func SetEmailVerification(email string, expireTSDuration time.Duration) (value string, err error) {
	token := make([]byte, EMAIL_VERIFICATION_TOKEN_LEN)
	_, err = rand.Read(token)
	if err != nil {
		return "", err
	}

	value, salt, err := emailVerificationSerializeValue(email, token)
	if err != nil {
		return "", err
	}

	emailSalt, err := serializeEmailSalt(email, salt)
	if err != nil {
		return "", err
	}

	err = db.RDBSetNX(rdb, RDB_PREFIX_EMAIL+value, emailSalt, expireTSDuration)
	if err != nil {
		return "", err
	}

	return value, nil
}

func serializeEmailSalt(email string, salt []byte) (emailSaltStr string, err error) {
	emailSalt := &EmailSalt{
		Email: email,
		Salt:  salt,
	}

	emailSaltBytes, err := json.Marshal(emailSalt)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(emailSaltBytes), nil
}

func deserializeEmailSalt(emailSaltStr string) (emailSalt *EmailSalt, err error) {
	emailSaltBytes, err := base64.URLEncoding.DecodeString(emailSaltStr)
	if err != nil {
		return nil, err
	}

	emailSalt = &EmailSalt{}

	err = json.Unmarshal(emailSaltBytes, emailSalt)
	if err != nil {
		return nil, err
	}

	return emailSalt, nil
}

func GetEmailVerification(tokenEmail string) (emailSalt *EmailSalt, err error) {
	emailSaltStr, err := db.RDBGet(rdb, RDB_PREFIX_EMAIL+tokenEmail)
	if err != nil {
		return nil, err
	}

	emailSalt, err = deserializeEmailSalt(emailSaltStr)
	if err != nil {
		return nil, err
	}

	return emailSalt, nil
}

func VerifyEmailVerification(tokenEmail string, emailSalt *EmailSalt) (isValid bool, err error) {
	_, hashedEmailBytes, err := emailVerificationDeserializeValue(tokenEmail)
	if err != nil {
		return false, err
	}

	email := emailSalt.Email
	salt := emailSalt.Salt

	emailBytes := []byte(email)
	hashedEmailBytes2 := argon2.IDKey(emailBytes, salt, ARGON2_TIME, ARGON2_MEMORY, ARGON2_THREADS, ARGON2_KEYLEN)
	return bytes.Equal(hashedEmailBytes, hashedEmailBytes2), nil
}

func emailVerificationSerializeValue(email string, token []byte) (value string, salt []byte, err error) {
	salt = make([]byte, ARGON2_SALT_LEN)
	_, err = rand.Read(salt)
	if err != nil {
		return "", nil, err
	}

	emailBytes := []byte(email)
	hashedEmailBytes := argon2.IDKey(emailBytes, salt, ARGON2_TIME, ARGON2_MEMORY, ARGON2_THREADS, ARGON2_KEYLEN)

	tokenEmailBytes := append(token, hashedEmailBytes...)

	tokenEmail := strings.TrimRight(base64.URLEncoding.EncodeToString(tokenEmailBytes), "=")

	return tokenEmail, salt, nil
}

func emailVerificationDeserializeValue(tokenEmail string) (token []byte, hashedEmailBytes []byte, err error) {
	tokenEmailWithPadding := b64WithPadding(tokenEmail)
	tokenEmailBytes, err := base64.URLEncoding.DecodeString(tokenEmailWithPadding)
	if err != nil {
		return nil, nil, err
	}
	token = tokenEmailBytes[:EMAIL_VERIFICATION_TOKEN_LEN]
	hashedEmailBytes = tokenEmailBytes[EMAIL_VERIFICATION_TOKEN_LEN:]

	return token, hashedEmailBytes, nil
}

func b64WithPadding(theStr string) string {
	nPadding := (4 - len(theStr)%4) % 4
	padding := strings.Repeat("=", nPadding)
	return theStr + padding
}
