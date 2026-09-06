package request

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "io"
    "mime/multipart"
    "net/http"
    "net/url"
    "strings"
    "time"
)

type Request struct {
    method  string
    url     string
    headers map[string]string
    body    []byte
    ctx     context.Context
    timeout time.Duration
    client  *http.Client
}

type Option func(*Request)

func New(method, url string, opts ...Option) *Request {
    r := &Request{
        method:  method,
        url:     url,
        headers: make(map[string]string),
        timeout: 30 * time.Second,
        ctx:     context.Background(),
    }
    for _, opt := range opts {
        opt(r)
    }
    return r
}

func WithHeader(key, value string) Option {
    return func(r *Request) {
        r.headers[key] = value
    }
}

func WithHeaders(headers map[string]string) Option {
    return func(r *Request) {
        for k, v := range headers {
            r.headers[k] = v
        }
    }
}

func WithBody(data []byte) Option {
    return func(r *Request) {
        r.body = data
    }
}

func WithJSONBody(v interface{}) Option {
    return func(r *Request) {
        data, err := json.Marshal(v)
        if err != nil {
            return
        }
        r.body = data
        r.headers["Content-Type"] = "application/json"
    }
}

func WithContext(ctx context.Context) Option {
    return func(r *Request) {
        r.ctx = ctx
    }
}

func WithTimeout(d time.Duration) Option {
    return func(r *Request) {
        r.timeout = d
    }
}

func WithClient(client *http.Client) Option {
    return func(r *Request) {
        r.client = client
    }
}

func (r *Request) Do() (*Response, error) {
    client := r.client
    if client == nil {
        client = &http.Client{Timeout: r.timeout}
    }
    var bodyReader io.Reader
    if r.body != nil {
        bodyReader = bytes.NewReader(r.body)
    }
    req, err := http.NewRequestWithContext(r.ctx, r.method, r.url, bodyReader)
    if err != nil {
        return nil, err
    }
    for k, v := range r.headers {
        req.Header.Set(k, v)
    }
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    return &Response{Response: resp}, nil
}
