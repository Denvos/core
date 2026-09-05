package port

import (
	"fmt"
	"strconv"
)

type Validator struct {
	Min int
	Max int
}

func New() *Validator {
	return &Validator{Min: 1, Max: 65535}
}

func (v *Validator) Validate(port interface{}) error {
	var p int
	switch val := port.(type) {
	case int:
		p = val
	case string:
		i, err := strconv.Atoi(val)
		if err != nil {
			return fmt.Errorf("port must be a number: %s", val)
		}
		p = i
	default:
		return fmt.Errorf("port must be int or string")
	}
	if p < v.Min || p > v.Max {
		return fmt.Errorf("port %d out of range [%d-%d]", p, v.Min, v.Max)
	}
	return nil
}
