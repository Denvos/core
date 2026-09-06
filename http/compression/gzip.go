package compression

import (
    "compress/gzip"
    "io"
    "strings"
)

type GzipCompressor struct {
    level int
}

func NewGzip(opts ...func(*GzipCompressor)) *GzipCompressor {
    g := &GzipCompressor{level: gzip.DefaultCompression}
    for _, opt := range opts {
        opt(g)
    }
    return g
}

func WithGzipLevel(level int) func(*GzipCompressor) {
    return func(g *GzipCompressor) {
        g.level = level
    }
}

func (g *GzipCompressor) Name() string {
    return "gzip"
}

func (g *GzipCompressor) Encoding() string {
    return "gzip"
}

func (g *GzipCompressor) NewWriter(w io.Writer) (io.WriteCloser, error) {
    return gzip.NewWriterLevel(w, g.level)
}

func (g *GzipCompressor) NewReader(r io.Reader) (io.ReadCloser, error) {
    return gzip.NewReader(r)
}

func (g *GzipCompressor) Match(acceptEncoding string) bool {
    return strings.Contains(acceptEncoding, "gzip")
}
