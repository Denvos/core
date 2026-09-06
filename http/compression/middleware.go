package compression

import (
    "bytes"
    "compress/gzip"
    "io"
    "net/http"
    "strings"
)

type Middleware struct {
    registry *Registry
    config   Config
}

func NewMiddleware(registry *Registry, config Config) *Middleware {
    if registry == nil {
        registry = NewRegistry()
    }
    if config.MinSize <= 0 {
        config.MinSize = DefaultConfig.MinSize
    }
    if len(config.Enabled) == 0 {
        config.Enabled = DefaultConfig.Enabled
    }
    return &Middleware{
        registry: registry,
        config:   config,
    }
}

func (m *Middleware) Handler(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Check if client accepts compression
        acceptEncoding := r.Header.Get("Accept-Encoding")
        if acceptEncoding == "" {
            next.ServeHTTP(w, r)
            return
        }

        // Negotiate compressor
        compressor := m.registry.Negotiate(acceptEncoding)
        if compressor == nil {
            next.ServeHTTP(w, r)
            return
        }

        // Wrap response writer
        cw := &compressWriter{
            ResponseWriter: w,
            compressor:     compressor,
            minSize:        m.config.MinSize,
            contentTypes:   m.config.ContentTypes,
        }
        next.ServeHTTP(cw, r)
        cw.Flush()
    })
}

type compressWriter struct {
    http.ResponseWriter
    compressor   Compressor
    minSize      int
    contentTypes []string
    buf          *bytes.Buffer
    wroteHeader  bool
    compressed   bool
}

func (cw *compressWriter) Write(b []byte) (int, error) {
    if !cw.wroteHeader {
        cw.WriteHeader(http.StatusOK)
    }
    if !cw.compressed {
        // Check content type and size
        contentType := cw.Header().Get("Content-Type")
        if cw.shouldCompress(contentType, len(b)) {
            cw.compressed = true
            // Create compressor writer
            w, err := cw.compressor.NewWriter(cw.ResponseWriter)
            if err != nil {
                // Fallback to uncompressed
                cw.compressed = false
                return cw.ResponseWriter.Write(b)
            }
            // Set Content-Encoding header
            cw.Header().Set("Content-Encoding", cw.compressor.Encoding())
            cw.Header().Del("Content-Length")
            // Write original data through compressor
            n, err := w.Write(b)
            if err == nil {
                w.Close()
            }
            return n, err
        }
        // No compression
        return cw.ResponseWriter.Write(b)
    }
    // Already compressed, write directly
    return cw.ResponseWriter.Write(b)
}

func (cw *compressWriter) WriteHeader(statusCode int) {
    if cw.wroteHeader {
        return
    }
    cw.wroteHeader = true
    cw.ResponseWriter.WriteHeader(statusCode)
}

func (cw *compressWriter) shouldCompress(contentType string, size int) bool {
    if size < cw.minSize {
        return false
    }
    if len(cw.contentTypes) == 0 {
        return true
    }
    for _, ct := range cw.contentTypes {
        if strings.HasPrefix(contentType, ct) {
            return true
        }
    }
    return false
}

func (cw *compressWriter) Flush() {
    if flusher, ok := cw.ResponseWriter.(http.Flusher); ok {
        flusher.Flush()
    }
}
