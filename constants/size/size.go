package size

const (
	Byte     = 1
	Kilobyte = 1024 * Byte
	Megabyte = 1024 * Kilobyte
	Gigabyte = 1024 * Megabyte
	Terabyte = 1024 * Gigabyte
	Petabyte = 1024 * Terabyte
)

const (
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
)

var (
	DefaultBufferSize = 64 * Kilobyte
	LargeBufferSize   = 1 * Megabyte
	MaxFileSize       = 100 * Megabyte
)
