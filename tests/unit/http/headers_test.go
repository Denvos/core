package headers

import (
    "net/http"
    "testing"
)

func TestGetBearerToken(t *testing.T) {
    r, _ := http.NewRequest("GET", "/", nil)
    r.Header.Set(HeaderAuthorization, "Bearer token123")
    if GetBearerToken(r) != "token123" {
        t.Error("expected token123")
    }
}

func TestParseAccept(t *testing.T) {
    header := "application/json;q=0.8, text/html;q=0.9"
    items := ParseAccept(header)
    if len(items) != 2 {
        t.Fatal("expected 2 items")
    }
    if items[0].Type != "text" || items[0].Subtype != "html" {
        t.Error("expected text/html first")
    }
    if items[1].Type != "application" || items[1].Subtype != "json" {
        t.Error("expected application/json second")
    }
}

func TestBuildUserAgent(t *testing.T) {
    ua := BuildUserAgent("MyApp", "1.0", "linux", "amd64")
    expected := "MyApp/1.0 (linux; amd64)"
    if ua != expected {
        t.Errorf("expected %s, got %s", expected, ua)
    }
}
