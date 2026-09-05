package hex

import "github.com/Denvos/core/encoding"

import (
	"encoding/hex"
)

type Codec struct{}

func (c Codec) Marshal(v interface{}) ([]byte, error) {
	data, ok := v.([]byte)
	if !ok {
		return nil, nil
	}
	dst := make([]byte, hex.EncodedLen(len(data)))
	hex.Encode(dst, data)
	return dst, nil
}

func (c Codec) Unmarshal(data []byte, v interface{}) error {
	dst, ok := v.(*[]byte)
	if !ok {
		return nil
	}
	*dst = make([]byte, hex.DecodedLen(len(data)))
	n, err := hex.Decode(*dst, data)
	if err != nil {
		return err
	}
	*dst = (*dst)[:n]
	return nil
}

func (c Codec) ContentType() string {
	return "application/hex"
}

func init() {
	encoding.Register("hex", Codec{})
}
