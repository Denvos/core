package username

import (
	"fmt"

	"github.com/Denvos/core/validation"
)

type Validator struct {
	MinLength int
	MaxLength int
}

func New() *Validator {
	return &Validator{MinLength: 3, MaxLength: 32}
}

func (v *Validator) Validate(username string) error {
	if len(username) < v.MinLength || len(username) > v.MaxLength {
		return fmt.Errorf("username must be between %d and %d characters", v.MinLength, v.MaxLength)
	}
	if !validation.UsernameRegex.MatchString(username) {
		return fmt.Errorf("username can only contain letters, numbers, and underscores")
	}
	return nil
}
