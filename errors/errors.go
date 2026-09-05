package errors

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/Denvos/core/errors/stack"
)

type Error struct {
	code    string
	message string
	cause   error
	stack   []uintptr
	fields  map[string]interface{}
}

type Option func(*Error)

func New(code, message string, opts ...Option) *Error {
	e := &Error{
		code:    code,
		message: message,
		stack:   stack.Capture(2),
		fields:  make(map[string]interface{}),
	}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

func Wrap(err error, message string, opts ...Option) *Error {
	if err == nil {
		return nil
	}
	e := New("", message, opts...)
	e.cause = err
	return e
}

func WithCause(cause error) Option {
	return func(e *Error) {
		e.cause = cause
	}
}

func WithField(key string, value interface{}) Option {
	return func(e *Error) {
		e.fields[key] = value
	}
}

func (e *Error) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("%s: %v", e.message, e.cause)
	}
	return e.message
}

func (e *Error) Code() string {
	return e.code
}

func (e *Error) Message() string {
	return e.message
}

func (e *Error) Cause() error {
	return e.cause
}

func (e *Error) Stack() []uintptr {
	return e.stack
}

func (e *Error) Fields() map[string]interface{} {
	return e.fields
}

func (e *Error) Unwrap() error {
	return e.cause
}

func Is(err error, code string) bool {
	e, ok := err.(*Error)
	if !ok {
		return false
	}
	return e.code == code
}
