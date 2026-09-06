package logging

import (
    "bytes"
    "io"
    "net/http"
    "net/http/httptest"
    "testing"
)

type testLogger struct {
    messages []string
}

func (t *testLogger) Debug(msg string, fields ...interface{}) {
    t.messages = append(t.messages, "DEBUG: "+msg)
}

func (t *testLogger) Info(msg string, fields ...interface{}) {
    t.messages = append(t.messages, "INFO: "+msg)
}

func (t *testLogger) Warn(msg string, fields ...interface{}) {
    t.messages = append(t.messages, "WARN: "+msg)
}

func (t *testLogger) Error(msg string, fields ...interface{}) {
    t.messages = append(t.messages, "ERROR: "+msg)
}

func TestMiddleware(t *testing.T) {
    logger := &testLogger{}
    handler := Middleware(logger)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("ok"))
    }))

    req := httptest.NewRequest("GET", "/test", nil)
    w := httptest.NewRecorder()
    handler.ServeHTTP(w, req)

    if len(logger.messages) == 0 {
        t.Error("expected log message")
    }
}

func TestRequestBodyLogging(t *testing.T) {
    logger := &testLogger{}
    handler := Middleware(logger, WithRequestBody())(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        body, _ := io.ReadAll(r.Body)
        if string(body) != "hello" {
            t.Errorf("expected hello, got %s", string(body))
        }
        w.WriteHeader(http.StatusOK)
    }))

    req := httptest.NewRequest("POST", "/", bytes.NewReader([]byte("hello")))
    w := httptest.NewRecorder()
    handler.ServeHTTP(w, req)
}
