package stack

import (
	"runtime"
	"strings"
)

func Capture(skip int) []uintptr {
	pc := make([]uintptr, 32)
	n := runtime.Callers(skip+1, pc)
	return pc[:n]
}

func Format(pcs []uintptr) string {
	if len(pcs) == 0 {
		return ""
	}
	var b strings.Builder
	frames := runtime.CallersFrames(pcs)
	for {
		frame, more := frames.Next()
		if !more {
			break
		}
		b.WriteString(frame.File)
		b.WriteString(":")
		b.WriteString(frame.Line)
		b.WriteString(" ")
		b.WriteString(frame.Function)
		b.WriteString("\n")
	}
	return b.String()
}
