package encoding

import (
    "testing"

    "github.com/Denvos/core/encoding/msgpack"
)

func TestMsgpackCodec(t *testing.T) {
    c := msgpack.Codec{}
    data, err := c.Marshal(nil)
    if err != nil {
        t.Fatal(err)
    }
    if data != nil {
        t.Error("expected nil for invalid input")
    }
}
