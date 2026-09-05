package regex

import (
	"fmt"
	"regexp"
)

type Validator struct {
	Pattern *regexp.Regexp
}

func New(pattern string) (*Validator, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return &Validator{Pattern: re}, nil
}

func (v *Validator) Validate(input string) error {
	if input == "" {
		return fmt.Errorf("input cannot be empty")
	}
	if !v.Pattern.MatchString(input) {
		return fmt.Errorf("input does not match pattern")
	}
	return nil
}
