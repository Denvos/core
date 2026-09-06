package auth

import (
    "net/http"
)

type Middleware struct {
    auth Authenticator
    onSuccess func(w http.ResponseWriter, r *http.Request)
    onFailure func(w http.ResponseWriter, r *http.Request)
}

func NewMiddleware(auth Authenticator) *Middleware {
    return &Middleware{
        auth: auth,
        onSuccess: func(w http.ResponseWriter, r *http.Request) {},
        onFailure: func(w http.ResponseWriter, r *http.Request) {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
        },
    }
}

func (m *Middleware) OnSuccess(fn func(w http.ResponseWriter, r *http.Request)) *Middleware {
    m.onSuccess = fn
    return m
}

func (m *Middleware) OnFailure(fn func(w http.ResponseWriter, r *http.Request)) *Middleware {
    m.onFailure = fn
    return m
}

func (m *Middleware) Handler(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ok, err := m.auth.Authenticate(r)
        if err != nil {
            http.Error(w, "Authentication error", http.StatusInternalServerError)
            return
        }
        if !ok {
            m.onFailure(w, r)
            return
        }
        m.onSuccess(w, r)
        next.ServeHTTP(w, r)
    })
}
