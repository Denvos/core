package prerelease

import (
	"fmt"
	"strings"
)

type Prerelease struct {
	Label    string
	Version  string
	Metadata string
}

func Parse(s string) (*Prerelease, error) {
	if s == "" {
		return nil, nil
	}
	parts := strings.SplitN(s, "+", 2)
	label := parts[0]
	metadata := ""
	if len(parts) > 1 {
		metadata = parts[1]
	}
	return &Prerelease{
		Label:    label,
		Metadata: metadata,
	}, nil
}

func (p *Prerelease) String() string {
	if p == nil {
		return ""
	}
	s := p.Label
	if p.Metadata != "" {
		s += "+" + p.Metadata
	}
	return s
}
