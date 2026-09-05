package yaml

import (
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
