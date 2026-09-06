package cors

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestMiddleware(t *testing.T) {
    cfg := AllowAll()
    handler := Middleware(cfg)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    }))

    req := httptest.NewRequest(http.MethodGet, "/", nil)
    req.Header.Set("Origin", "https://example.com")
    w := httptest.NewRecorder()
    handler.ServeHTTP(w, req)

    if w.Header().Get("Access-Control-Allow-Origin") != "https://example.com" {
        t.Error("expected Access-Control-Allow-Origin: https://example.com")
    }
}

func TestPreflight(t *testing.T) {
    cfg := AllowAll()
    handler := Middleware(cfg)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    }))

    req := httptest.NewRequest(http.MethodOptions, "/", nil)
    req.Header.Set("Origin", "https://example.com")
    w := httptest.NewRecorder()
    handler.ServeHTTP(w, req)

    if w.Code != http.StatusNoContent {
        t.Error("expected 204 No Content")
    }
    if w.Header().Get("Access-Control-Allow-Methods") == "" {
        t.Error("expected Access-Control-Allow-Methods header")
    }
}

func TestOriginRestriction(t *testing.T) {
    cfg := NewConfig(WithAllowOrigins("https://allowed.com"))
    handler := Middleware(cfg)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    }))

    req := httptest.NewRequest(http.MethodGet, "/", nil)
    req.Header.Set("Origin", "https://notallowed.com")
    w := httptest.NewRecorder()
    handler.ServeHTTP(w, req)

    if w.Header().Get("Access-Control-Allow-Origin") != "" {
        t.Error("expected no Access-Control-Allow-Origin header")
    }
}
