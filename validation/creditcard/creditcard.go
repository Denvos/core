package creditcard

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Denvos/core/validation"
)

type Validator struct {
	CheckLuhn bool
}

func New() *Validator {
	return &Validator{CheckLuhn: true}
}

func (v *Validator) Validate(card string) error {
	if card == "" {
		return fmt.Errorf("credit card cannot be empty")
	}
	// Remove spaces and dashes
	clean := strings.ReplaceAll(strings.ReplaceAll(card, " ", ""), "-", "")
	if !validation.CreditCardRegex.MatchString(clean) {
		return fmt.Errorf("invalid credit card format")
	}
	if v.CheckLuhn {
		if !luhn(clean) {
			return fmt.Errorf("invalid credit card number (Luhn check failed)")
		}
	}
	return nil
}

func luhn(card string) bool {
	sum := 0
	alternate := false
	for i := len(card) - 1; i >= 0; i-- {
		n, err := strconv.Atoi(string(card[i]))
		if err != nil {
			return false
		}
		if alternate {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
		alternate = !alternate
	}
	return sum%10 == 0
}
