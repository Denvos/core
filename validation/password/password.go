package password

import (
	"fmt"
	"unicode"
)

type Validator struct {
	MinLength    int
	RequireUpper bool
	RequireLower bool
	RequireDigit bool
	RequireSymbol bool
}

func New() *Validator {
	return &Validator{
		MinLength:    8,
		RequireUpper: true,
		RequireLower: true,
		RequireDigit: true,
		RequireSymbol: true,
	}
}

func (v *Validator) Validate(password string) error {
	if len(password) < v.MinLength {
		return fmt.Errorf("password must be at least %d characters", v.MinLength)
	}
	var hasUpper, hasLower, hasDigit, hasSymbol bool
	for _, ch := range password {
		switch {
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsDigit(ch):
			hasDigit = true
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			hasSymbol = true
		}
	}
	if v.RequireUpper && !hasUpper {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}
	if v.RequireLower && !hasLower {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}
	if v.RequireDigit && !hasDigit {
		return fmt.Errorf("password must contain at least one digit")
	}
	if v.RequireSymbol && !hasSymbol {
		return fmt.Errorf("password must contain at least one symbol")
	}
	return nil
}
