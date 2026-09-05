package json

import "github.com/Denvos/core/encoding"

import (
	"encoding/json"
)

type Codec struct{}

func (c Codec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (c Codec) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (c Codec) ContentType() string {
	return "application/json"
}

func init() {
	encoding.Register("json", Codec{})
}
