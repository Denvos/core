package compression

import (
    "io"
    "net/http"
)

type Compressor interface {
    Name() string
    Encoding() string
    NewWriter(w io.Writer) (io.WriteCloser, error)
    NewReader(r io.Reader) (io.ReadCloser, error)
    Match(acceptEncoding string) bool
}

type Config struct {
    // Enabled compressors (by name)
    Enabled []string
    // Minimum response size to compress (bytes)
    MinSize int
    // Content types to compress (if empty, compress all)
    ContentTypes []string
}

var DefaultConfig = Config{
    Enabled:     []string{"gzip", "zstd", "br"},
    MinSize:     1024,
    ContentTypes: []string{
        "text/plain",
        "text/html",
        "text/css",
        "text/javascript",
        "application/json",
        "application/xml",
        "application/yaml",
        "application/javascript",
        "application/x-javascript",
    },
}
