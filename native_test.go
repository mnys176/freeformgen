package main

import (
	"testing"
)

type nullTester struct{}

func (tester nullTester) assertNil() func(*testing.T) {
	return func(t *testing.T) {
		got := null()
		assertNil(t, got)
	}
}

func TestNull(t *testing.T) {
	t.Run("baseline", nullTester{}.assertNil())
}
