package time

import (
	"fmt"
	"time"
)

type Validator struct {
	Format string
}

func New() *Validator {
	return &Validator{Format: "15:04:05Z07:00"}
}

func (v *Validator) Validate(timestr string) error {
	if timestr == "" {
		return fmt.Errorf("time cannot be empty")
	}
	_, err := time.Parse(v.Format, timestr)
	if err != nil {
		return fmt.Errorf("invalid time format (expected %s): %s", v.Format, timestr)
	}
	return nil
}
