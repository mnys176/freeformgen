package main

import (
	"errors"
	"fmt"
	"testing"
)

type nullDirectiveTester struct{}

func (tester nullDirectiveTester) assertNil() func(*testing.T) {
	return func(t *testing.T) {
		nullDirective()
	}
}

type integerDirectiveTester struct {
	iMin int
	iMax int
	oErr error
}

func (tester integerDirectiveTester) assertNumber() func(*testing.T) {
	return func(t *testing.T) {
		got, oErr := integerDirective(tester.iMin, tester.iMax)
		assertZeroed(t, oErr)
		assertWildNumber(t, got, tester.iMin, tester.iMax)
	}
}

func (tester integerDirectiveTester) assertMinGreaterThanMaxError() func(*testing.T) {
	return func(t *testing.T) {
		oInt, got := integerDirective(tester.iMin, tester.iMax)
		assertZeroed(t, oInt)
		assertMinGreaterThanMaxError(t, got, tester.oErr)
	}
}

type floatDirectiveTester struct {
	iMin float64
	iMax float64
	oErr error
}

func (tester floatDirectiveTester) assertNumber() func(*testing.T) {
	return func(t *testing.T) {
		got, oErr := floatDirective(tester.iMin, tester.iMax)
		assertZeroed(t, oErr)
		assertWildNumber(t, got, tester.iMin, tester.iMax)
	}
}

func (tester floatDirectiveTester) assertMinGreaterThanMaxError() func(*testing.T) {
	return func(t *testing.T) {
		oFloat, got := floatDirective(tester.iMin, tester.iMax)
		assertZeroed(t, oFloat)
		assertMinGreaterThanMaxError(t, got, tester.oErr)
	}
}

type stringDirectiveTester struct {
	iMinLength int
	iMaxLength int
	iCharset   string
	oErr       error
}

func (tester stringDirectiveTester) assertString() func(*testing.T) {
	return func(t *testing.T) {
		got, oErr := stringDirective(tester.iMinLength, tester.iMaxLength, tester.iCharset)
		assertZeroed(t, oErr)
		assertWildString(t, got, tester.iMinLength, tester.iMaxLength)
	}
}

func (tester stringDirectiveTester) assertMinGreaterThanMaxError() func(*testing.T) {
	return func(t *testing.T) {
		oString, got := stringDirective(tester.iMinLength, tester.iMaxLength, tester.iCharset)
		assertZeroed(t, oString)
		assertMinGreaterThanMaxError(t, got, tester.oErr)
	}
}

func (tester stringDirectiveTester) assertInvalidLengthError() func(*testing.T) {
	return func(t *testing.T) {
		oString, got := stringDirective(tester.iMinLength, tester.iMaxLength, tester.iCharset)
		assertZeroed(t, oString)
		assertInvalidLengthError(t, got, tester.oErr)
	}
}

type booleanDirectiveTester struct{}

func (tester booleanDirectiveTester) assertBoolean() func(*testing.T) {
	return func(t *testing.T) {
		booleanDirective()
	}
}

type primitiveDirectiveTester struct{}

func (tester primitiveDirectiveTester) assertPrimitive() func(*testing.T) {
	return func(t *testing.T) {
		got := primitiveDirective()
		assertWildPrimitive(t, got)
	}
}

type vPrimitiveDirectiveTester struct {
	iMinLength int
	iMaxLength int
	oErr       error
}

func (tester vPrimitiveDirectiveTester) assertVector() func(*testing.T) {
	return func(t *testing.T) {
		got, oErr := vPrimitiveDirective(tester.iMinLength, tester.iMaxLength)
		assertZeroed(t, oErr)
		assertWildVector(t, got, tester.iMinLength, tester.iMaxLength)
	}
}

func (tester vPrimitiveDirectiveTester) assertMinGreaterThanMaxError() func(*testing.T) {
	return func(t *testing.T) {
		oVector, got := vPrimitiveDirective(tester.iMinLength, tester.iMaxLength)
		assertZeroed(t, oVector)
		assertMinGreaterThanMaxError(t, got, tester.oErr)
	}
}

func (tester vPrimitiveDirectiveTester) assertInvalidLengthError() func(*testing.T) {
	return func(t *testing.T) {
		oVector, got := vPrimitiveDirective(tester.iMinLength, tester.iMaxLength)
		assertZeroed(t, oVector)
		assertInvalidLengthError(t, got, tester.oErr)
	}
}

const limit int = 100

func TestNullDirective(t *testing.T) {
	t.Run("baseline", nullDirectiveTester{}.assertNil())
}

