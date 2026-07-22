package schema

import (
	"encoding/json"
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/zitadel/oidc/v3/pkg/oidc"
	"golang.org/x/text/language"
)

type OIDCAuthRequest struct {
	ID       string `json:"id" bson:"id"`
	ClientID string `json:"client_id" bson:"client_id"`

	CreationDate time.Time `json:"create_time" bson:"create_time"`

	Subject  string    `json:"sub" bson:"sub"`
	ACR      string    `json:"acr" bson:"acr"`
	AMR      []string  `json:"amr" bson:"amr"`
	Audience []string  `json:"aud" bson:"aud"`
	AuthTime time.Time `json:"auth_time" bson:"auth_time"`

	Scopes []string `json:"scope" bson:"scope"`

	State string `json:"state" bson:"state"`

	CodeChallenge *oidc.CodeChallenge `json:"challenges" bson:"challenges"`
	Nonce         string              `json:"nonce" bson:"nonce"`

	RedirectURI string `json:"uri" bson:"uri"`

	ResponseMode oidc.ResponseMode `json:"mode" bson:"mode"`
	ResponseType oidc.ResponseType `json:"type" bson:"type"`

	IsAuth bool `json:"done" bson:"done"`

	Prompt     []string       `json:"prompt" bson:"prompt"`
	UiLocales  []language.Tag `json:"locale" bson:"locale"`
	LoginHint  string         `json:"login_hint" bson:"login_hint"`
	MaxAuthAge time.Duration  `json:"max_auth_age" bson:"max_auth_age"`
	Username   string         `json:"username" bson:"username"`
}

type OIDCAuthRequestIsAuthInfo struct {
	Subject  string    `json:"sub" bson:"sub"`
	AuthTime time.Time `json:"auth_time" bson:"auth_time"`
}

func NewAuthRequestByUsername(username string, clientID string) *OIDCAuthRequest {
	return &OIDCAuthRequest{
		CreationDate: time.Now(),
		ClientID:     clientID,
		Subject:      username,
		Username:     username,
		Scopes:       []string{"openid", "offline_access"},
	}
}

func NewAuthRequest(authReq *oidc.AuthRequest, username string) *OIDCAuthRequest {
	var codeChallenge *oidc.CodeChallenge
	if authReq.CodeChallenge != "" {
		codeChallenge = &oidc.CodeChallenge{
			Challenge: authReq.CodeChallenge,
			Method:    authReq.CodeChallengeMethod,
		}
	}

	return &OIDCAuthRequest{
		CreationDate:  time.Now(),
		ClientID:      authReq.ClientID,
		RedirectURI:   authReq.RedirectURI,
		State:         authReq.State,
		Prompt:        PromptToInternal(authReq.Prompt),
		UiLocales:     authReq.UILocales,
		LoginHint:     authReq.LoginHint,
		MaxAuthAge:    MaxAgeToInternal(authReq.MaxAge),
		Subject:       username,
		Username:      username,
		Scopes:        authReq.Scopes,
		ResponseType:  authReq.ResponseType,
		ResponseMode:  authReq.ResponseMode,
		Nonce:         authReq.Nonce,
		CodeChallenge: codeChallenge,
	}
}

func SaveAuthRequest(authReq *OIDCAuthRequest) (err error) {
	authReqBytes, err := json.Marshal(authReq)
	if err != nil {
		return err
	}

	// ensure that the challenge is unique.
	err = db.RDBSetNX(rdb, RDB_PREFIX_OIDC_OP_REQUEST_CHALLENGE+authReq.CodeChallenge.Challenge, RDB_OIDC_OP_REQUEST_TRUE, types.EXPIRE_OIDC_AUTH_REQUEST_TS_DURATION+1*time.Second)
	if err != nil {
		return err
	}

	err = db.RDBSetNX(rdb, RDB_PREFIX_OIDC_OP_REQUEST+authReq.ID, string(authReqBytes), types.EXPIRE_OIDC_AUTH_REQUEST_TS_DURATION)
	if err != nil {
		return err
	}
	return nil
}

func SetAuthRequestIsAuth(authReqID string, username string) (err error) {
	info := &OIDCAuthRequestIsAuthInfo{
		Subject:  username,
		AuthTime: time.Now(),
	}
	infoStr, err := json.Marshal(info)
	if err != nil {
		return err
	}
	err = db.RDBSet(rdb, RDB_PREFIX_OIDC_OP_REQUEST_IS_AUTH+authReqID, string(infoStr), types.EXPIRE_OIDC_AUTH_REQUEST_TS_DURATION)
	if err != nil {
		return err
	}

	return nil
}

