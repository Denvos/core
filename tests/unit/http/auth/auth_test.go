package auth

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestBasicAuth(t *testing.T) {
    auth := &BasicAuth{Username: "admin", Password: "pass"}
    r := httptest.NewRequest("GET", "/", nil)
    r.SetBasicAuth("admin", "pass")
    ok, err := auth.Authenticate(r)
    if err != nil {
        t.Fatal(err)
    }
    if !ok {
        t.Error("expected auth to pass")
    }
}

func TestBearerAuth(t *testing.T) {
    auth := &BearerAuth{Token: "abc123"}
    r := httptest.NewRequest("GET", "/", nil)
    r.Header.Set("Authorization", "Bearer abc123")
    ok, err := auth.Authenticate(r)
    if err != nil {
        t.Fatal(err)
    }
    if !ok {
        t.Error("expected auth to pass")
    }
}

func TestAPIKeyAuth(t *testing.T) {
    auth := &APIKeyAuth{Value: "key123"}
    r := httptest.NewRequest("GET", "/", nil)
    r.Header.Set("X-API-Key", "key123")
    ok, err := auth.Authenticate(r)
    if err != nil {
        t.Fatal(err)
    }
    if !ok {
        t.Error("expected auth to pass")
    }
}

func TestMiddleware(t *testing.T) {
    auth := &BasicAuth{Username: "admin", Password: "pass"}
    mw := NewMiddleware(auth)
    handler := mw.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    }))

    r := httptest.NewRequest("GET", "/", nil)
    r.SetBasicAuth("admin", "pass")
    w := httptest.NewRecorder()
    handler.ServeHTTP(w, r)
    if w.Code != http.StatusOK {
        t.Error("expected 200")
    }
}
