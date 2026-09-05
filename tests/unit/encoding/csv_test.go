package encoding

import (
    "testing"

    "github.com/Denvos/core/encoding/csv"
)

func TestCSVCodec(t *testing.T) {
    c := csv.Codec{}
    v := [][]string{{"name", "age"}, {"test", "30"}}

    data, err := c.Marshal(v)
    if err != nil {
        t.Fatal(err)
    }

    var out [][]string
    if err := c.Unmarshal(data, &out); err != nil {
        t.Fatal(err)
    }

    if len(out) != 2 {
        t.Errorf("expected 2 rows, got %d", len(out))
    }
    if out[0][0] != "name" {
        t.Errorf("expected name, got %s", out[0][0])
    }
}
