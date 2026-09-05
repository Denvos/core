package uuid

import (
	"fmt"

	"github.com/Denvos/core/validation"
)

type Validator struct{}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) Validate(uuid string) error {
	if uuid == "" {
		return fmt.Errorf("UUID cannot be empty")
	}
	if !validation.UUIDRegex.MatchString(uuid) {
		return fmt.Errorf("invalid UUID format: %s", uuid)
	}
	return nil
}
