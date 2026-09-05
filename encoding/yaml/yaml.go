package yaml

import (
	"github.com/Denvos/core/encoding"
	"gopkg.in/yaml.v3"
)

type Codec struct{}

func (c Codec) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (c Codec) Unmarshal(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}

func (c Codec) ContentType() string {
	return "application/yaml"
}

func init() {
	encoding.Register("yaml", Codec{})
}
