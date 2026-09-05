package messages

const (
	DefaultUnknown      = "an unknown error occurred"
	DefaultInvalidArg   = "invalid argument provided"
	DefaultNotFound     = "resource not found"
	DefaultPermission   = "permission denied"
	DefaultTimeout      = "operation timed out"
	DefaultRetryable    = "retryable error"
	DefaultFatal        = "fatal error"
)

var Default = map[string]string{
	"OK":                  "",
	"CANCELED":            "operation was canceled",
	"UNKNOWN":             DefaultUnknown,
	"INVALID_ARGUMENT":    DefaultInvalidArg,
	"DEADLINE_EXCEEDED":   DefaultTimeout,
	"NOT_FOUND":           DefaultNotFound,
	"ALREADY_EXISTS":      "resource already exists",
	"PERMISSION_DENIED":   DefaultPermission,
	"RESOURCE_EXHAUSTED":  "resource exhausted",
	"FAILED_PRECONDITION": "failed precondition",
	"ABORTED":             "operation aborted",
	"OUT_OF_RANGE":        "out of range",
	"UNIMPLEMENTED":       "not implemented",
	"INTERNAL":            "internal server error",
	"UNAVAILABLE":         "service unavailable",
	"DATA_LOSS":           "data loss",
	"UNAUTHENTICATED":     "unauthenticated",
	"CONFLICT":            "conflict",
	"TOO_MANY_REQUESTS":   "too many requests",
	"TIMEOUT":             DefaultTimeout,
	"RETRYABLE":           DefaultRetryable,
	"FATAL":               DefaultFatal,
}
