package grpc

const (
	DefaultMaxRecvMsgSize = 4 * 1024 * 1024 // 4MB
	DefaultMaxSendMsgSize = 4 * 1024 * 1024 // 4MB
)

const (
	CodeOK                 = 0
	CodeCanceled           = 1
	CodeUnknown            = 2
	CodeInvalidArgument    = 3
	CodeDeadlineExceeded   = 4
	CodeNotFound           = 5
	CodeAlreadyExists      = 6
	CodePermissionDenied   = 7
	CodeResourceExhausted  = 8
	CodeFailedPrecondition = 9
	CodeAborted            = 10
	CodeOutOfRange         = 11
	CodeUnimplemented      = 12
	CodeInternal           = 13
	CodeUnavailable        = 14
	CodeDataLoss           = 15
	CodeUnauthenticated    = 16
)

var CodeText = map[int]string{
	0:  "OK",
	1:  "Canceled",
	2:  "Unknown",
	3:  "Invalid Argument",
	4:  "Deadline Exceeded",
	5:  "Not Found",
	6:  "Already Exists",
	7:  "Permission Denied",
	8:  "Resource Exhausted",
	9:  "Failed Precondition",
	10: "Aborted",
	11: "Out Of Range",
	12: "Unimplemented",
	13: "Internal",
	14: "Unavailable",
	15: "Data Loss",
	16: "Unauthenticated",
}
