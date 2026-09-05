package health

import (
    "context"
    "time"
)

type Status string

const (
    StatusPass  Status = "pass"
    StatusFail  Status = "fail"
    StatusWarn  Status = "warn"
    StatusUnknown Status = "unknown"
)

type Checker interface {
    Name() string
    Check(ctx context.Context) (*Result, error)
}

type Result struct {
    Status  Status                 `json:"status"`
    Message string                 `json:"message,omitempty"`
    Data    map[string]interface{} `json:"data,omitempty"`
    Time    time.Time              `json:"time"`
    Error   error                  `json:"-"`
}

type CheckFunc func(ctx context.Context) (*Result, error)

func (f CheckFunc) Name() string {
    return "check"
}

func (f CheckFunc) Check(ctx context.Context) (*Result, error) {
    return f(ctx)
}
