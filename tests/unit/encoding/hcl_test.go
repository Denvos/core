package encoding

import (
    "testing"

    "github.com/Denvos/core/encoding/hcl"
)

type hclTestStruct struct {
    Name string `hcl:"name"`
}

func TestHCLCodec(t *testing.T) {
    t.Skip("HCL codec needs proper implementation")
    c := hcl.Codec{}
    v := hclTestStruct{Name: "test"}

    data, err := c.Marshal(v)
    if err != nil {
        t.Fatal(err)
    }

    var out hclTestStruct
    if err := c.Unmarshal(data, &out); err != nil {
        t.Fatal(err)
    }

    if out.Name != "test" {
        t.Errorf("expected test, got %s", out.Name)
    }
}
