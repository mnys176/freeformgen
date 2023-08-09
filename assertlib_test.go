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

func assertZero[T int | float64](t *testing.T, got T) {
	if float64(got) != 0.0 {
		t.Errorf("got %v but should have gotten %v", got, T(0.0))
	}
}

func assertEmptyString(t *testing.T, got string) {
	if got != "" {
		t.Errorf("got %q but should have gotten an empty string", got)
	}
}

func assertFalse(t *testing.T, got bool) {
	if got != false {
		t.Errorf("got %t but should have gotten false", got)
	}
}

func assertNil(t *testing.T, got any) {
	if got != nil {
		t.Errorf("got %+v but should have gotten nil", got)
	}
}

func assertZeroed(t *testing.T, got any) {
	switch got := got.(type) {
	case int:
	case float64:
		assertZero(t, got)
	case string:
		assertEmptyString(t, got)
	case bool:
		assertFalse(t, got)
	default:
		assertNil(t, got)
	}
}

func assertWildNumber[T int | float64](t *testing.T, got, wantMin, wantMax T) {
	if got < wantMin || got > wantMax {
		t.Errorf("got %v but should have gotten a number in the range [%v,%v]", got, wantMin, wantMax)
	}
}

func assertWildString(t *testing.T, got string, wantMinLength, wantMaxLength int) {
	pattern := fmt.Sprintf(`[\p{S}\p{L}\p{N}]{%d,%d}`, wantMinLength, wantMaxLength)
	if matches, _ := regexp.MatchString(pattern, got); !matches {
		t.Errorf("got %q which does not match the pattern %q", got, pattern)
	}
}

func assertWildPrimitive(t *testing.T, got any) {
	switch got.(type) {
	case int:
	case float64:
	case string:
	case bool:
	case nil:
		break
	default:
		t.Errorf("got value of type %T which is not a primitive", got)
	}
}
