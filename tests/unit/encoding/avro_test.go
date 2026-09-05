package encoding

import (
    "testing"

    "github.com/Denvos/core/encoding/avro"
)

func TestAvroCodec(t *testing.T) {
    c := avro.Codec{}
    data, err := c.Marshal(nil)
    if err != nil {
        t.Fatal(err)
    }
    if data != nil {
        t.Error("expected nil for invalid input")
    }
}
