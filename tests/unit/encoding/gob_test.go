package encoding

import (
    "testing"

    "github.com/Denvos/core/encoding/gob"
)

type gobTestStruct struct {
    Name string
}

func TestGOBCodec(t *testing.T) {
    c := gob.Codec{}
    v := gobTestStruct{Name: "test"}

    data, err := c.Marshal(v)
    if err != nil {
        t.Fatal(err)
    }

    var out gobTestStruct
    if err := c.Unmarshal(data, &out); err != nil {
        t.Fatal(err)
    }

    if out.Name != "test" {
        t.Errorf("expected test, got %s", out.Name)
    }
}
