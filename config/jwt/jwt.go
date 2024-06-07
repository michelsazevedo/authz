package jwt

import (
	"context"
	"github.com/michelsazevedo/authz/domain"
)

const authentication = AuthenticationContext("settings")

type AuthenticationContext string

func FromContext(ctx context.Context) domain.JwtToken {
	return ctx.Value(authentication).(domain.JwtToken)
}

func ToContext(ctx context.Context, jwt domain.JwtToken) context.Context {
	return context.WithValue(ctx, authentication, jwt)
}
