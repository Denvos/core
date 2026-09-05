package encoding

import (
    "testing"

    "github.com/Denvos/core/encoding/bson"
)

func TestBSONCodec(t *testing.T) {
    c := bson.Codec{}
    data, err := c.Marshal(nil)
    if err != nil {
        t.Fatal(err)
    }
    if data != nil {
        t.Error("expected nil for invalid input")
    }
}
