package oidc

import (
	"context"
	"fmt"

	"github.com/coreos/go-oidc/v3/oidc"
)

type OIDC struct {
	provider  *oidc.Provider
	verifier  *oidc.IDTokenVerifier
	clientID  string
}

type Config struct {
	IssuerURL      string
	ClientID       string
	ClientSecret   string
	RedirectURL    string
	Scopes         []string
}

func New(cfg *Config) (*OIDC, error) {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, cfg.IssuerURL)
	if err != nil {
		return nil, err
	}
	verifier := provider.Verifier(&oidc.Config{ClientID: cfg.ClientID})
	return &OIDC{
		provider: provider,
		verifier: verifier,
		clientID: cfg.ClientID,
	}, nil
}

func (o *OIDC) AuthURL(state string) string {
	// return a constructed auth URL
	return ""
}

func (o *OIDC) Exchange(ctx context.Context, code string) (*oidc.IDToken, error) {
	return nil, nil
}

func (o *OIDC) Verify(ctx context.Context, rawIDToken string) (*oidc.IDToken, error) {
	return o.verifier.Verify(ctx, rawIDToken)
}
