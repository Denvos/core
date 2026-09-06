package compression

import (
    "bytes"
    "io"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestGzipCompression(t *testing.T) {
    reg := NewRegistry()
    comp := reg.Get("gzip")
    if comp == nil {
        t.Fatal("gzip compressor not registered")
    }

    data := []byte("hello world")
    buf := new(bytes.Buffer)
    w, err := comp.NewWriter(buf)
    if err != nil {
        t.Fatal(err)
    }
    w.Write(data)
    w.Close()

    // Decompress
    r, err := comp.NewReader(buf)
    if err != nil {
        t.Fatal(err)
    }
    out, err := io.ReadAll(r)
    if err != nil {
        t.Fatal(err)
    }
    if string(out) != "hello world" {
        t.Errorf("expected hello world, got %s", string(out))
    }
}

func TestMiddleware(t *testing.T) {
    reg := NewRegistry()
    cfg := DefaultConfig
    cfg.MinSize = 10
    mw := NewMiddleware(reg, cfg)

    handler := mw.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/plain")
        w.Write([]byte("hello world"))
    }))

    req := httptest.NewRequest("GET", "/", nil)
    req.Header.Set("Accept-Encoding", "gzip")
    w := httptest.NewRecorder()
    handler.ServeHTTP(w, req)

    // Check Content-Encoding header
    if w.Header().Get("Content-Encoding") != "gzip" {
        t.Error("expected gzip encoding")
    }
}
