package encoding

import (
	"testing"

	"github.com/Denvos/core/encoding/json"
)

type jsonTestStruct struct {
	Name string `json:"name"`
}

func TestJSONCodec(t *testing.T) {
	c := json.Codec{}
	v := jsonTestStruct{Name: "test"}

	data, err := c.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}

	var out jsonTestStruct
	if err := c.Unmarshal(data, &out); err != nil {
		t.Fatal(err)
	}

	if out.Name != "test" {
		t.Errorf("expected test, got %s", out.Name)
	}
}
