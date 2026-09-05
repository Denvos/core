package debug

import (
	"fmt"
	"io"
	"time"
)

type Debug struct {
	w io.Writer
}

func New(w io.Writer) *Debug {
	return &Debug{w: w}
}

func (d *Debug) Write(msg string, fields map[string]interface{}) {
	timestamp := time.Now().UTC().Format(time.RFC3339)
	fmt.Fprintf(d.w, "[%s] DEBUG %s\n", timestamp, msg)
}
