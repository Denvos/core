package ascii85

import (
    "encoding/ascii85"
    "github.com/Denvos/core/encoding"
)

type Codec struct{}

func (c Codec) Marshal(v interface{}) ([]byte, error) {
    data, ok := v.([]byte)
    if !ok {
        return nil, nil
    }
    dst := make([]byte, ascii85.MaxEncodedLen(len(data)))
    n := ascii85.Encode(dst, data)
    return dst[:n], nil
}

func (c Codec) Unmarshal(data []byte, v interface{}) error {
    dst, ok := v.(*[]byte)
    if !ok {
        return nil
    }
    *dst = make([]byte, len(data))
    n, _, err := ascii85.Decode(*dst, data, true)
    if err != nil {
        return err
    }
    *dst = (*dst)[:n]
    return nil
}

func (c Codec) ContentType() string {
    return "application/ascii85"
}

func init() {
    encoding.Register("ascii85", Codec{})
}
