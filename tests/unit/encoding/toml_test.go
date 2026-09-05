package encoding

import (
	"testing"

	"github.com/Denvos/core/encoding/toml"
)

type tomlTestStruct struct {
	Name string `toml:"name"`
}

func TestTOMLCodec(t *testing.T) {
	c := toml.Codec{}
	v := tomlTestStruct{Name: "test"}

	data, err := c.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}

	var out tomlTestStruct
	if err := c.Unmarshal(data, &out); err != nil {
		t.Fatal(err)
	}

	if out.Name != "test" {
		t.Errorf("expected test, got %s", out.Name)
	}
}
