package schema

import (
	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

type OIDCCode struct {
	Code      string `bson:"code"`
	RequestID string `bson:"req"`
}

func SaveOIDCCode(requestID string, code string) (err error) {
	err = db.RDBSetNX(rdb, RDB_PREFIX_OIDC_OP_CHALLENGE+code, requestID, types.EXPIRE_OIDC_AUTH_REQUEST_TS_DURATION)
	if err != nil {
		return err
	}

	return nil
}

func DeleteOIDCCodeByCode(code string) (err error) {
	return db.RDBDel(rdb, RDB_PREFIX_OIDC_OP_CHALLENGE+code)
}

func GetRequestIDByOIDCCode(code string) (requestID string, err error) {
	requestID, err = db.RDBGet(rdb, RDB_PREFIX_OIDC_OP_CHALLENGE+code)
	if err != nil {
		return "", err
	}

	return requestID, nil
}
