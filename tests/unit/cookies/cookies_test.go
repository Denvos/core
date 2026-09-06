package cookies

import (
    "net/http"
    "testing"
    "time"
)

func TestCookieBuilder(t *testing.T) {
    c := New("test", "value").
        SetPath("/api").
        SetDomain("example.com").
        SetMaxAge(3600).
        SetSecure(true).
        SetHttpOnly(true)

    if c.Name != "test" {
        t.Error("expected name test")
    }
    if c.Path != "/api" {
        t.Error("expected path /api")
    }
    if c.MaxAge != 3600 {
        t.Error("expected max-age 3600")
    }
    if !c.Secure {
        t.Error("expected secure true")
    }
    if !c.HttpOnly {
        t.Error("expected httpOnly true")
    }
}

func TestSetGet(t *testing.T) {
    w := &mockResponseWriter{header: make(http.Header)}
    r := &http.Request{Header: make(http.Header)}

    SetString(w, "test", "value")
    r.Header.Set("Cookie", "test=value")

    val := Get(r, "test")
    if val != "value" {
        t.Errorf("expected value, got %s", val)
    }

    Delete(w, "test")
    // Check that the cookie was set with MaxAge=-1
    setCookie := w.header.Get("Set-Cookie")
    if setCookie == "" {
        t.Error("expected Set-Cookie header")
    }
}

type mockResponseWriter struct {
    header http.Header
}

func (m *mockResponseWriter) Header() http.Header {
    return m.header
}

func (m *mockResponseWriter) Write(b []byte) (int, error) {
    return len(b), nil
}

func (m *mockResponseWriter) WriteHeader(statusCode int) {}
