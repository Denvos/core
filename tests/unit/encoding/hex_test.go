package encoding

import (
	"testing"

	"github.com/Denvos/core/encoding/hex"
)

func TestHexCodec(t *testing.T) {
	c := hex.Codec{}
	v := []byte("hello")

	data, err := c.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}

	var out []byte
	if err := c.Unmarshal(data, &out); err != nil {
		t.Fatal(err)
	}

	if string(out) != "hello" {
		t.Errorf("expected hello, got %s", string(out))
	}
}
