package avro

import "github.com/Denvos/core/encoding"

type Codec struct{}

func (c Codec) Marshal(v interface{}) ([]byte, error) {
	return nil, nil
}

func (c Codec) Unmarshal(data []byte, v interface{}) error {
	return nil
}

func (c Codec) ContentType() string {
	return "application/avro"
}

func init() {
	encoding.Register("avro", Codec{})
}
