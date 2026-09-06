package logging

import (
    "bytes"
    "io"
    "net/http"
    "time"
)

func Middleware(logger Logger, opts ...Option) func(http.Handler) http.Handler {
    cfg := defaultConfig()
    for _, opt := range opts {
        opt(&cfg)
    }

    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            start := time.Now()
            rw := &responseWriter{
                ResponseWriter: w,
                status:         200,
                size:           0,
            }

            var body []byte
            if cfg.LogRequestBody {
                body, _ = io.ReadAll(r.Body)
                r.Body = io.NopCloser(bytes.NewReader(body))
            }

            next.ServeHTTP(rw, r)

            duration := time.Since(start)
            fields := map[string]interface{}{
                "method":     r.Method,
                "path":       r.URL.Path,
                "query":      r.URL.RawQuery,
                "status":     rw.status,
                "size":       rw.size,
                "duration":   duration.String(),
                "duration_ms": duration.Milliseconds(),
                "remote_addr": r.RemoteAddr,
                "user_agent":  r.UserAgent(),
                "referer":    r.Referer(),
                "request_id": r.Header.Get("X-Request-ID"),
            }

            if cfg.LogRequestBody && len(body) > 0 {
                fields["request_body"] = string(body)
            }

            if cfg.LogResponseBody && rw.body != nil {
                fields["response_body"] = string(rw.body)
            }

            if rw.status >= 500 {
                logger.Error("request error", fields)
            } else if rw.status >= 400 {
                logger.Warn("request warning", fields)
            } else {
                logger.Info("request", fields)
            }
        })
    }
}

type responseWriter struct {
    http.ResponseWriter
    status int
    size   int
    body   []byte
}

func (rw *responseWriter) Write(b []byte) (int, error) {
    if rw.body == nil {
        rw.body = make([]byte, 0, len(b))
    }
    rw.body = append(rw.body, b...)
    n, err := rw.ResponseWriter.Write(b)
    rw.size += n
    return n, err
}

func (rw *responseWriter) WriteHeader(statusCode int) {
    rw.status = statusCode
    rw.ResponseWriter.WriteHeader(statusCode)
}
