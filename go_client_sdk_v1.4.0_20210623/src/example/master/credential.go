package master

import (
	"context"
)

const (
	JWT_TOKEN_KEY = "token"
)


type JWTCredential struct {
	Token string
}

func (c* JWTCredential) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string {
		JWT_TOKEN_KEY: c.Token,
	}, nil
}

func (c* JWTCredential) RequireTransportSecurity() bool {
	return true
}
