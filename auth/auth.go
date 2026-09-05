package auth

import (
	"context"
	"time"
)

type Authenticator interface {
	Authenticate(ctx context.Context, credentials interface{}) (*Identity, error)
}

type Authorizer interface {
	Authorize(ctx context.Context, identity *Identity, resource string, action string) (bool, error)
}

type TokenManager interface {
	Generate(ctx context.Context, identity *Identity, ttl time.Duration) (string, error)
	Validate(ctx context.Context, token string) (*Identity, error)
	Revoke(ctx context.Context, token string) error
}

type Identity struct {
	ID          string                 `json:"id"`
	Subject     string                 `json:"subject"`
	Username    string                 `json:"username"`
	Email       string                 `json:"email"`
	Roles       []string               `json:"roles"`
	Permissions []string               `json:"permissions"`
	Claims      map[string]interface{} `json:"claims"`
	ExpiresAt   time.Time              `json:"expires_at"`
	IssuedAt    time.Time              `json:"issued_at"`
}

type Credentials struct {
	Type     string                 `json:"type"`
	Username string                 `json:"username,omitempty"`
	Password string                 `json:"password,omitempty"`
	Token    string                 `json:"token,omitempty"`
	APIKey   string                 `json:"api_key,omitempty"`
	Provider string                 `json:"provider,omitempty"`
	Code     string                 `json:"code,omitempty"`
	Extra    map[string]interface{} `json:"extra,omitempty"`
}
