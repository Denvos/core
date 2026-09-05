package structured

import (
	"encoding/json"
	"io"
)

type Structured struct {
	w io.Writer
}

func New(w io.Writer) *Structured {
	return &Structured{w: w}
}

func (s *Structured) Write(p []byte) (int, error) {
	return s.w.Write(p)
}

func Write(s *Structured, entry map[string]interface{}) {
	data, _ := json.Marshal(entry)
	s.w.Write(append(data, '\n'))
}
