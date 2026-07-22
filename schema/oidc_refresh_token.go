package schema

import (
	"encoding/json"
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

type OIDCRefreshToken struct {
	ID            string    `json:"id" bson:"id"`
	Token         string    `json:"token" bson:"token"`
	AuthTime      time.Time `json:"auth_time" bson:"auth_time"`
	AMR           []string  `json:"amr" bson:"amr"`
	Audience      []string  `json:"aud" bson:"aud"`
	Subject       string    `json:"sub" bson:"sub"`
	ApplicationID string    `json:"appid" bson:"appid"`
	Expiration    time.Time `json:"exp" bson:"exp"`
	Scopes        []string  `json:"scope" bson:"scope"`
	AccessToken   string    `json:"access" bson:"access"`
}

func GetOIDCRefreshToken(refreshTokenID string) (refreshToken *OIDCRefreshToken, err error) {
	val, err := db.RDBGet(rdb, RDB_PREFIX_OIDC_OP_REFRESH_TOKEN+refreshTokenID)
	if err != nil {
		return nil, err
	}
	if val == "" {
		return nil, ErrNotFound
	}

	refreshToken = &OIDCRefreshToken{}
	err = json.Unmarshal([]byte(val), refreshToken)
	if err != nil {
		return nil, err
	}

	return refreshToken, nil
}

func SetOIDCRefreshToken(refreshToken *OIDCRefreshToken) (err error) {
	//nolint:gosec // marshal to internal db.
	val, err := json.Marshal(refreshToken)
	if err != nil {
		return err
	}

	return db.RDBSet(rdb, RDB_PREFIX_OIDC_OP_REFRESH_TOKEN+refreshToken.ID, string(val), types.REFRESH_TOKEN_EXPIRE_TS_DURATION)
}

func DeleteOIDCRefreshToken(refreshTokenID string) (err error) {
	return db.RDBDel(rdb, RDB_PREFIX_OIDC_OP_REFRESH_TOKEN+refreshTokenID)
}
