package version

import "testing"

func TestParse(t *testing.T) {
	v, err := Parse("0.1")
	if err != nil {
		t.Fatal(err)
	}
	if v.Major != 0 || v.Minor != 1 {
		t.Errorf("expected 0.1, got %d.%d", v.Major, v.Minor)
	}
}

func TestCompare(t *testing.T) {
	a, _ := Parse("0.1")
	b, _ := Parse("0.2")
	if !a.LessThan(b) {
		t.Error("0.1 should be less than 0.2")
	}
	if !b.GreaterThan(a) {
		t.Error("0.2 should be greater than 0.1")
	}
}

func TestNext(t *testing.T) {
	v, _ := Parse("0.1")
	next := v.NextMinor()
	if next.String() != "0.2" {
		t.Errorf("expected 0.2, got %s", next.String())
	}
	nextMajor := v.NextMajor()
	if nextMajor.String() != "1.0" {
		t.Errorf("expected 1.0, got %s", nextMajor.String())
	}
}
