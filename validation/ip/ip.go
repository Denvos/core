package ip

import (
	"fmt"
	"net"
)

type Validator struct {
	IPv4 bool
	IPv6 bool
}

func New() *Validator {
	return &Validator{IPv4: true, IPv6: true}
}

func (v *Validator) Validate(ip string) error {
	if ip == "" {
		return fmt.Errorf("IP cannot be empty")
	}
	parsed := net.ParseIP(ip)
	if parsed == nil {
		return fmt.Errorf("invalid IP address: %s", ip)
	}
	if v.IPv4 && parsed.To4() != nil {
		return nil
	}
	if v.IPv6 && parsed.To16() != nil && parsed.To4() == nil {
		return nil
	}
	return fmt.Errorf("IP version not allowed")
}
