package encoding

import (
	"testing"

	"github.com/Denvos/core/encoding/yaml"
)

type testStruct struct {
	Name string `yaml:"name"`
}

func TestYAMLCodec(t *testing.T) {
	c := yaml.Codec{}
	v := testStruct{Name: "test"}

	data, err := c.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}

	var out testStruct
	if err := c.Unmarshal(data, &out); err != nil {
		t.Fatal(err)
	}

	if out.Name != "test" {
		t.Errorf("expected test, got %s", out.Name)
	}
}
