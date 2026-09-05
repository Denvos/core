package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

type Provider struct {
	Name         string
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Endpoint     oauth2.Endpoint
	Scopes       []string
	UserInfoURL  string
}

type Config struct {
	Providers []Provider
}

type OAuth struct {
	configs map[string]*oauth2.Config
	providers map[string]Provider
}

func New(cfg *Config) *OAuth {
	o := &OAuth{
		configs:   make(map[string]*oauth2.Config),
		providers: make(map[string]Provider),
	}
	for _, p := range cfg.Providers {
		o.configs[p.Name] = &oauth2.Config{
			ClientID:     p.ClientID,
			ClientSecret: p.ClientSecret,
			RedirectURL:  p.RedirectURL,
			Endpoint:     p.Endpoint,
			Scopes:       p.Scopes,
		}
		o.providers[p.Name] = p
	}
	return o
}

func (o *OAuth) AuthCodeURL(provider string, state string) (string, error) {
	cfg, ok := o.configs[provider]
	if !ok {
		return "", fmt.Errorf("unknown provider: %s", provider)
	}
	return cfg.AuthCodeURL(state, oauth2.AccessTypeOffline), nil
}

func (o *OAuth) Exchange(ctx context.Context, provider string, code string) (*oauth2.Token, error) {
	cfg, ok := o.configs[provider]
	if !ok {
		return nil, fmt.Errorf("unknown provider: %s", provider)
	}
	return cfg.Exchange(ctx, code)
}

func (o *OAuth) UserInfo(ctx context.Context, provider string, token *oauth2.Token) (map[string]interface{}, error) {
	p, ok := o.providers[provider]
	if !ok {
		return nil, fmt.Errorf("unknown provider: %s", provider)
	}
	client := o.configs[provider].Client(ctx, token)
	resp, err := client.Get(p.UserInfoURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return data, nil
}
