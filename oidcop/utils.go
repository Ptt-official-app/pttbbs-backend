package oidcop

import (
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/zitadel/oidc/v3/pkg/op"
)

// getInfoFromRequest returns the clientID, authTime and amr depending on the op.TokenRequest type / implementation
func getInfoFromRequest(req op.TokenRequest) (clientID string, authTime time.Time, amr []string) {
	authReq, ok := req.(*schema.OIDCAuthRequest) // Code Flow (with scope offline_access)
	if ok {
		return authReq.ClientID, authReq.AuthTime, authReq.GetAMR()
	}

	refreshReq, ok := req.(*schema.OIDCRefreshTokenRequest) // Refresh Token Request
	if ok {
		return refreshReq.ApplicationID, refreshReq.AuthTime, refreshReq.AMR
	}

	return "", time.Time{}, nil
}
