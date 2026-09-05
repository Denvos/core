package logger

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/Denvos/core/logger/console"
	"github.com/Denvos/core/logger/file"
	"github.com/Denvos/core/logger/json"
	"github.com/Denvos/core/logger/structured"
)

type Level int

const (
	SilentLevel Level = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
)

type Logger struct {
	mu      sync.Mutex
	level   Level
	outputs []io.Writer
	fields  map[string]interface{}
}

type Option func(*Logger)

func New(opts ...Option) *Logger {
	l := &Logger{
		level:   InfoLevel,
		outputs: []io.Writer{os.Stdout},
		fields:  make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

func WithLevel(level Level) Option {
	return func(l *Logger) {
		l.level = level
	}
}

func WithOutput(w io.Writer) Option {
	return func(l *Logger) {
		l.outputs = append(l.outputs, w)
	}
}

func WithFields(fields map[string]interface{}) Option {
	return func(l *Logger) {
		for k, v := range fields {
			l.fields[k] = v
		}
	}
}

func (l *Logger) Debug(msg string, fields ...interface{}) {
	l.log(DebugLevel, "DEBUG", msg, fields...)
}

func (l *Logger) Info(msg string, fields ...interface{}) {
	l.log(InfoLevel, "INFO", msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...interface{}) {
	l.log(WarnLevel, "WARN", msg, fields...)
}

func (l *Logger) Error(msg string, fields ...interface{}) {
	l.log(ErrorLevel, "ERROR", msg, fields...)
}

func (l *Logger) log(level Level, levelStr, msg string, fields ...interface{}) {
	if level < l.level {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()

	entry := map[string]interface{}{
		"time":  time.Now().UTC().Format(time.RFC3339),
		"level": levelStr,
		"msg":   msg,
	}
	for k, v := range l.fields {
		entry[k] = v
	}
	for i := 0; i+1 < len(fields); i += 2 {
		key, ok := fields[i].(string)
		if !ok {
			continue
		}
		entry[key] = fields[i+1]
	}

	for _, out := range l.outputs {
		switch o := out.(type) {
		case *console.Console:
			console.Write(o, entry, level)
		case *file.File:
			file.Write(o, entry)
		case *json.JSON:
			json.Write(o, entry)
		case *structured.Structured:
			structured.Write(o, entry)
		default:
			// fallback: simple text
			out.Write([]byte(levelStr + " " + msg + "\n"))
		}
	}
}

func (l *Logger) WithFields(fields map[string]interface{}) *Logger {
	newLogger := &Logger{
		level:   l.level,
		outputs: l.outputs,
		fields:  make(map[string]interface{}),
	}
	for k, v := range l.fields {
		newLogger.fields[k] = v
	}
	for k, v := range fields {
		newLogger.fields[k] = v
	}
	return newLogger
}
