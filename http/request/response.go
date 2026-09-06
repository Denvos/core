package request

import (
    "encoding/json"
    "io"
    "net/http"
)

type Response struct {
    *http.Response
}

func (r *Response) BodyBytes() ([]byte, error) {
    if r.Response == nil {
        return nil, nil
    }
    defer r.Response.Body.Close()
    return io.ReadAll(r.Response.Body)
}

func (r *Response) BodyString() (string, error) {
    data, err := r.BodyBytes()
    if err != nil {
        return "", err
    }
    return string(data), nil
}

func (r *Response) JSON(v interface{}) error {
    data, err := r.BodyBytes()
    if err != nil {
        return err
    }
    return json.Unmarshal(data, v)
}

func (r *Response) IsSuccess() bool {
    return r.Response != nil && r.Response.StatusCode >= 200 && r.Response.StatusCode < 300
}

func (r *Response) IsError() bool {
    return !r.IsSuccess()
}

func (r *Response) Status() string {
    if r.Response == nil {
        return "no response"
    }
    return r.Response.Status
}
