package validate

import (
	"regexp"
)

var pattern = regexp.MustCompile(`^v?\d+\.\d+$`)

func IsValid(v string) bool {
	return pattern.MatchString(v)
}

func Must(v string) bool {
	return IsValid(v)
}
