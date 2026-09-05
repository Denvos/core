package errors

import (
	"testing"
)

func TestError(t *testing.T) {
	err := New("TEST", "test message")
	if err.Code() != "TEST" {
		t.Errorf("expected TEST, got %s", err.Code())
	}
	if err.Message() != "test message" {
		t.Errorf("expected test message, got %s", err.Message())
	}
}

func TestWrap(t *testing.T) {
	base := New("BASE", "base error")
	wrapped := Wrap(base, "wrapped message")
	if wrapped.Cause() != base {
		t.Error("cause not set correctly")
	}
}
