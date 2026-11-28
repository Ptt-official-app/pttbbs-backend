package oidcop

import (
	"context"

	"github.com/zitadel/oidc/v3/pkg/op"
)

var (
	_ op.Storage                  = &Storage{}
	_ op.ClientCredentialsStorage = &Storage{}

	PROVIDER          *op.Provider
	AUTH_CALLBACK_URL func(ctx context.Context, oidcID string, callbackSecret string) string
)

const (
	// op appends "/callback" to the authorization endpoint ("/api/oidc/auth").
	AUTH_CALLBACK_PATH         = "/api/oidc/auth/callback"
	AUTH_CALLBACK_SECRET_PARAM = "s"
)
