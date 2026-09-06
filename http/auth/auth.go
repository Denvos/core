package auth

import (
    "encoding/base64"
    "net/http"
    "strings"
)

type Authenticator interface {
    Authenticate(r *http.Request) (bool, error)
}

type BasicAuth struct {
    Username string
    Password string
}

func (b *BasicAuth) Authenticate(r *http.Request) (bool, error) {
    auth := r.Header.Get("Authorization")
    if auth == "" {
        return false, nil
    }
    if !strings.HasPrefix(auth, "Basic ") {
        return false, nil
    }
    payload, err := base64.StdEncoding.DecodeString(auth[6:])
    if err != nil {
        return false, err
    }
    parts := strings.SplitN(string(payload), ":", 2)
    if len(parts) != 2 {
        return false, nil
    }
    return parts[0] == b.Username && parts[1] == b.Password, nil
}

type BearerAuth struct {
    Token string
}

func (b *BearerAuth) Authenticate(r *http.Request) (bool, error) {
    auth := r.Header.Get("Authorization")
    if auth == "" {
        return false, nil
    }
    if !strings.HasPrefix(auth, "Bearer ") {
        return false, nil
    }
    token := auth[7:]
    return token == b.Token, nil
}

type APIKeyAuth struct {
    Header string
    Key    string
    Value  string
}

func (a *APIKeyAuth) Authenticate(r *http.Request) (bool, error) {
    if a.Header == "" {
        a.Header = "X-API-Key"
    }
    val := r.Header.Get(a.Header)
    if val == "" {
        return false, nil
    }
    return val == a.Value, nil
}

type MultiAuth struct {
    Authenticators []Authenticator
}

func (m *MultiAuth) Authenticate(r *http.Request) (bool, error) {
    for _, auth := range m.Authenticators {
        ok, err := auth.Authenticate(r)
        if err != nil {
            return false, err
        }
        if ok {
            return true, nil
        }
    }
    return false, nil
}
