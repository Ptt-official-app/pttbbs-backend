package schema

import (
	"crypto/ed25519"

	"github.com/Ptt-official-app/pttbbs-backend/db"
	"go.mongodb.org/mongo-driver/bson"
)

var OIDCKey_c *db.Collection

// OIDCKey is for client-id with the possibility of multiple keys.
// A very typical scenario is to transition the expiring keys.
type OIDCKey struct {
	ClientID string             `json:"client_id" bson:"client_id"`
	KeyID    string             `json:"key_id" bson:"key_id"`
	Key      *ed25519.PublicKey `json:"key" bson:"key"`
}

var EMPTY_OIDC_KEY = &OIDCKey{}

var (
	OIDC_KEY_CLIENT_ID_b = getBSONName(EMPTY_OIDC_KEY, "ClientID")
	OIDC_KEY_KEY_ID_b    = getBSONName(EMPTY_OIDC_KEY, "KeyID")
	OIDC_KEY_KEY_b       = getBSONName(EMPTY_OIDC_KEY, "Key")
)

func GetOIDCKey(clientID string, keyID string) (key *OIDCKey, err error) {
	query := bson.M{
		OIDC_KEY_CLIENT_ID_b: clientID,
		OIDC_KEY_KEY_ID_b:    keyID,
	}

	key = &OIDCKey{}

	err = OIDCKey_c.FindOne(query, key, nil)
	if err != nil {
		return nil, err
	}

	return key, nil
}
