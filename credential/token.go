package credential

import "context"

type TokenCreds struct {
	Token string
}

func (c *TokenCreds) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		":authorization": "Bearer " + c.Token,
	}, nil
}

func (c *TokenCreds) RequireTransportSecurity() bool {
	return false
}
