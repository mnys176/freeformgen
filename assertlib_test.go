package main

import (
	"errors"
	"testing"
)

func assertError(t *testing.T, got, want error) {
	if !errors.Is(got, want) {
		t.Errorf("got %q error but wanted %q", got, want)
	}
}

func assertMinGreaterThanMaxError(t *testing.T, got, want error) {
	if got == nil {
		t.Fatal("no error returned with a minimum greater than a maximum")
	}
	assertError(t, got, want)
}

func assertMinGreaterThanMaxPanic(t *testing.T, want error) {
	r := recover()
	if r == nil {
		t.Fatal("no panic with a minimum greater than a maximum")
	}
	assertMinGreaterThanMaxError(t, r.(error), want)
}

func assertNil(t *testing.T, got any) {
	if got != nil {
		t.Errorf("got %+v but should have gotten nil", got)
	}
}

func assertNumber[T int | float64](t *testing.T, got, wantMin, wantMax T) {
	if got < wantMin || got > wantMax {
		t.Errorf("got %v but should have gotten a number in the range [%v,%v)", got, wantMin, wantMax)
	}
}
