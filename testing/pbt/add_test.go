package pbt_test

import (
	"testing"
	"testing/quick"

	. "github.com/aranw/go_examples/testing/pbt"
)

func TestAdd(t *testing.T) {
	assertion := func(a, b int) bool {
		return Add(a, b) == Add(b, a)
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error(err)
	}
}

func TestAddToZero(t *testing.T) {
	assertion := func(a int) bool {
		return Add(a, 0) == Add(0, a)
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error(err)
	}
}
