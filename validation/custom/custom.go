package custom

import "fmt"

type Validator struct {
	Fn func(interface{}) error
}

func New(fn func(interface{}) error) *Validator {
	return &Validator{Fn: fn}
}

func (v *Validator) Validate(input interface{}) error {
	if v.Fn == nil {
		return fmt.Errorf("no validation function provided")
	}
	return v.Fn(input)
}
