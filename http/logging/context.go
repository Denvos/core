package logging

import (
    "context"
)

type ctxKey struct{}

var key ctxKey

func WithLogger(ctx context.Context, logger Logger) context.Context {
    return context.WithValue(ctx, key, logger)
}

func FromContext(ctx context.Context) Logger {
    logger, ok := ctx.Value(key).(Logger)
    if !ok {
        return NoopLogger()
    }
    return logger
}

func Debug(ctx context.Context, msg string, fields ...interface{}) {
    FromContext(ctx).Debug(msg, fields...)
}

func Info(ctx context.Context, msg string, fields ...interface{}) {
    FromContext(ctx).Info(msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...interface{}) {
    FromContext(ctx).Warn(msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...interface{}) {
    FromContext(ctx).Error(msg, fields...)
}
