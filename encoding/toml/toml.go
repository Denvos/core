package toml

import "github.com/Denvos/core/encoding"

import (
	"github.com/pelletier/go-toml/v2"
)

type Codec struct{}

func (c Codec) Marshal(v interface{}) ([]byte, error) {
	return toml.Marshal(v)
}

func (c Codec) Unmarshal(data []byte, v interface{}) error {
	return toml.Unmarshal(data, v)
}

func (c Codec) ContentType() string {
	return "application/toml"
}

func init() {
	encoding.Register("toml", Codec{})
}
