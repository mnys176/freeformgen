package main

import (
	"errors"
	"fmt"
	"regexp"
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

func assertInvalidLengthError(t *testing.T, got, want error) {
	if got == nil {
		t.Fatal("no error returned with an invalid length")
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

func assertInvalidLengthPanic(t *testing.T, want error) {
	r := recover()
	if r == nil {
		t.Fatal("no panic with an invalid length")
	}
	assertInvalidLengthError(t, r.(error), want)
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

func assertString(t *testing.T, got string, wantMinLength, wantMaxLength int, wantCharset string) {
	pattern := fmt.Sprintf("[%s]{%d,%d}", wantCharset, wantMinLength, wantMaxLength)
	if matches, _ := regexp.MatchString(pattern, got); !matches {
		t.Errorf("got %q which does not match the pattern %q", got, pattern)
	}
}
