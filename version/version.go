package version

import (
	"fmt"
	"strconv"
	"strings"
)

var Version = "0.1"

type Version struct {
	Major int
	Minor int
}

func Parse(v string) (*Version, error) {
	v = strings.TrimPrefix(v, "v")
	parts := strings.Split(v, ".")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid version format: %s", v)
	}
	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("invalid major version: %s", parts[0])
	}
	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid minor version: %s", parts[1])
	}
	return &Version{Major: major, Minor: minor}, nil
}

func (v *Version) String() string {
	return fmt.Sprintf("%d.%d", v.Major, v.Minor)
}

func (v *Version) Compare(other *Version) int {
	if v.Major != other.Major {
		if v.Major > other.Major {
			return 1
		}
		return -1
	}
	if v.Minor != other.Minor {
		if v.Minor > other.Minor {
			return 1
		}
		return -1
	}
	return 0
}

func (v *Version) Equal(other *Version) bool {
	return v.Compare(other) == 0
}

func (v *Version) GreaterThan(other *Version) bool {
	return v.Compare(other) > 0
}

func (v *Version) LessThan(other *Version) bool {
	return v.Compare(other) < 0
}

func (v *Version) IsZero() bool {
	return v.Major == 0 && v.Minor == 0
}

func (v *Version) NextMajor() *Version {
	return &Version{Major: v.Major + 1, Minor: 0}
}

func (v *Version) NextMinor() *Version {
	return &Version{Major: v.Major, Minor: v.Minor + 1}
}
