package schema

import (
	"encoding/json"
	"time"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/db"
)

type TwoFactor struct {
	Token  string `json:"token" bson:"token"`
	OIDCID string `json:"oidc_id" bson:"oidc_id"`
}

func Set2FA(userID bbs.UUserID, email string, token string, oidcID string, expireTSDuration time.Duration) (err error) {
	twoFactor := &TwoFactor{
		Token:  token,
		OIDCID: oidcID,
	}
	twoFactorBytes, err := json.Marshal(twoFactor)
	if err != nil {
		return err
	}
	err = db.RDBSet(rdb, RDB_PREFIX_2FA+string(userID), string(twoFactorBytes), expireTSDuration)
	if err != nil {
		return err
	}

	return nil
}

func Get2FA(userID bbs.UUserID) (twoFactor *TwoFactor, err error) {
	twoFactorStr, err := db.RDBGet(rdb, RDB_PREFIX_2FA+string(userID))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(twoFactorStr), &twoFactor)
	if err != nil {
		return nil, err
	}

	return twoFactor, nil
}

// Incr2FAAttempt increments the number of 2FA verification attempts of the user.
//
// The counter expires expireTSDuration after the first attempt.
// It is intentionally not reset by Set2FA, so requesting a new token
// does not grant more attempts.
func Incr2FAAttempt(userID bbs.UUserID, expireTSDuration time.Duration) (attempts int64, err error) {
	return db.RDBIncrWithExpire(rdb, RDB_PREFIX_2FA_ATTEMPT+string(userID), expireTSDuration)
}

// Delete2FA removes the 2FA token of the user, so the token can not be used again.
func Delete2FA(userID bbs.UUserID) (err error) {
	return db.RDBDel(rdb, RDB_PREFIX_2FA+string(userID))
}

// Reset2FAAttempt removes the 2FA attempt counter of the user.
func Reset2FAAttempt(userID bbs.UUserID) (err error) {
	return db.RDBDel(rdb, RDB_PREFIX_2FA_ATTEMPT+string(userID))
}
