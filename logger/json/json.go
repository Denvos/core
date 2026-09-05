package json

import (
	"encoding/json"
	"io"
)

type JSON struct {
	w io.Writer
}

func New(w io.Writer) *JSON {
	return &JSON{w: w}
}

func (j *JSON) Write(p []byte) (int, error) {
	return j.w.Write(p)
}

func Write(j *JSON, entry map[string]interface{}) {
	data, _ := json.Marshal(entry)
	j.w.Write(append(data, '\n'))
}