func PromptToInternal(oidcPrompt oidc.SpaceDelimitedArray) []string {
	prompts := make([]string, 0, len(oidcPrompt))
	for _, oidcPrompt := range oidcPrompt {
		switch oidcPrompt {
		case oidc.PromptNone,
			oidc.PromptLogin,
			oidc.PromptConsent,
			oidc.PromptSelectAccount:
			prompts = append(prompts, oidcPrompt)
		}
	}
	return prompts
}

func MaxAgeToInternal(maxAge *uint) time.Duration {
	if maxAge == nil {
		return time.Duration(0)
	}
	dur := time.Duration(*maxAge) * time.Second
	return dur
}

// Done implements [op.AuthRequest].
func (o *OIDCAuthRequest) Done() bool {
	if o.IsAuth {
		return o.IsAuth
	}

	err := o.fillAuthInfo()
	if err != nil {
		return false
	}

	return o.IsAuth
}

func (o *OIDCAuthRequest) fillAuthInfo() (err error) {
	val, err := db.RDBGet(rdb, RDB_PREFIX_OIDC_OP_REQUEST_IS_AUTH+o.ID)
	if err != nil {
		return err
	}
	if val == "" {
		return err
	}

	var info *OIDCAuthRequestIsAuthInfo
	err = json.Unmarshal([]byte(val), &info)
	if err != nil {
		return err
	}

	// cache o.IsAuth
	o.IsAuth = true
	o.Subject = info.Subject
	o.AuthTime = info.AuthTime

	return nil
}

// GetACR implements [op.AuthRequest].
func (o *OIDCAuthRequest) GetACR() string {
	return o.ACR
}

// GetAMR implements [op.AuthRequest].
func (o *OIDCAuthRequest) GetAMR() []string {
	return o.AMR
}

// GetAudience implements [op.AuthRequest].
func (o *OIDCAuthRequest) GetAudience() []string {
	return o.Audience
}

// GetAuthTime implements [op.AuthRequest].
func (o *OIDCAuthRequest) GetAuthTime() time.Time {
	if !o.AuthTime.IsZero() {
		return o.AuthTime
	}
	err := o.fillAuthInfo()
	if err != nil {
		return time.Time{}
	}
	return o.AuthTime
}

// GetClientID implements [op.AuthRequest].
func (o *OIDCAuthRequest) GetClientID() string {
	return o.ClientID
}

// GetCodeChallenge implements [op.AuthRequest].
func (o *OIDCAuthRequest) GetCodeChallenge() *oidc.CodeChallenge {
	return o.CodeChallenge
}

// GetID implements [op.AuthRequest].
func (o *OIDCAuthRequest) GetID() string {
	return o.ID
}

// GetNonce implements [op.AuthRequest].
func (o *OIDCAuthRequest) GetNonce() string {
	return o.Nonce
}

// GetRedirectURI implements [op.AuthRequest].
func (o *OIDCAuthRequest) GetRedirectURI() string {
	return o.RedirectURI
}

// GetResponseMode implements [op.AuthRequest].
func (o *OIDCAuthRequest) GetResponseMode() oidc.ResponseMode {
	return o.ResponseMode
}

// GetResponseType implements [op.AuthRequest].
func (o *OIDCAuthRequest) GetResponseType() oidc.ResponseType {
	return o.ResponseType
}

// GetScopes implements [op.AuthRequest].
func (o *OIDCAuthRequest) GetScopes() []string {
	return o.Scopes
}

// GetState implements [op.AuthRequest].
func (o *OIDCAuthRequest) GetState() string {
	return o.State
}

// GetSubject implements [op.AuthRequest].
func (o *OIDCAuthRequest) GetSubject() string {
	if o.Subject != "" {
		return o.Subject
	}

	err := o.fillAuthInfo()
	if err != nil {
		return ""
	}

	return o.Subject
}

func DeleteOIDCAuthRequest(requestID string) (err error) {
	authRequest, err := GetOIDCAuthRequestByRequestID(requestID)
	if err != nil {
		return err
	}
	err = db.RDBDel(rdb, RDB_PREFIX_OIDC_OP_REQUEST+requestID)
	if err != nil {
		return err
	}

	if authRequest == nil {
		return nil
	}

	_ = db.RDBDel(rdb, RDB_PREFIX_OIDC_OP_REQUEST_CHALLENGE+authRequest.CodeChallenge.Challenge)

	return nil
}

func GetOIDCAuthRequestByRequestID(requestID string) (authRequest *OIDCAuthRequest, err error) {
	authRequestStr, err := db.RDBGet(rdb, RDB_PREFIX_OIDC_OP_REQUEST+requestID)
	if err != nil {
		return nil, err
	}

	authRequest = &OIDCAuthRequest{}
	err = json.Unmarshal([]byte(authRequestStr), authRequest)
	if err != nil {
		return nil, err
	}

	return authRequest, nil
}
