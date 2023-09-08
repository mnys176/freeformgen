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

func assertWildNullVector(t *testing.T, got []any, wantMinLength, wantMaxLength int) {
	if len(got) < wantMinLength || len(got) > wantMaxLength {
		t.Fatalf(
			"vector has a length of %d but should have a length in the range [%d,%d]",
			len(got),
			wantMinLength,
			wantMaxLength,
		)
	}
	for _, v := range got {
		assertNil(t, v)
	}
}

func assertWildNumberVector[T int | float64](t *testing.T, got []T, wantMinLength, wantMaxLength int, wantMin, wantMax T) {
	if len(got) < wantMinLength || len(got) > wantMaxLength {
		t.Fatalf(
			"vector has a length of %d but should have a length in the range [%d,%d]",
			len(got),
			wantMinLength,
			wantMaxLength,
		)
	}
	for _, v := range got {
		assertWildNumber(t, v, wantMin, wantMax)
	}
}

func assertWildStringVector(t *testing.T, got []string, wantMinLength, wantMaxLength, wantMinStrLength, wantMaxStrLength int) {
	if len(got) < wantMinLength || len(got) > wantMaxLength {
		t.Fatalf(
			"vector has a length of %d but should have a length in the range [%d,%d]",
			len(got),
			wantMinLength,
			wantMaxLength,
		)
	}
	for _, v := range got {
		assertWildString(t, v, wantMinStrLength, wantMaxStrLength)
	}
}

func assertWildBooleanVector(t *testing.T, got []bool, wantMinLength, wantMaxLength int) {
	if len(got) < wantMinLength || len(got) > wantMaxLength {
		t.Fatalf(
			"vector has a length of %d but should have a length in the range [%d,%d]",
			len(got),
			wantMinLength,
			wantMaxLength,
		)
	}
}

func assertWildPrimitiveVector(t *testing.T, got []any, wantMinLength, wantMaxLength int) {
	if len(got) < wantMinLength || len(got) > wantMaxLength {
		t.Fatalf(
			"vector has a length of %d but should have a length in the range [%d,%d]",
			len(got),
			wantMinLength,
			wantMaxLength,
		)
	}
	for _, v := range got {
		assertWildPrimitive(t, v)
	}
}

func assertWildNullMatrix(t *testing.T, got [][]any, wantMinRowCount, wantMaxRowCount, wantMinColCount, wantMaxColCount int) {
	if len(got) < wantMinRowCount || len(got) > wantMaxRowCount {
		t.Fatalf(
			"matrix has a row count of %d but should have a row count in the range [%d,%d]",
			len(got),
			wantMinRowCount,
			wantMaxRowCount,
		)
	}
	for _, r := range got {
		assertWildNullVector(t, r, wantMinColCount, wantMaxColCount)
	}
}

func assertWildNumberMatrix[T int | float64](t *testing.T, got [][]T, wantMinRowCount, wantMaxRowCount, wantMinColCount, wantMaxColCount int, wantMin, wantMax T) {
	if len(got) < wantMinRowCount || len(got) > wantMaxRowCount {
		t.Fatalf(
			"matrix has a row count of %d but should have a row count in the range [%d,%d]",
			len(got),
			wantMinRowCount,
			wantMaxRowCount,
		)
	}
	for _, r := range got {
		assertWildNumberVector(t, r, wantMinColCount, wantMaxColCount, wantMin, wantMax)
	}
}

func assertWildStringMatrix(t *testing.T, got [][]string, wantMinRowCount, wantMaxRowCount, wantMinColCount, wantMaxColCount, wantMinStrLength, wantMaxStrLength int) {
	if len(got) < wantMinRowCount || len(got) > wantMaxRowCount {
		t.Fatalf(
			"matrix has a row count of %d but should have a row count in the range [%d,%d]",
			len(got),
			wantMinRowCount,
			wantMaxRowCount,
		)
	}
	for _, r := range got {
		assertWildStringVector(t, r, wantMinColCount, wantMaxColCount, wantMinStrLength, wantMaxStrLength)
	}
}

func assertWildBooleanMatrix(t *testing.T, got [][]bool, wantMinRowCount, wantMaxRowCount, wantMinColCount, wantMaxColCount int) {
	if len(got) < wantMinRowCount || len(got) > wantMaxRowCount {
		t.Fatalf(
			"matrix has a row count of %d but should have a row count in the range [%d,%d]",
			len(got),
			wantMinRowCount,
			wantMaxRowCount,
		)
	}
	for _, r := range got {
		assertWildBooleanVector(t, r, wantMinColCount, wantMaxColCount)
	}
}

func assertWildPrimitiveMatrix(t *testing.T, got [][]any, wantMinRowCount, wantMaxRowCount, wantMinColCount, wantMaxColCount int) {
	if len(got) < wantMinRowCount || len(got) > wantMaxRowCount {
		t.Fatalf(
			"matrix has a row count of %d but should have a row count in the range [%d,%d]",
			len(got),
			wantMinRowCount,
			wantMaxRowCount,
		)
	}
	for _, r := range got {
		assertWildPrimitiveVector(t, r, wantMinColCount, wantMaxColCount)
	}
}
