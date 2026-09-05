package datadog

import "io"

type Datadog struct{ w io.Writer }

func New(w io.Writer) *Datadog { return &Datadog{w: w} }
func (d *Datadog) Write(p []byte) (int, error) { return d.w.Write(p) }
