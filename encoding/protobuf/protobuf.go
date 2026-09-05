package protobuf

import "github.com/Denvos/core/encoding"

import (
	"google.golang.org/protobuf/proto"
)

type Codec struct{}

func (c Codec) Marshal(v interface{}) ([]byte, error) {
	msg, ok := v.(proto.Message)
	if !ok {
		return nil, nil
	}
	return proto.Marshal(msg)
}

func (c Codec) Unmarshal(data []byte, v interface{}) error {
	msg, ok := v.(proto.Message)
	if !ok {
		return nil
	}
	return proto.Unmarshal(data, msg)
}

func (c Codec) ContentType() string {
	return "application/protobuf"
}

func init() {
	encoding.Register("protobuf", Codec{})
}
