package encoding

import (
	"testing"

	"github.com/Denvos/core/encoding/xml"
)

type xmlTestStruct struct {
	Name string `xml:"name"`
}

func TestXMLCodec(t *testing.T) {
	c := xml.Codec{}
	v := xmlTestStruct{Name: "test"}

	data, err := c.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}

	var out xmlTestStruct
	if err := c.Unmarshal(data, &out); err != nil {
		t.Fatal(err)
	}

	if out.Name != "test" {
		t.Errorf("expected test, got %s", out.Name)
	}
}
