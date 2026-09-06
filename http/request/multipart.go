package request

import (
    "bytes"
    "fmt"
    "io"
    "mime/multipart"
    "os"
    "path/filepath"
)

type MultipartBuilder struct {
    body   *bytes.Buffer
    writer *multipart.Writer
}

func NewMultipartBuilder() *MultipartBuilder {
    b := &MultipartBuilder{
        body: &bytes.Buffer{},
    }
    b.writer = multipart.NewWriter(b.body)
    return b
}

func (m *MultipartBuilder) Field(name, value string) *MultipartBuilder {
    m.writer.WriteField(name, value)
    return m
}

func (m *MultipartBuilder) File(name, path string) (*MultipartBuilder, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    part, err := m.writer.CreateFormFile(name, filepath.Base(path))
    if err != nil {
        return nil, err
    }
    _, err = io.Copy(part, file)
    return m, err
}

func (m *MultipartBuilder) FileReader(name, filename string, reader io.Reader) (*MultipartBuilder, error) {
    part, err := m.writer.CreateFormFile(name, filename)
    if err != nil {
        return nil, err
    }
    _, err = io.Copy(part, reader)
    return m, err
}

func (m *MultipartBuilder) Close() error {
    return m.writer.Close()
}

func (m *MultipartBuilder) ContentType() string {
    return m.writer.FormDataContentType()
}

func (m *MultipartBuilder) Body() []byte {
    return m.body.Bytes()
}
