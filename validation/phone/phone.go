package phone

import (
	"fmt"

	"github.com/Denvos/core/validation"
)

type Validator struct{}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) Validate(phone string) error {
	if phone == "" {
		return fmt.Errorf("phone cannot be empty")
	}
	if !validation.PhoneRegex.MatchString(phone) {
		return fmt.Errorf("invalid phone number format: %s", phone)
	}
	return nil
}