func TestIntegerDirective(t *testing.T) {
	for i := 0; i < limit; i++ {
		t.Run(fmt.Sprintf("%d baseline", i), integerDirectiveTester{
			iMin: 0,
			iMax: 1,
		}.assertNumber())
		t.Run(fmt.Sprintf("%d broad range", i), integerDirectiveTester{
			iMin: -10,
			iMax: 10,
		}.assertNumber())
		t.Run(fmt.Sprintf("%d equal", i), integerDirectiveTester{
			iMin: 0,
			iMax: 0,
		}.assertNumber())
	}
	t.Run("min greater than max", integerDirectiveTester{
		iMin: 1,
		iMax: -1,
		oErr: errors.New("freeformgen: min cannot exceed max"),
	}.assertMinGreaterThanMaxError())
}

func TestFloatDirective(t *testing.T) {
	for i := 0; i < limit; i++ {
		t.Run(fmt.Sprintf("%d baseline", i), floatDirectiveTester{
			iMin: 0.0,
			iMax: 1.0,
		}.assertNumber())
		t.Run(fmt.Sprintf("%d broad range", i), floatDirectiveTester{
			iMin: -10.0,
			iMax: 10.0,
		}.assertNumber())
		t.Run(fmt.Sprintf("%d equal", i), floatDirectiveTester{
			iMin: 0.0,
			iMax: 0.0,
		}.assertNumber())
	}
	t.Run("min greater than max", floatDirectiveTester{
		iMin: 10.0,
		iMax: -10.0,
		oErr: errors.New("freeformgen: min cannot exceed max"),
	}.assertMinGreaterThanMaxError())
}

func TestStringDirective(t *testing.T) {
	for i := 0; i < limit; i++ {
		t.Run(fmt.Sprintf("%d baseline", i), stringDirectiveTester{
			iMinLength: 3,
			iMaxLength: 6,
			iCharset:   "abc",
		}.assertString())
		t.Run(fmt.Sprintf("%d emojis", i), stringDirectiveTester{
			iMinLength: 3,
			iMaxLength: 6,
			iCharset:   "🔴🟡🟢",
		}.assertString())
		t.Run(fmt.Sprintf("%d length of zero", i), stringDirectiveTester{
			iMinLength: 0,
			iMaxLength: 0,
			iCharset:   "abc",
		}.assertString())
		t.Run(fmt.Sprintf("%d equal lengths", i), stringDirectiveTester{
			iMinLength: 6,
			iMaxLength: 6,
			iCharset:   "abc",
		}.assertString())
	}
	t.Run("invalid min length", stringDirectiveTester{
		iMinLength: -1,
		iMaxLength: 6,
		oErr:       errors.New("freeformgen: string cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("invalid max length", stringDirectiveTester{
		iMinLength: 3,
		iMaxLength: -1,
		oErr:       errors.New("freeformgen: string cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("min length greater than max length", stringDirectiveTester{
		iMinLength: 6,
		iMaxLength: 3,
		oErr:       errors.New("freeformgen: min length cannot exceed max length"),
	}.assertMinGreaterThanMaxError())
}

func TestBooleanDirective(t *testing.T) {
	t.Run("baseline", booleanDirectiveTester{}.assertBoolean())
}

func TestPrimitiveDirective(t *testing.T) {
	for i := 0; i < limit; i++ {
		t.Run(fmt.Sprintf("%d baseline", i), primitiveDirectiveTester{}.assertPrimitive())
	}
}

func TestVPrimitiveDirective(t *testing.T) {
	for i := 0; i < limit; i++ {
		t.Run(fmt.Sprintf("%d baseline", i), vPrimitiveDirectiveTester{
			iMinLength: 3,
			iMaxLength: 6,
		}.assertVector())
		t.Run(fmt.Sprintf("%d length of zero", i), vPrimitiveDirectiveTester{
			iMinLength: 0,
			iMaxLength: 0,
		}.assertVector())
		t.Run(fmt.Sprintf("%d equal lengths", i), vPrimitiveDirectiveTester{
			iMinLength: 6,
			iMaxLength: 6,
		}.assertVector())
	}
	t.Run("invalid min length", vPrimitiveDirectiveTester{
		iMinLength: -1,
		iMaxLength: 6,
		oErr:       errors.New("freeformgen: vector cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("invalid max length", vPrimitiveDirectiveTester{
		iMinLength: 3,
		iMaxLength: -1,
		oErr:       errors.New("freeformgen: vector cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("min length greater than max length", vPrimitiveDirectiveTester{
		iMinLength: 6,
		iMaxLength: 3,
		oErr:       errors.New("freeformgen: min length cannot exceed max length"),
	}.assertMinGreaterThanMaxError())
}
