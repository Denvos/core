package constraints

import (
	"fmt"
	"strings"

	"github.com/Denvos/core/version"
)

type Constraint struct {
	operator string
	value    *version.Version
}

func Parse(expr string) ([]Constraint, error) {
	var constraints []Constraint
	parts := strings.Split(expr, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		var op string
		var ver string
		if strings.HasPrefix(part, ">=") {
			op = ">="
			ver = strings.TrimSpace(part[2:])
		} else if strings.HasPrefix(part, "<=") {
			op = "<="
			ver = strings.TrimSpace(part[2:])
		} else if strings.HasPrefix(part, "!=") {
			op = "!="
			ver = strings.TrimSpace(part[2:])
		} else if strings.HasPrefix(part, ">") {
			op = ">"
			ver = strings.TrimSpace(part[1:])
		} else if strings.HasPrefix(part, "<") {
			op = "<"
			ver = strings.TrimSpace(part[1:])
		} else if strings.HasPrefix(part, "=") {
			op = "="
			ver = strings.TrimSpace(part[1:])
		} else {
			op = "="
			ver = part
		}
		v, err := version.Parse(ver)
		if err != nil {
			return nil, fmt.Errorf("invalid version in constraint: %s", ver)
		}
		constraints = append(constraints, Constraint{operator: op, value: v})
	}
	return constraints, nil
}

func Check(v *version.Version, constraints []Constraint) bool {
	for _, c := range constraints {
		if !checkSingle(v, c) {
			return false
		}
	}
	return true
}

func checkSingle(v *version.Version, c Constraint) bool {
	switch c.operator {
	case "=":
		return v.Equal(c.value)
	case "!=":
		return !v.Equal(c.value)
	case ">":
		return v.GreaterThan(c.value)
	case "<":
		return v.LessThan(c.value)
	case ">=":
		return v.GreaterThan(c.value) || v.Equal(c.value)
	case "<=":
		return v.LessThan(c.value) || v.Equal(c.value)
	default:
		return false
	}
}
