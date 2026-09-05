package hostname

import (
	"fmt"

	"github.com/Denvos/core/validation"
)

type Validator struct{}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) Validate(hostname string) error {
	if hostname == "" {
		return fmt.Errorf("hostname cannot be empty")
	}
	if len(hostname) > 253 {
		return fmt.Errorf("hostname too long (max 253 characters)")
	}
	if !validation.HostnameRegex.MatchString(hostname) {
		return fmt.Errorf("invalid hostname format: %s", hostname)
	}
	return nil
}
