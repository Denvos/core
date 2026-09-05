package codes

const (
	OK                  = "OK"
	Canceled            = "CANCELED"
	Unknown             = "UNKNOWN"
	InvalidArgument     = "INVALID_ARGUMENT"
	DeadlineExceeded    = "DEADLINE_EXCEEDED"
	NotFound            = "NOT_FOUND"
	AlreadyExists       = "ALREADY_EXISTS"
	PermissionDenied    = "PERMISSION_DENIED"
	ResourceExhausted   = "RESOURCE_EXHAUSTED"
	FailedPrecondition  = "FAILED_PRECONDITION"
	Aborted             = "ABORTED"
	OutOfRange          = "OUT_OF_RANGE"
	Unimplemented       = "UNIMPLEMENTED"
	Internal            = "INTERNAL"
	Unavailable         = "UNAVAILABLE"
	DataLoss            = "DATA_LOSS"
	Unauthenticated     = "UNAUTHENTICATED"
	Conflict            = "CONFLICT"
	TooManyRequests     = "TOO_MANY_REQUESTS"
	Timeout             = "TIMEOUT"
	Retryable           = "RETRYABLE"
	Fatal               = "FATAL"
)

var DefaultHTTPStatus = map[string]int{
	OK:                  200,
	Canceled:            499,
	Unknown:             500,
	InvalidArgument:     400,
	DeadlineExceeded:    504,
	NotFound:            404,
	AlreadyExists:       409,
	PermissionDenied:    403,
	ResourceExhausted:   429,
	FailedPrecondition:  400,
	Aborted:             409,
	OutOfRange:          400,
	Unimplemented:       501,
	Internal:            500,
	Unavailable:         503,
	DataLoss:            500,
	Unauthenticated:     401,
	Conflict:            409,
	TooManyRequests:     429,
	Timeout:             504,
	Retryable:           503,
	Fatal:               500,
}
