package logger

import (
	"bytes"
	"testing"
)

func TestLogger(t *testing.T) {
	var buf bytes.Buffer
	l := New(WithOutput(&buf), WithLevel(DebugLevel))
	l.Info("test message", "key", "value")
	if buf.Len() == 0 {
		t.Error("expected log output")
	}
	if !bytes.Contains(buf.Bytes(), []byte("test message")) {
		t.Error("expected message in output")
	}
}
