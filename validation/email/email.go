package email

import (
	"fmt"
	"net"
	"strings"

	"github.com/Denvos/core/validation"
)

type Validator struct {
	CheckDomain bool
}

func New() *Validator {
	return &Validator{CheckDomain: false}
}

func (v *Validator) Validate(email string) error {
	if email == "" {
		return fmt.Errorf("email cannot be empty")
	}
	if !validation.EmailRegex.MatchString(email) {
		return fmt.Errorf("invalid email format: %s", email)
	}
	if v.CheckDomain {
		parts := strings.Split(email, "@")
		if len(parts) != 2 {
			return fmt.Errorf("invalid email structure")
		}
		domain := parts[1]
		mx, err := net.LookupMX(domain)
		if err != nil || len(mx) == 0 {
			return fmt.Errorf("domain %s has no MX records", domain)
		}
	}
	return nil
}
