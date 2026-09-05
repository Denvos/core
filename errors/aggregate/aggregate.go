package aggregate

import (
	"strings"
)

type Aggregate []error

func (a Aggregate) Error() string {
	var msgs []string
	for _, err := range a {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

func (a Aggregate) Unwrap() []error {
	return a
}
