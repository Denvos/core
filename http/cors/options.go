package cors

import "net/http"

type Option func(*Config)

func WithAllowOrigins(origins ...string) Option {
    return func(c *Config) {
        c.AllowOrigins = origins
    }
}

func WithAllowMethods(methods ...string) Option {
    return func(c *Config) {
        c.AllowMethods = methods
    }
}

func WithAllowHeaders(headers ...string) Option {
    return func(c *Config) {
        c.AllowHeaders = headers
    }
}

func WithExposeHeaders(headers ...string) Option {
    return func(c *Config) {
        c.ExposeHeaders = headers
    }
}

func WithAllowCredentials(allow bool) Option {
    return func(c *Config) {
        c.AllowCredentials = allow
    }
}

func WithMaxAge(seconds int) Option {
    return func(c *Config) {
        c.MaxAge = seconds
    }
}

func NewConfig(opts ...Option) Config {
    cfg := DefaultConfig
    for _, opt := range opts {
        opt(&cfg)
    }
    return cfg
}

func AllowAll() Config {
    return Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch, http.MethodOptions},
        AllowHeaders:     []string{"*"},
        ExposeHeaders:    []string{},
        AllowCredentials: false,
        MaxAge:           86400,
    }
}

func DisallowAll() Config {
    return Config{
        AllowOrigins:     []string{},
        AllowMethods:     []string{},
        AllowHeaders:     []string{},
        ExposeHeaders:    []string{},
        AllowCredentials: false,
        MaxAge:           0,
    }
}
