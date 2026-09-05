package metrics

import (
	"io"
	"sync/atomic"
)

type Metrics struct {
	w          io.Writer
	lineCount  uint64
	errorCount uint64
}

func New(w io.Writer) *Metrics {
	return &Metrics{w: w}
}

func (m *Metrics) Write(p []byte) (int, error) {
	atomic.AddUint64(&m.lineCount, 1)
	return m.w.Write(p)
}

func (m *Metrics) IncErrors() {
	atomic.AddUint64(&m.errorCount, 1)
}

func (m *Metrics) LineCount() uint64 {
	return atomic.LoadUint64(&m.lineCount)
}

func (m *Metrics) ErrorCount() uint64 {
	return atomic.LoadUint64(&m.errorCount)
}
