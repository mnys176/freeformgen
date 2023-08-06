package main

import (
	"errors"
	"testing"
)

type nullDirectiveTester struct{}

func (tester nullDirectiveTester) assertNil() func(*testing.T) {
	return func(t *testing.T) {
		got := nullDirective()
		assertNil(t, got)
	}
}

type numberDirectiveTester[T int | float64] struct {
	iMin   T
	iMax   T
	oPanic error
}

func (tester numberDirectiveTester[T]) assertNumber() func(*testing.T) {
	return func(t *testing.T) {
		got := number[T](tester.iMin, tester.iMax)
		assertNumber(t, got, tester.iMin, tester.iMax)
	}
}

func (tester numberDirectiveTester[T]) assertMinGreaterThanMaxErrorPanic() func(*testing.T) {
	return func(t *testing.T) {
		defer assertMinGreaterThanMaxPanic(t, tester.oPanic)
		number[T](tester.iMin, tester.iMax)
	}
}

type stringDirectiveTester struct {
	iMinLength int
	iMaxLength int
	iCharset   string
	oPanic     error
}

func (tester stringDirectiveTester) assertString() func(*testing.T) {
	return func(t *testing.T) {
		got := stringDirective(tester.iMinLength, tester.iMaxLength, tester.iCharset)
		assertString(t, got, tester.iMinLength, tester.iMaxLength, tester.iCharset)
	}
}

func (tester stringDirectiveTester) assertMinGreaterThanMaxErrorPanic() func(*testing.T) {
	return func(t *testing.T) {
		defer assertMinGreaterThanMaxPanic(t, tester.oPanic)
		stringDirective(tester.iMinLength, tester.iMaxLength, tester.iCharset)
	}
}

func (tester stringDirectiveTester) assertInvalidLengthErrorPanic() func(*testing.T) {
	return func(t *testing.T) {
		defer assertInvalidLengthPanic(t, tester.oPanic)
		stringDirective(tester.iMinLength, tester.iMaxLength, tester.iCharset)
	}
}

func TestNullDirective(t *testing.T) {
	t.Run("baseline", nullDirectiveTester{}.assertNil())
}

func TestNumberDirective(t *testing.T) {
	t.Run("baseline", numberDirectiveTester[float64]{
		iMin: 0.0,
		iMax: 1.0,
	}.assertNumber())
	t.Run("broad range", numberDirectiveTester[float64]{
		iMin: -10.0,
		iMax: 10.0,
	}.assertNumber())
	t.Run("equal", numberDirectiveTester[float64]{
		iMin: 0.0,
		iMax: 0.0,
	}.assertNumber())
	t.Run("min greater than max", numberDirectiveTester[float64]{
		iMin:   10.0,
		iMax:   -10.0,
		oPanic: errors.New("freeformgen: min cannot exceed max"),
	}.assertMinGreaterThanMaxErrorPanic())
	t.Run("int baseline", numberDirectiveTester[int]{
		iMin: 0,
		iMax: 1,
	}.assertNumber())
	t.Run("int broad range", numberDirectiveTester[int]{
		iMin: -10,
		iMax: 10,
	}.assertNumber())
	t.Run("int equal", numberDirectiveTester[int]{
		iMin: 0,
		iMax: 0,
	}.assertNumber())
	t.Run("int min greater than max", numberDirectiveTester[int]{
		iMin:   1,
		iMax:   -1,
		oPanic: errors.New("freeformgen: min cannot exceed max"),
	}.assertMinGreaterThanMaxErrorPanic())
}

func TestStringDirective(t *testing.T) {
	t.Run("baseline", stringDirectiveTester{
		iMinLength: 3,
		iMaxLength: 6,
		iCharset:   "abc",
	}.assertString())
	t.Run("emojis", stringDirectiveTester{
		iMinLength: 3,
		iMaxLength: 6,
		iCharset:   "🔴🟡🟢",
	}.assertString())
	t.Run("length of zero", stringDirectiveTester{
		iMinLength: 0,
		iMaxLength: 0,
		iCharset:   "abc",
	}.assertString())
	t.Run("equal lengths", stringDirectiveTester{
		iMinLength: 6,
		iMaxLength: 6,
		iCharset:   "abc",
	}.assertString())
	t.Run("invalid min length", stringDirectiveTester{
		iMinLength: -1,
		iMaxLength: 6,
		oPanic:     errors.New("freeformgen: string cannot have a negative length"),
	}.assertInvalidLengthErrorPanic())
	t.Run("invalid max length", stringDirectiveTester{
		iMinLength: 3,
		iMaxLength: -1,
		oPanic:     errors.New("freeformgen: string cannot have a negative length"),
	}.assertInvalidLengthErrorPanic())
	t.Run("min length greater than max length", stringDirectiveTester{
		iMinLength: 6,
		iMaxLength: 3,
		oPanic:     errors.New("freeformgen: min length cannot exceed max length"),
	}.assertMinGreaterThanMaxErrorPanic())
}
