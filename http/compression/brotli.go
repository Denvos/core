package compression

import (
    "io"
    "strings"

    "github.com/andybalholm/brotli"
)

type BrotliCompressor struct {
    quality int
}

func NewBrotli(opts ...func(*BrotliCompressor)) *BrotliCompressor {
    b := &BrotliCompressor{quality: brotli.DefaultQuality}
    for _, opt := range opts {
        opt(b)
    }
    return b
}

func WithBrotliQuality(quality int) func(*BrotliCompressor) {
    return func(b *BrotliCompressor) {
        b.quality = quality
    }
}

func (b *BrotliCompressor) Name() string {
    return "br"
}

func (b *BrotliCompressor) Encoding() string {
    return "br"
}

func (b *BrotliCompressor) NewWriter(w io.Writer) (io.WriteCloser, error) {
    return brotli.NewWriterLevel(w, b.quality), nil
}

func (b *BrotliCompressor) NewReader(r io.Reader) (io.ReadCloser, error) {
    return brotli.NewReader(r), nil
}

func (b *BrotliCompressor) Match(acceptEncoding string) bool {
    return strings.Contains(acceptEncoding, "br")
}
