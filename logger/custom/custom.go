package custom

import "io"

type Custom struct {
	w   io.Writer
	fn  func([]byte) (int, error)
}

func New(w io.Writer, fn func([]byte) (int, error)) *Custom {
	return &Custom{w: w, fn: fn}
}

func (c *Custom) Write(p []byte) (int, error) {
	if c.fn != nil {
		return c.fn(p)
	}
	return c.w.Write(p)
}
