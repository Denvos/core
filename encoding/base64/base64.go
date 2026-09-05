package base64
import "github.com/Denvos/core/encoding"

import (
    "encoding/base64"
)

type Codec struct{}

func (c Codec) Marshal(v interface{}) ([]byte, error) {
    data, ok := v.([]byte)
    if !ok {
        return nil, nil
    }
    dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
    base64.StdEncoding.Encode(dst, data)
    return dst, nil
}

func (c Codec) Unmarshal(data []byte, v interface{}) error {
    dst, ok := v.(*[]byte)
    if !ok {
        return nil
    }
    *dst = make([]byte, base64.StdEncoding.DecodedLen(len(data)))
    n, err := base64.StdEncoding.Decode(*dst, data)
    if err != nil {
        return err
    }
    *dst = (*dst)[:n]
    return nil
}

func (c Codec) ContentType() string {
    return "application/base64"
}

func init() {
    encoding.Register("base64", Codec{})
}
