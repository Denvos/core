package mac

import (
	"fmt"

	"github.com/Denvos/core/validation"
)

type Validator struct{}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) Validate(mac string) error {
	if mac == "" {
		return fmt.Errorf("MAC address cannot be empty")
	}
	if !validation.MACRegex.MatchString(mac) {
		return fmt.Errorf("invalid MAC address format: %s", mac)
	}
	return nil
}
