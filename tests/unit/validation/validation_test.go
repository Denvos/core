package validation

import (
	"testing"

	"github.com/Denvos/core/validation/email"
	"github.com/Denvos/core/validation/url"
)

func TestEmail(t *testing.T) {
	v := email.New()
	if err := v.Validate("test@example.com"); err != nil {
		t.Error(err)
	}
	if err := v.Validate("invalid"); err == nil {
		t.Error("expected invalid email to fail")
	}
}

func TestURL(t *testing.T) {
	v := url.New()
	if err := v.Validate("https://example.com"); err != nil {
		t.Error(err)
	}
	if err := v.Validate("example"); err == nil {
		t.Error("expected invalid URL to fail")
	}
}
