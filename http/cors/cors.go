package cors

import (
    "net/http"
    "strings"
)

type Config struct {
    AllowOrigins     []string
    AllowMethods     []string
    AllowHeaders     []string
    ExposeHeaders    []string
    AllowCredentials bool
    MaxAge           int
}

var DefaultConfig = Config{
    AllowOrigins:  []string{"*"},
    AllowMethods:  []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch, http.MethodOptions},
    AllowHeaders:  []string{"*"},
    ExposeHeaders: []string{},
    MaxAge:        86400,
}

func Middleware(cfg Config) func(http.Handler) http.Handler {
    if len(cfg.AllowOrigins) == 0 {
        cfg.AllowOrigins = DefaultConfig.AllowOrigins
    }
    if len(cfg.AllowMethods) == 0 {
        cfg.AllowMethods = DefaultConfig.AllowMethods
    }
    if len(cfg.AllowHeaders) == 0 {
        cfg.AllowHeaders = DefaultConfig.AllowHeaders
    }
    if len(cfg.ExposeHeaders) == 0 {
        cfg.ExposeHeaders = DefaultConfig.ExposeHeaders
    }
    if cfg.MaxAge == 0 {
        cfg.MaxAge = DefaultConfig.MaxAge
    }

    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            origin := r.Header.Get("Origin")
            if origin == "" {
                next.ServeHTTP(w, r)
                return
            }

            if !isOriginAllowed(origin, cfg.AllowOrigins) {
                next.ServeHTTP(w, r)
                return
            }

            w.Header().Set("Access-Control-Allow-Origin", origin)

            if cfg.AllowCredentials {
                w.Header().Set("Access-Control-Allow-Credentials", "true")
            }

            if len(cfg.ExposeHeaders) > 0 {
                w.Header().Set("Access-Control-Expose-Headers", strings.Join(cfg.ExposeHeaders, ", "))
            }

            if r.Method == http.MethodOptions {
                w.Header().Set("Access-Control-Allow-Methods", strings.Join(cfg.AllowMethods, ", "))
                w.Header().Set("Access-Control-Allow-Headers", strings.Join(cfg.AllowHeaders, ", "))
                if cfg.MaxAge > 0 {
                    w.Header().Set("Access-Control-Max-Age", string(rune(cfg.MaxAge)))
                }
                w.WriteHeader(http.StatusNoContent)
                return
            }

            next.ServeHTTP(w, r)
        })
    }
}

func isOriginAllowed(origin string, allowed []string) bool {
    for _, a := range allowed {
        if a == "*" {
            return true
        }
        if a == origin {
            return true
        }
        if strings.HasPrefix(a, "*.") {
            suffix := a[1:]
            if strings.HasSuffix(origin, suffix) {
                return true
            }
        }
    }
    return false
}
