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
