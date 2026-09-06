package request

import (
    "encoding/json"
    "fmt"
    "strings"
)

type Builder struct {
    method  string
    url     string
    params  map[string]string
    headers map[string]string
    body    []byte
}

func NewBuilder(method, url string) *Builder {
    return &Builder{
        method:  method,
        url:     url,
        params:  make(map[string]string),
        headers: make(map[string]string),
    }
}

func (b *Builder) Param(key, value string) *Builder {
    b.params[key] = value
    return b
}

func (b *Builder) Params(params map[string]string) *Builder {
    for k, v := range params {
        b.params[k] = v
    }
    return b
}

func (b *Builder) Header(key, value string) *Builder {
    b.headers[key] = value
    return b
}

func (b *Builder) Headers(headers map[string]string) *Builder {
    for k, v := range headers {
        b.headers[k] = v
    }
    return b
}

func (b *Builder) Body(data []byte) *Builder {
    b.body = data
    return b
}

func (b *Builder) JSON(v interface{}) *Builder {
    data, err := json.Marshal(v)
    if err == nil {
        b.body = data
        b.headers["Content-Type"] = "application/json"
    }
    return b
}

func (b *Builder) Build() (*Request, error) {
    urlStr := b.url
    if len(b.params) > 0 {
        q := url.Values{}
        for k, v := range b.params {
            q.Set(k, v)
        }
        if strings.Contains(urlStr, "?") {
            urlStr += "&" + q.Encode()
        } else {
            urlStr += "?" + q.Encode()
        }
    }
    req := New(b.method, urlStr)
    for k, v := range b.headers {
        req.headers[k] = v
    }
    if b.body != nil {
        req.body = b.body
    }
    return req, nil
}
