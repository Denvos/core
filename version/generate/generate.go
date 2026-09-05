package generate

import (
	"fmt"

	"github.com/Denvos/core/version"
)

func NextMajor(v *version.Version) string {
	return v.NextMajor().String()
}

func NextMinor(v *version.Version) string {
	return v.NextMinor().String()
}

func FromParts(major, minor int) string {
	return fmt.Sprintf("%d.%d", major, minor)
}
