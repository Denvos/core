package logging

type Logger interface {
    Debug(msg string, fields ...interface{})
    Info(msg string, fields ...interface{})
    Warn(msg string, fields ...interface{})
    Error(msg string, fields ...interface{})
}

type loggerFunc struct {
    fn func(msg string, fields ...interface{})
}

func (l loggerFunc) Debug(msg string, fields ...interface{}) {
    l.fn(msg, fields...)
}

func (l loggerFunc) Info(msg string, fields ...interface{}) {
    l.fn(msg, fields...)
}

func (l loggerFunc) Warn(msg string, fields ...interface{}) {
    l.fn(msg, fields...)
}

func (l loggerFunc) Error(msg string, fields ...interface{}) {
    l.fn(msg, fields...)
}

func LoggerFunc(fn func(msg string, fields ...interface{})) Logger {
    return loggerFunc{fn: fn}
}

type noopLogger struct{}

func (n noopLogger) Debug(msg string, fields ...interface{}) {}
func (n noopLogger) Info(msg string, fields ...interface{})  {}
func (n noopLogger) Warn(msg string, fields ...interface{})  {}
func (n noopLogger) Error(msg string, fields ...interface{}) {}

func NoopLogger() Logger {
    return noopLogger{}
}
