package claims

import (
	"time"
)

type Claims struct {
	Subject   string                 `json:"sub"`
	Issuer    string                 `json:"iss"`
	Audience  string                 `json:"aud"`
	ExpiresAt int64                  `json:"exp"`
	IssuedAt  int64                  `json:"iat"`
	NotBefore int64                  `json:"nbf"`
	Extra     map[string]interface{} `json:"extra"`
}

func New() *Claims {
	return &Claims{
		Extra: make(map[string]interface{}),
	}
}

func (c *Claims) Valid() error {
	now := time.Now().Unix()
	if c.ExpiresAt > 0 && now > c.ExpiresAt {
		return &ValidationError{Field: "exp", Message: "token expired"}
	}
	if c.NotBefore > 0 && now < c.NotBefore {
		return &ValidationError{Field: "nbf", Message: "token not yet valid"}
	}
	return nil
}

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Message
}
