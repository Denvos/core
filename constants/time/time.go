package time

import "time"

const (
	Nanosecond  = 1
	Microsecond = 1000 * Nanosecond
	Millisecond = 1000 * Microsecond
	Second      = 1000 * Millisecond
	Minute      = 60 * Second
	Hour        = 60 * Minute
	Day         = 24 * Hour
	Week        = 7 * Day
)

const (
	DateFormatISO     = "2006-01-02"
	DateFormatRFC3339 = time.RFC3339
	DateFormatRFC822  = time.RFC822
	DateFormatRFC850  = time.RFC850
	DateFormatRFC1123 = time.RFC1123
	DateFormatKitchen = "3:04PM"
)

const (
	LayoutANSIC       = time.ANSIC
	LayoutUnixDate    = time.UnixDate
	LayoutRubyDate    = time.RubyDate
	LayoutRFC822      = time.RFC822
	LayoutRFC822Z     = time.RFC822Z
	LayoutRFC850      = time.RFC850
	LayoutRFC1123     = time.RFC1123
	LayoutRFC1123Z    = time.RFC1123Z
	LayoutRFC3339     = time.RFC3339
	LayoutRFC3339Nano = time.RFC3339Nano
	LayoutKitchen     = time.Kitchen
	LayoutStamp       = time.Stamp
	LayoutStampMilli  = time.StampMilli
	LayoutStampMicro  = time.StampMicro
	LayoutStampNano   = time.StampNano
)

var (
	DefaultTimeout = 30 * Second
	ShortTimeout   = 5 * Second
	LongTimeout    = 5 * Minute
)
