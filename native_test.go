package main

import (
	"errors"
	"fmt"
	"testing"
)

type nullDirectiveTester struct{}

func (tester nullDirectiveTester) assertNil() func(*testing.T) {
	return func(t *testing.T) {
		got := nullDirective()
		assertNil(t, got)
	}
}

type integerDirectiveTester struct {
	iMin   int
	iMax   int
	oPanic error
}

func (tester integerDirectiveTester) assertNumber() func(*testing.T) {
	return func(t *testing.T) {
		got := integerDirective(tester.iMin, tester.iMax)
		assertNumber(t, got, tester.iMin, tester.iMax)
	}
}

func (tester integerDirectiveTester) assertMinGreaterThanMaxErrorPanic() func(*testing.T) {
	return func(t *testing.T) {
		defer assertMinGreaterThanMaxPanic(t, tester.oPanic)
		integerDirective(tester.iMin, tester.iMax)
	}
}

type floatDirectiveTester struct {
	iMin   float64
	iMax   float64
	oPanic error
}

func (tester floatDirectiveTester) assertNumber() func(*testing.T) {
	return func(t *testing.T) {
		got := floatDirective(tester.iMin, tester.iMax)
		assertNumber(t, got, tester.iMin, tester.iMax)
	}
}

func (tester floatDirectiveTester) assertMinGreaterThanMaxErrorPanic() func(*testing.T) {
	return func(t *testing.T) {
		defer assertMinGreaterThanMaxPanic(t, tester.oPanic)
		floatDirective(tester.iMin, tester.iMax)
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
		assertString(t, got, tester.iMinLength, tester.iMaxLength)
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
		assertPrimitive(t, got)
	}
}

const limit int = 100

func TestNullDirective(t *testing.T) {
	for i := 0; i < limit; i++ {
		t.Run(fmt.Sprintf("%d baseline", i), nullDirectiveTester{}.assertNil())
	}
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
		t.Run(fmt.Sprintf("%d min greater than max", i), integerDirectiveTester{
			iMin:   1,
			iMax:   -1,
			oPanic: errors.New("freeformgen: min cannot exceed max"),
		}.assertMinGreaterThanMaxErrorPanic())
	}
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
		t.Run(fmt.Sprintf("%d min greater than max", i), floatDirectiveTester{
			iMin:   10.0,
			iMax:   -10.0,
			oPanic: errors.New("freeformgen: min cannot exceed max"),
		}.assertMinGreaterThanMaxErrorPanic())
	}
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
		t.Run(fmt.Sprintf("%d invalid min length", i), stringDirectiveTester{
			iMinLength: -1,
			iMaxLength: 6,
			oPanic:     errors.New("freeformgen: string cannot have a negative length"),
		}.assertInvalidLengthErrorPanic())
		t.Run(fmt.Sprintf("%d invalid max length", i), stringDirectiveTester{
			iMinLength: 3,
			iMaxLength: -1,
			oPanic:     errors.New("freeformgen: string cannot have a negative length"),
		}.assertInvalidLengthErrorPanic())
		t.Run(fmt.Sprintf("%d min length greater than max length", i), stringDirectiveTester{
			iMinLength: 6,
			iMaxLength: 3,
			oPanic:     errors.New("freeformgen: min length cannot exceed max length"),
		}.assertMinGreaterThanMaxErrorPanic())
	}
}

func TestBooleanDirective(t *testing.T) {
	t.Run("baseline", booleanDirectiveTester{}.assertBoolean())
}

func TestPrimitiveDirective(t *testing.T) {
	for i := 0; i < limit; i++ {
		t.Run(fmt.Sprintf("%d baseline", i), primitiveDirectiveTester{}.assertPrimitive())
	}
}
