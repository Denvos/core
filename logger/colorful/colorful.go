package colorful

import (
	"fmt"
	"io"
)

type Colorful struct {
	w io.Writer
}

func New(w io.Writer) *Colorful {
	return &Colorful{w: w}
}

func (c *Colorful) Write(level string, msg string, fields map[string]interface{}) {
	colorMap := map[string]string{
		"DEBUG": "\033[36m",
		"INFO":  "\033[32m",
		"WARN":  "\033[33m",
		"ERROR": "\033[31m",
	}
	reset := "\033[0m"
	color := colorMap[level]
	if color == "" {
		color = reset
	}
	fmt.Fprintf(c.w, "%s%s: %s%s\n", color, level, msg, reset)
}
