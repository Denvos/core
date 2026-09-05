package date

import (
	"fmt"
	"time"
)

type Validator struct {
	Format string
}

func New() *Validator {
	return &Validator{Format: "2006-01-02"}
}

func (v *Validator) Validate(date string) error {
	if date == "" {
		return fmt.Errorf("date cannot be empty")
	}
	_, err := time.Parse(v.Format, date)
	if err != nil {
		return fmt.Errorf("invalid date format (expected %s): %s", v.Format, date)
	}
	return nil
}
