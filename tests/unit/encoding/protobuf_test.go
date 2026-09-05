package encoding

import (
	"testing"

	"github.com/Denvos/core/encoding/protobuf"
)

func TestProtobufCodec(t *testing.T) {
	c := protobuf.Codec{}
	data, err := c.Marshal(nil)
	if err != nil {
		t.Fatal(err)
	}
	if data != nil {
		t.Error("expected nil for invalid input")
	}
}
