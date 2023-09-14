package main

import (
	"errors"
	"fmt"
	"reflect"
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

func assertInvalidRowCountError(t *testing.T, got, want error) {
	if got == nil {
		t.Fatal("no error returned with an invalid row count")
	}
	assertError(t, got, want)
}

func assertInvalidColCountError(t *testing.T, got, want error) {
	if got == nil {
		t.Fatal("no error returned with an invalid column count")
	}
	assertError(t, got, want)
}

func assertEmptyCharsetError(t *testing.T, got, want error) {
	if got == nil {
		t.Fatal("no error returned with an empty charset")
	}
	assertError(t, got, want)
}

func assertIncorrectArgsError(t *testing.T, got, want error) {
	if got == nil {
		t.Fatal("no error returned with incorrect template args")
	}
	assertError(t, got, want)
}

func assertInvalidTypeError(t *testing.T, got, want error) {
	if got == nil {
		t.Fatal("no error returned with invalid type")
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

func assertEmptySlice[T any](t *testing.T, got []T) {
	if len(got) > 0 {
		t.Errorf("got %+v but should have gotten an empty slice", got)
	}
}

func assertNil(t *testing.T, got any) {
	if got != nil {
		t.Errorf("got %+v but should have gotten nil", got)
	}
}

func assertZeroed(t *testing.T, got any) {
	// TODO: Use reflection for this.
	switch got := got.(type) {
	case int:
	case float64:
		assertZero(t, got)
	case string:
		assertEmptyString(t, got)
	case bool:
		assertFalse(t, got)
	case []any:
	case [][]any:
	case []int:
	case [][]int:
	case []float64:
	case [][]float64:
	case []string:
	case [][]string:
	case []bool:
	case [][]bool:
		assertEmptySlice(t, got)
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

func assertWildPrimitiveVector(t *testing.T, got any, wantMinLength, wantMaxLength int) {
	v := reflect.ValueOf(got)
	r := make([]any, v.Len())
	for i := 0; i < v.Len(); i++ {
		r[i] = v.Index(i).Interface()
	}
	if len(r) < wantMinLength || len(r) > wantMaxLength {
		t.Fatalf(
			"vector has a length of %d but should have a length in the range [%d,%d]",
			len(r),
			wantMinLength,
			wantMaxLength,
		)
	}
	for _, v := range r {
		assertWildPrimitive(t, v)
	}
}

func assertWildPrimitiveMatrix(t *testing.T, got any, wantMinRowCount, wantMaxRowCount, wantMinColCount, wantMaxColCount int) {
	v := reflect.ValueOf(got)
	r := make([][]any, v.Len())
	for i := 0; i < v.Len(); i++ {
		vv := reflect.ValueOf(v.Index(i).Interface())
		r[i] = make([]any, vv.Len())
		for j := 0; j < v.Index(i).Len(); j++ {
			r[i][j] = v.Index(i).Index(j).Interface()
		}
	}
	if len(r) < wantMinRowCount || len(r) > wantMaxRowCount {
		t.Fatalf(
			"matrix has a row count of %d but should have a row count in the range [%d,%d]",
			len(r),
			wantMinRowCount,
			wantMaxRowCount,
		)
	}
	for _, r := range r {
		assertWildPrimitiveVector(t, r, wantMinColCount, wantMaxColCount)
	}
}
