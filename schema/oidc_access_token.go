package schema

import (
	"encoding/json"
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

type OIDCAccessToken struct {
	ID             string    `json:"id" bson:"id"`
	ApplicationID  string    `json:"appid" bson:"appid"`
	Subject        string    `json:"sub" bson:"sub"`
	RefreshTokenID string    `json:"refresh" bson:"refresh"`
	Audience       []string  `json:"aud" bson:"aud"`
	Expiration     time.Time `json:"exp" bson:"exp"`
	Scopes         []string  `json:"scopes" bson:"scopes"`
}

func GetOIDCAccessTokensBySubjectAndClientID(subject string, clientID string) (accessTokens []*OIDCAccessToken, err error) {
	subjectClientID := serializeOIDCAccessTokenSubjectClientID(subject, clientID)

	accessTokenIDs, err := db.RDBSMembers(rdb, RDB_PREFIX_OIDC_OP_ACCESS_TOKEN_SUBJECT_CLIENTID+subjectClientID)
	if err != nil {
		return nil, err
	}

	accessTokens = make([]*OIDCAccessToken, 0, len(accessTokenIDs))
	for _, eachID := range accessTokenIDs {
		each, err := GetOIDCAccessToken(eachID)
		if err == ErrNotFound {
			continue
		}
		if err != nil {
			return nil, err
		}
		if each == nil {
			continue
		}
		accessTokens = append(accessTokens, each)
	}

	return accessTokens, nil
}

func GetOIDCAccessToken(accessTokenID string) (accessToken *OIDCAccessToken, err error) {
	val, err := db.RDBGet(rdb, RDB_PREFIX_OIDC_OP_ACCESS_TOKEN+accessTokenID)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(val), &accessToken)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func SetOIDCAccessToken(accessToken *OIDCAccessToken) (err error) {
	val, err := json.Marshal(accessToken)
	if err != nil {
		return nil
	}

	subjectClientID := serializeOIDCAccessTokenSubjectClientID(accessToken.Subject, accessToken.ApplicationID)

	err = db.RDBSAdd(rdb, RDB_PREFIX_OIDC_OP_ACCESS_TOKEN_SUBJECT_CLIENTID+subjectClientID, accessToken.ID, types.ACCESS_TOKEN_EXPIRE_TS_DURATION+1*time.Second)
	if err != nil {
		return err
	}
	err = db.RDBSetNX(rdb, RDB_PREFIX_OIDC_OP_ACCESS_TOKEN+accessToken.ID, string(val), types.ACCESS_TOKEN_EXPIRE_TS_DURATION)
	if err != nil {
		return err
	}

	return nil
}

func DeleteOIDCAccessToken(accessTokenID string) (err error) {
	accessToken, err := GetOIDCAccessToken(accessTokenID)
	if err != nil {
		return err
	}

	subjectClientID := serializeOIDCAccessTokenSubjectClientID(accessToken.Subject, accessToken.ApplicationID)

	err = db.RDBDel(rdb, RDB_PREFIX_OIDC_OP_ACCESS_TOKEN+accessToken.ID)
	if err != nil {
		return err
	}

	return db.RDBSRem(rdb, RDB_PREFIX_OIDC_OP_ACCESS_TOKEN_SUBJECT_CLIENTID+subjectClientID, accessTokenID)
}

func serializeOIDCAccessTokenSubjectClientID(subject string, clientID string) string {
	return subject + RDB_PREFIX_SEPARATOR + clientID
}
