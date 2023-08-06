package main

import (
	"errors"
	"testing"
)

type nullTester struct{}

func (tester nullTester) assertNil() func(*testing.T) {
	return func(t *testing.T) {
		got := null()
		assertNil(t, got)
	}
}

type numberTester[T int | float64] struct {
	iMin   T
	iMax   T
	oPanic error
}

func (tester numberTester[T]) assertNumber() func(*testing.T) {
	return func(t *testing.T) {
		got := number[T](tester.iMin, tester.iMax)
		assertNumber(t, got, tester.iMin, tester.iMax)
	}
}

func (tester numberTester[T]) assertMinGreaterThanMaxErrorPanic() func(*testing.T) {
	return func(t *testing.T) {
		defer assertMinGreaterThanMaxPanic(t, tester.oPanic)
		number[T](tester.iMin, tester.iMax)
	}
}

type strTester struct {
	iMinLength int
	iMaxLength int
	iCharset   string
	oPanic     error
}

func (tester strTester) assertStr() func(*testing.T) {
	return func(t *testing.T) {
		got := str(tester.iMinLength, tester.iMaxLength, tester.iCharset)
		assertStr(t, got, tester.iMinLength, tester.iMaxLength, tester.iCharset)
	}
}

func (tester strTester) assertMinGreaterThanMaxErrorPanic() func(*testing.T) {
	return func(t *testing.T) {
		defer assertMinGreaterThanMaxPanic(t, tester.oPanic)
		str(tester.iMinLength, tester.iMaxLength, tester.iCharset)
	}
}

func (tester strTester) assertInvalidLengthErrorPanic() func(*testing.T) {
	return func(t *testing.T) {
		defer assertInvalidLengthPanic(t, tester.oPanic)
		str(tester.iMinLength, tester.iMaxLength, tester.iCharset)
	}
}

func TestNull(t *testing.T) {
	t.Run("baseline", nullTester{}.assertNil())
}

func TestNumber(t *testing.T) {
	t.Run("baseline", numberTester[float64]{
		iMin: 0.0,
		iMax: 1.0,
	}.assertNumber())
	t.Run("broad range", numberTester[float64]{
		iMin: -10.0,
		iMax: 10.0,
	}.assertNumber())
	t.Run("equal", numberTester[float64]{
		iMin: 0.0,
		iMax: 0.0,
	}.assertNumber())
	t.Run("min greater than max", numberTester[float64]{
		iMin:   10.0,
		iMax:   -10.0,
		oPanic: errors.New("freeformgen: min cannot exceed max"),
	}.assertMinGreaterThanMaxErrorPanic())
	t.Run("int baseline", numberTester[int]{
		iMin: 0,
		iMax: 1,
	}.assertNumber())
	t.Run("int broad range", numberTester[int]{
		iMin: -10,
		iMax: 10,
	}.assertNumber())
	t.Run("int equal", numberTester[int]{
		iMin: 0,
		iMax: 0,
	}.assertNumber())
	t.Run("int min greater than max", numberTester[int]{
		iMin:   1,
		iMax:   -1,
		oPanic: errors.New("freeformgen: min cannot exceed max"),
	}.assertMinGreaterThanMaxErrorPanic())
}

func TestStr(t *testing.T) {
	t.Run("baseline", strTester{
		iMinLength: 3,
		iMaxLength: 6,
		iCharset:   "abc",
	}.assertStr())
	t.Run("emojis", strTester{
		iMinLength: 3,
		iMaxLength: 6,
		iCharset:   "🔴🟡🟢",
	}.assertStr())
	t.Run("length of zero", strTester{
		iMinLength: 0,
		iMaxLength: 0,
		iCharset:   "abc",
	}.assertStr())
	t.Run("equal lengths", strTester{
		iMinLength: 6,
		iMaxLength: 6,
		iCharset:   "abc",
	}.assertStr())
	t.Run("invalid min length", strTester{
		iMinLength: -1,
		iMaxLength: 6,
		oPanic:     errors.New("freeformgen: string cannot have a negative length"),
	}.assertInvalidLengthErrorPanic())
	t.Run("invalid max length", strTester{
		iMinLength: 3,
		iMaxLength: -1,
		oPanic:     errors.New("freeformgen: string cannot have a negative length"),
	}.assertInvalidLengthErrorPanic())
	t.Run("min length greater than max length", strTester{
		iMinLength: 6,
		iMaxLength: 3,
		oPanic:     errors.New("freeformgen: min length cannot exceed max length"),
	}.assertMinGreaterThanMaxErrorPanic())
}
