package schema

import (
	"time"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/db"
)

func Set2FA(userID bbs.UUserID, email string, token string, expireTSDuration time.Duration) (err error) {
	err = db.RDBSetNX(rdb, RDB_PREFIX_2FA+string(userID), token, expireTSDuration)
	if err != nil {
		return err
	}

	return nil
}

func Get2FA(userID bbs.UUserID) (token string, err error) {
	token, err = db.RDBGet(rdb, RDB_PREFIX_2FA+string(userID))
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetAll2FAKeys() (keys []string, err error) {
	return db.RDBGetAllKeys(rdb, RDB_PREFIX_2FA+"*")
}
