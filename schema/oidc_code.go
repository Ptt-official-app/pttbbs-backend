package schema

import (
	"errors"

	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/go-redis/redis/v8"
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

// ConsumeOIDCCode returns the request-id of the code and deletes the code, so that a code can be exchanged only once.
// Concurrent exchanges of the same code get ErrNotFound except for one.
func ConsumeOIDCCode(code string) (requestID string, err error) {
	requestID, err = db.RDBGet(rdb, RDB_PREFIX_OIDC_OP_CHALLENGE+code)
	if errors.Is(err, redis.Nil) {
		return "", ErrNotFound
	}
	if err != nil {
		return "", err
	}

	isDeleted, err := db.RDBDelExisting(rdb, RDB_PREFIX_OIDC_OP_CHALLENGE+code)
	if err != nil {
		return "", err
	}
	if !isDeleted {
		return "", ErrNotFound
	}

	return requestID, nil
}
