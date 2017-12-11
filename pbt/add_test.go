package pbt

import (
	"testing"
	"testing/quick"
)

func TestAdd(t *testing.T) {
	assertion := func(a, b int) bool {
		if Add(a, b) != Add(b, a) {
			return false
		}
		return true
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error(err)
	}
}

func TestAddToZero(t *testing.T) {
	assertion := func(a int) bool {
		if Add(a, 0) != Add(0, a) {
			return false
		}
		return true
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error(err)
	}
}
