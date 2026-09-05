package formatting

import (
	"fmt"
	"strings"

	"github.com/Denvos/core/errors"
)

func Format(e *errors.Error, verbose bool) string {
	if e == nil {
		return ""
	}
	var b strings.Builder
	b.WriteString(fmt.Sprintf("[%s] %s", e.Code(), e.Message()))
	if cause := e.Cause(); cause != nil {
		b.WriteString(fmt.Sprintf(" (caused by: %v)", cause))
	}
	if verbose {
		if fields := e.Fields(); len(fields) > 0 {
			b.WriteString(" | fields: ")
			for k, v := range fields {
				b.WriteString(fmt.Sprintf("%s=%v ", k, v))
			}
		}
		if stack := e.Stack(); len(stack) > 0 {
			b.WriteString("\nstack:\n")
			b.WriteString(stack.Format(stack))
		}
	}
	return b.String()
}
