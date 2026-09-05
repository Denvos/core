package console

import (
	"fmt"
	"os"
)

type Console struct {
	Color bool
}

func New() *Console {
	return &Console{Color: true}
}

func (c *Console) Write(p []byte) (int, error) {
	return os.Stdout.Write(p)
}

func Write(c *Console, entry map[string]interface{}, level int) {
	// Simple color mapping
	color := ""
	reset := "\033[0m"
	if c.Color {
		switch level {
		case 1: // debug
			color = "\033[36m" // cyan
		case 2: // info
			color = "\033[32m" // green
		case 3: // warn
			color = "\033[33m" // yellow
		case 4: // error
			color = "\033[31m" // red
		}
	}
	msg := fmt.Sprintf("%s %s %s", entry["time"], entry["level"], entry["msg"])
	if color != "" {
		fmt.Fprintf(os.Stdout, "%s%s%s\n", color, msg, reset)
	} else {
		fmt.Fprintln(os.Stdout, msg)
	}
}
