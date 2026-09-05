package url

import (
	"fmt"
	"net/url"
)

type Validator struct {
	RequireScheme bool
}

func New() *Validator {
	return &Validator{RequireScheme: true}
}

func (v *Validator) Validate(raw string) error {
	if raw == "" {
		return fmt.Errorf("URL cannot be empty")
	}
	u, err := url.Parse(raw)
	if err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}
	if v.RequireScheme && u.Scheme == "" {
		return fmt.Errorf("URL requires scheme (http, https, etc.)")
	}
	if u.Host == "" {
		return fmt.Errorf("URL requires host")
	}
	return nil
}
