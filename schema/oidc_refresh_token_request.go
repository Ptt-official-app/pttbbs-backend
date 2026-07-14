package schema

import "time"

type OIDCRefreshTokenRequest struct {
	*OIDCRefreshToken
}

func NewOIDCRefreshTokenRequest(token *OIDCRefreshToken) (req *OIDCRefreshTokenRequest) {
	return &OIDCRefreshTokenRequest{token}
}

func (r *OIDCRefreshTokenRequest) GetAMR() []string {
	return r.AMR
}

func (r *OIDCRefreshTokenRequest) GetAudience() []string {
	return r.Audience
}

func (r *OIDCRefreshTokenRequest) GetAuthTime() time.Time {
	return r.AuthTime
}

func (r *OIDCRefreshTokenRequest) GetClientID() string {
	return r.ApplicationID
}

func (r *OIDCRefreshTokenRequest) GetScopes() []string {
	return r.Scopes
}

func (r *OIDCRefreshTokenRequest) GetSubject() string {
	return r.Subject
}

func (r *OIDCRefreshTokenRequest) SetCurrentScopes(scopes []string) {
	r.Scopes = scopes
}
