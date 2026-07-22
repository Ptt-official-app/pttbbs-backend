package oidcop

import (
	"context"

	"github.com/zitadel/oidc/v3/pkg/op"
)

var (
	_ op.Storage                  = &Storage{}
	_ op.ClientCredentialsStorage = &Storage{}

	PROVIDER          *op.Provider
	AUTH_CALLBACK_URL func(context.Context, string) string
)
