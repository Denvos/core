package gob

import "github.com/Denvos/core/encoding"

import (
	"bytes"
	"encoding/gob"
)

type Codec struct{}

func (c Codec) Marshal(v interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c Codec) Unmarshal(data []byte, v interface{}) error {
	dec := gob.NewDecoder(bytes.NewReader(data))
	return dec.Decode(v)
}

func (c Codec) ContentType() string {
	return "application/gob"
}

func init() {
	encoding.Register("gob", Codec{})
}
