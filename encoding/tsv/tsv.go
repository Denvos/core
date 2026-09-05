package tsv

import "github.com/Denvos/core/encoding"

import (
	"bytes"
	"encoding/csv"
)

type Codec struct{}

func (c Codec) Marshal(v interface{}) ([]byte, error) {
	records, ok := v.([][]string)
	if !ok {
		return nil, nil
	}
	buf := new(bytes.Buffer)
	w := csv.NewWriter(buf)
	w.Comma = '\t'
	for _, r := range records {
		if err := w.Write(r); err != nil {
			return nil, err
		}
	}
	w.Flush()
	return buf.Bytes(), w.Error()
}

func (c Codec) Unmarshal(data []byte, v interface{}) error {
	records, ok := v.(*[][]string)
	if !ok {
		return nil
	}
	r := csv.NewReader(bytes.NewReader(data))
	r.Comma = '\t'
	rows, err := r.ReadAll()
	if err != nil {
		return err
	}
	*records = rows
	return nil
}

func (c Codec) ContentType() string {
	return "text/tab-separated-values"
}

func init() {
	encoding.Register("tsv", Codec{})
}
