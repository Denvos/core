package compression

import (
    "io"
    "strings"

    "github.com/klauspost/compress/zstd"
)

type ZstdCompressor struct {
    level int
}

func NewZstd(opts ...func(*ZstdCompressor)) *ZstdCompressor {
    z := &ZstdCompressor{level: zstd.SpeedDefault}
    for _, opt := range opts {
        opt(z)
    }
    return z
}

func WithZstdLevel(level int) func(*ZstdCompressor) {
    return func(z *ZstdCompressor) {
        z.level = level
    }
}

func (z *ZstdCompressor) Name() string {
    return "zstd"
}

func (z *ZstdCompressor) Encoding() string {
    return "zstd"
}

func (z *ZstdCompressor) NewWriter(w io.Writer) (io.WriteCloser, error) {
    enc, err := zstd.NewWriter(w, zstd.WithEncoderLevel(zstd.EncoderLevel(z.level)))
    if err != nil {
        return nil, err
    }
    return enc, nil
}

func (z *ZstdCompressor) NewReader(r io.Reader) (io.ReadCloser, error) {
    return zstd.NewReader(r), nil
}

func (z *ZstdCompressor) Match(acceptEncoding string) bool {
    return strings.Contains(acceptEncoding, "zstd") || strings.Contains(acceptEncoding, "zstd")
}
