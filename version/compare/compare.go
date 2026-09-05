package compare

import "github.com/Denvos/core/version"

func Compare(a, b *version.Version) int {
	if a == nil && b == nil {
		return 0
	}
	if a == nil {
		return -1
	}
	if b == nil {
		return 1
	}
	return a.Compare(b)
}

func Equal(a, b *version.Version) bool {
	return Compare(a, b) == 0
}

func Greater(a, b *version.Version) bool {
	return Compare(a, b) > 0
}

func Less(a, b *version.Version) bool {
	return Compare(a, b) < 0
}

func Min(a, b *version.Version) *version.Version {
	if Less(a, b) {
		return a
	}
	return b
}

func Max(a, b *version.Version) *version.Version {
	if Greater(a, b) {
		return a
	}
	return b
}
