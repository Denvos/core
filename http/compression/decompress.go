package compression

import (
    "io"
    "net/http"
    "strings"
)

func DecompressRequest(r *http.Request) (io.ReadCloser, error) {
    contentEncoding := r.Header.Get("Content-Encoding")
    if contentEncoding == "" {
        return r.Body, nil
    }

    // Handle multiple encodings (e.g., "gzip, deflate")
    encodings := strings.Split(contentEncoding, ",")
    var reader io.ReadCloser = r.Body
    var err error

    // Process in reverse order (last applied first)
    for i := len(encodings) - 1; i >= 0; i-- {
        enc := strings.TrimSpace(encodings[i])
        switch enc {
        case "gzip", "deflate":
            // gzip.Reader handles both gzip and deflate
            gz, err := gzip.NewReader(reader)
            if err != nil {
                return nil, err
            }
            reader = gz
        case "br":
            // Brotli
            br := brotli.NewReader(reader)
            reader = &brotliReadCloser{Reader: br, closer: reader}
        case "zstd":
            // Zstd
            zr, err := zstd.NewReader(reader)
            if err != nil {
                return nil, err
            }
            reader = &zstdReadCloser{Decoder: zr, closer: reader}
        default:
            // Unknown encoding, ignore
            continue
        }
    }
    return reader, nil
}

type brotliReadCloser struct {
    io.Reader
    closer io.Closer
}

func (b *brotliReadCloser) Close() error {
    if b.closer != nil {
        return b.closer.Close()
    }
    return nil
}

type zstdReadCloser struct {
    *zstd.Decoder
    closer io.Closer
}

func (z *zstdReadCloser) Close() error {
    z.Decoder.Close()
    if z.closer != nil {
        return z.closer.Close()
    }
    return nil
}
