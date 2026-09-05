package semver

import (
	"fmt"
	"strconv"
	"strings"
)

type Semver struct {
	Major      int
	Minor      int
	Patch      int
	Prerelease string
	Build      string
}

func Parse(v string) (*Semver, error) {
	v = strings.TrimPrefix(v, "v")
	parts := strings.Split(v, "-")
	versionPart := parts[0]
	prerelease := ""
	if len(parts) > 1 {
		prerelease = parts[1]
	}
	build := ""
	if idx := strings.Index(prerelease, "+"); idx != -1 {
		build = prerelease[idx+1:]
		prerelease = prerelease[:idx]
	}
	nums := strings.Split(versionPart, ".")
	if len(nums) < 2 {
		return nil, fmt.Errorf("invalid semver: %s", v)
	}
	major, _ := strconv.Atoi(nums[0])
	minor, _ := strconv.Atoi(nums[1])
	patch := 0
	if len(nums) > 2 {
		patch, _ = strconv.Atoi(nums[2])
	}
	return &Semver{
		Major:      major,
		Minor:      minor,
		Patch:      patch,
		Prerelease: prerelease,
		Build:      build,
	}, nil
}

func (s *Semver) String() string {
	base := fmt.Sprintf("%d.%d", s.Major, s.Minor)
	if s.Patch > 0 {
		base += fmt.Sprintf(".%d", s.Patch)
	}
	if s.Prerelease != "" {
		base += "-" + s.Prerelease
	}
	if s.Build != "" {
		base += "+" + s.Build
	}
	return base
}

func (s *Semver) Compare(other *Semver) int {
	if s.Major != other.Major {
		if s.Major > other.Major {
			return 1
		}
		return -1
	}
	if s.Minor != other.Minor {
		if s.Minor > other.Minor {
			return 1
		}
		return -1
	}
	if s.Patch != other.Patch {
		if s.Patch > other.Patch {
			return 1
		}
		return -1
	}
	// Prerelease: if one is empty, the non-empty is considered lower
	if s.Prerelease == "" && other.Prerelease != "" {
		return 1
	}
	if s.Prerelease != "" && other.Prerelease == "" {
		return -1
	}
	if s.Prerelease != "" && other.Prerelease != "" {
		if s.Prerelease < other.Prerelease {
			return -1
		}
		if s.Prerelease > other.Prerelease {
			return 1
		}
	}
	return 0
}

func (s *Semver) Equal(other *Semver) bool {
	return s.Compare(other) == 0
}
