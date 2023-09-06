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

func (tester stringDirectiveTester) assertEmptyCharsetError() func(*testing.T) {
	return func(t *testing.T) {
		oString, got := stringDirective(tester.iMinLength, tester.iMaxLength, tester.iCharset)
		assertZeroed(t, oString)
		assertEmptyCharsetError(t, got, tester.oErr)
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

type vNullDirectiveTester struct {
	iMinLength int
	iMaxLength int
	oErr       error
}

func (tester vNullDirectiveTester) assertVector() func(*testing.T) {
	return func(t *testing.T) {
		got, oErr := vNullDirective(tester.iMinLength, tester.iMaxLength)
		assertZeroed(t, oErr)
		assertWildNullVector(t, got, tester.iMinLength, tester.iMaxLength)
	}
}

func (tester vNullDirectiveTester) assertMinGreaterThanMaxError() func(*testing.T) {
	return func(t *testing.T) {
		oVector, got := vNullDirective(tester.iMinLength, tester.iMaxLength)
		assertZeroed(t, oVector)
		assertMinGreaterThanMaxError(t, got, tester.oErr)
	}
}

func (tester vNullDirectiveTester) assertInvalidLengthError() func(*testing.T) {
	return func(t *testing.T) {
		oVector, got := vNullDirective(tester.iMinLength, tester.iMaxLength)
		assertZeroed(t, oVector)
		assertInvalidLengthError(t, got, tester.oErr)
	}
}

type vIntegerDirectiveTester struct {
	iMinLength int
	iMaxLength int
	iMin       int
	iMax       int
	oErr       error
}

func (tester vIntegerDirectiveTester) assertVector() func(*testing.T) {
	return func(t *testing.T) {
		got, oErr := vIntegerDirective(tester.iMinLength, tester.iMaxLength, tester.iMin, tester.iMax)
		assertZeroed(t, oErr)
		assertWildNumberVector(t, got, tester.iMinLength, tester.iMaxLength, tester.iMin, tester.iMax)
	}
}

func (tester vIntegerDirectiveTester) assertMinGreaterThanMaxError() func(*testing.T) {
	return func(t *testing.T) {
		oVector, got := vIntegerDirective(tester.iMinLength, tester.iMaxLength, tester.iMin, tester.iMax)
		assertZeroed(t, oVector)
		assertMinGreaterThanMaxError(t, got, tester.oErr)
	}
}

func (tester vIntegerDirectiveTester) assertInvalidLengthError() func(*testing.T) {
	return func(t *testing.T) {
		oVector, got := vIntegerDirective(tester.iMinLength, tester.iMaxLength, tester.iMin, tester.iMax)
		assertZeroed(t, oVector)
		assertInvalidLengthError(t, got, tester.oErr)
	}
}

type vFloatDirectiveTester struct {
	iMinLength int
	iMaxLength int
	iMin       float64
	iMax       float64
	oErr       error
}

func (tester vFloatDirectiveTester) assertVector() func(*testing.T) {
	return func(t *testing.T) {
		got, oErr := vFloatDirective(tester.iMinLength, tester.iMaxLength, tester.iMin, tester.iMax)
		assertZeroed(t, oErr)
		assertWildNumberVector(t, got, tester.iMinLength, tester.iMaxLength, tester.iMin, tester.iMax)
	}
}

func (tester vFloatDirectiveTester) assertMinGreaterThanMaxError() func(*testing.T) {
	return func(t *testing.T) {
		oVector, got := vFloatDirective(tester.iMinLength, tester.iMaxLength, tester.iMin, tester.iMax)
		assertZeroed(t, oVector)
		assertMinGreaterThanMaxError(t, got, tester.oErr)
	}
}

func (tester vFloatDirectiveTester) assertInvalidLengthError() func(*testing.T) {
	return func(t *testing.T) {
		oVector, got := vFloatDirective(tester.iMinLength, tester.iMaxLength, tester.iMin, tester.iMax)
		assertZeroed(t, oVector)
		assertInvalidLengthError(t, got, tester.oErr)
	}
}

type vStringDirectiveTester struct {
	iMinLength    int
	iMaxLength    int
	iMinStrLength int
	iMaxStrLength int
	iCharset      string
	oErr          error
}

func (tester vStringDirectiveTester) assertVector() func(*testing.T) {
	return func(t *testing.T) {
		got, oErr := vStringDirective(tester.iMinLength, tester.iMaxLength, tester.iMinStrLength, tester.iMaxStrLength, tester.iCharset)
		assertZeroed(t, oErr)
		assertWildStringVector(t, got, tester.iMinLength, tester.iMaxLength, tester.iMinStrLength, tester.iMaxStrLength)
	}
}

func (tester vStringDirectiveTester) assertMinGreaterThanMaxError() func(*testing.T) {
	return func(t *testing.T) {
		oVector, got := vStringDirective(tester.iMinLength, tester.iMaxLength, tester.iMinStrLength, tester.iMaxStrLength, tester.iCharset)
		assertZeroed(t, oVector)
		assertMinGreaterThanMaxError(t, got, tester.oErr)
	}
}

func (tester vStringDirectiveTester) assertInvalidLengthError() func(*testing.T) {
	return func(t *testing.T) {
		oVector, got := vStringDirective(tester.iMinLength, tester.iMaxLength, tester.iMinStrLength, tester.iMaxStrLength, tester.iCharset)
		assertZeroed(t, oVector)
		assertInvalidLengthError(t, got, tester.oErr)
	}
}

func (tester vStringDirectiveTester) assertEmptyCharsetError() func(*testing.T) {
	return func(t *testing.T) {
		oVector, got := vStringDirective(tester.iMinLength, tester.iMaxLength, tester.iMinStrLength, tester.iMaxStrLength, tester.iCharset)
		assertZeroed(t, oVector)
		assertEmptyCharsetError(t, got, tester.oErr)
	}
}

type vBooleanDirectiveTester struct {
	iMinLength int
	iMaxLength int
	oErr       error
}

func (tester vBooleanDirectiveTester) assertVector() func(*testing.T) {
	return func(t *testing.T) {
		got, oErr := vBooleanDirective(tester.iMinLength, tester.iMaxLength)
		assertZeroed(t, oErr)
		assertWildBooleanVector(t, got, tester.iMinLength, tester.iMaxLength)
	}
}

func (tester vBooleanDirectiveTester) assertMinGreaterThanMaxError() func(*testing.T) {
	return func(t *testing.T) {
		oVector, got := vBooleanDirective(tester.iMinLength, tester.iMaxLength)
		assertZeroed(t, oVector)
		assertMinGreaterThanMaxError(t, got, tester.oErr)
	}
}

func (tester vBooleanDirectiveTester) assertInvalidLengthError() func(*testing.T) {
	return func(t *testing.T) {
		oVector, got := vBooleanDirective(tester.iMinLength, tester.iMaxLength)
		assertZeroed(t, oVector)
		assertInvalidLengthError(t, got, tester.oErr)
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
		assertWildPrimitiveVector(t, got, tester.iMinLength, tester.iMaxLength)
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

type mNullDirectiveTester struct {
	iMinRows int
	iMaxRows int
	iMinCols int
	iMaxCols int
	oErr     error
}

func (tester mNullDirectiveTester) assertMatrix() func(*testing.T) {
	return func(t *testing.T) {
		got, oErr := mNullDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols)
		assertZeroed(t, oErr)
		assertWildNullMatrix(t, got, tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols)
	}
}

func (tester mNullDirectiveTester) assertMinRowCountGreaterThanMaxRowCountError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := mNullDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols)
		assertZeroed(t, oMatrix)
		assertMinGreaterThanMaxError(t, got, tester.oErr)
	}
}

func (tester mNullDirectiveTester) assertMinColumnCountGreaterThanMaxColumnCountError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := mNullDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols)
		assertZeroed(t, oMatrix)
		assertMinGreaterThanMaxError(t, got, tester.oErr)
	}
}

func (tester mNullDirectiveTester) assertInvalidRowCountError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := mNullDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols)
		assertZeroed(t, oMatrix)
		assertInvalidRowCountError(t, got, tester.oErr)
	}
}

func (tester mNullDirectiveTester) assertInvalidColumnCountError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := mNullDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols)
		assertZeroed(t, oMatrix)
		assertInvalidColumnCountError(t, got, tester.oErr)
	}
}

const limit int = 10

func TestNullDirective(t *testing.T) {
	t.Run("baseline", nullDirectiveTester{}.assertNil())
}

func TestIntegerDirective(t *testing.T) {
	for i := 0; i < limit; i++ {
		t.Run(fmt.Sprintf("%d baseline", i), integerDirectiveTester{
			iMin: 0,
			iMax: 3,
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
	t.Run("empty charset", stringDirectiveTester{
		iMinLength: 3,
		iMaxLength: 6,
		iCharset:   "",
		oErr:       errors.New("freeformgen: charset cannot be empty"),
	}.assertEmptyCharsetError())
}

func TestBooleanDirective(t *testing.T) {
	t.Run("baseline", booleanDirectiveTester{}.assertBoolean())
}

func TestPrimitiveDirective(t *testing.T) {
	for i := 0; i < limit; i++ {
		t.Run(fmt.Sprintf("%d baseline", i), primitiveDirectiveTester{}.assertPrimitive())
	}
}

func TestVNullDirective(t *testing.T) {
	for i := 0; i < limit; i++ {
		t.Run(fmt.Sprintf("%d baseline", i), vNullDirectiveTester{
			iMinLength: 3,
			iMaxLength: 6,
		}.assertVector())
		t.Run(fmt.Sprintf("%d length of zero", i), vNullDirectiveTester{
			iMinLength: 0,
			iMaxLength: 0,
		}.assertVector())
		t.Run(fmt.Sprintf("%d equal lengths", i), vNullDirectiveTester{
			iMinLength: 6,
			iMaxLength: 6,
		}.assertVector())
	}
	t.Run("invalid min length", vNullDirectiveTester{
		iMinLength: -1,
		iMaxLength: 6,
		oErr:       errors.New("freeformgen: vector cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("invalid max length", vNullDirectiveTester{
		iMinLength: 3,
		iMaxLength: -1,
		oErr:       errors.New("freeformgen: vector cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("min length greater than max length", vNullDirectiveTester{
		iMinLength: 6,
		iMaxLength: 3,
		oErr:       errors.New("freeformgen: min length cannot exceed max length"),
	}.assertMinGreaterThanMaxError())
}

func TestVIntegerDirective(t *testing.T) {
	for i := 0; i < limit; i++ {
		t.Run(fmt.Sprintf("%d baseline", i), vIntegerDirectiveTester{
			iMin:       0,
			iMax:       3,
			iMinLength: 3,
			iMaxLength: 6,
		}.assertVector())
		t.Run(fmt.Sprintf("%d broad range", i), vIntegerDirectiveTester{
			iMin:       -10,
			iMax:       10,
			iMinLength: 3,
			iMaxLength: 6,
		}.assertVector())
		t.Run(fmt.Sprintf("%d equal", i), vIntegerDirectiveTester{
			iMinLength: 3,
			iMaxLength: 6,
		}.assertVector())
		t.Run(fmt.Sprintf("%d length of zero", i), vIntegerDirectiveTester{
			iMinLength: 0,
			iMaxLength: 0,
		}.assertVector())
		t.Run(fmt.Sprintf("%d equal lengths", i), vIntegerDirectiveTester{
			iMinLength: 6,
			iMaxLength: 6,
		}.assertVector())
	}
	t.Run("invalid min length", vIntegerDirectiveTester{
		iMinLength: -1,
		iMaxLength: 6,
		oErr:       errors.New("freeformgen: vector cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("invalid max length", vIntegerDirectiveTester{
		iMinLength: 3,
		iMaxLength: -1,
		oErr:       errors.New("freeformgen: vector cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("min greater than max", vIntegerDirectiveTester{
		iMin: 1,
		iMax: -1,
		oErr: errors.New("freeformgen: min cannot exceed max"),
	}.assertMinGreaterThanMaxError())
	t.Run("min length greater than max length", vIntegerDirectiveTester{
		iMinLength: 6,
		iMaxLength: 3,
		oErr:       errors.New("freeformgen: min length cannot exceed max length"),
	}.assertMinGreaterThanMaxError())
}

func TestVFloatDirective(t *testing.T) {
	for i := 0; i < limit; i++ {
		t.Run(fmt.Sprintf("%d baseline", i), vFloatDirectiveTester{
			iMin:       0.0,
			iMax:       1.0,
			iMinLength: 3,
			iMaxLength: 6,
		}.assertVector())
		t.Run(fmt.Sprintf("%d broad range", i), vFloatDirectiveTester{
			iMin:       -10.0,
			iMax:       10.0,
			iMinLength: 3,
			iMaxLength: 6,
		}.assertVector())
		t.Run(fmt.Sprintf("%d equal", i), vFloatDirectiveTester{
			iMinLength: 3,
			iMaxLength: 6,
		}.assertVector())
		t.Run(fmt.Sprintf("%d length of zero", i), vFloatDirectiveTester{
			iMinLength: 0,
			iMaxLength: 0,
		}.assertVector())
		t.Run(fmt.Sprintf("%d equal lengths", i), vFloatDirectiveTester{
			iMinLength: 6,
			iMaxLength: 6,
		}.assertVector())
	}
	t.Run("invalid min length", vFloatDirectiveTester{
		iMinLength: -1,
		iMaxLength: 6,
		oErr:       errors.New("freeformgen: vector cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("invalid max length", vFloatDirectiveTester{
		iMinLength: 3,
		iMaxLength: -1,
		oErr:       errors.New("freeformgen: vector cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("min greater than max", vFloatDirectiveTester{
		iMin: 1.0,
		iMax: -1.0,
		oErr: errors.New("freeformgen: min cannot exceed max"),
	}.assertMinGreaterThanMaxError())
	t.Run("min length greater than max length", vFloatDirectiveTester{
		iMinLength: 6,
		iMaxLength: 3,
		oErr:       errors.New("freeformgen: min length cannot exceed max length"),
	}.assertMinGreaterThanMaxError())
}

func TestVStringDirective(t *testing.T) {
	for i := 0; i < limit; i++ {
		t.Run(fmt.Sprintf("%d baseline", i), vStringDirectiveTester{
			iMinStrLength: 3,
			iMaxStrLength: 6,
			iMinLength:    3,
			iMaxLength:    6,
			iCharset:      "abc",
		}.assertVector())
		t.Run(fmt.Sprintf("%d emojis", i), vStringDirectiveTester{
			iMinStrLength: 3,
			iMaxStrLength: 6,
			iMinLength:    3,
			iMaxLength:    6,
			iCharset:      "🔴🟡🟢",
		}.assertVector())
		t.Run(fmt.Sprintf("%d length of zero", i), vStringDirectiveTester{
			iMinStrLength: 3,
			iMaxStrLength: 6,
			iMinLength:    0,
			iMaxLength:    0,
			iCharset:      "abc",
		}.assertVector())
		t.Run(fmt.Sprintf("%d equal lengths", i), vStringDirectiveTester{
			iMinStrLength: 3,
			iMaxStrLength: 6,
			iMinLength:    6,
			iMaxLength:    6,
			iCharset:      "abc",
		}.assertVector())
		t.Run(fmt.Sprintf("%d string length of zero", i), vStringDirectiveTester{
			iMinStrLength: 0,
			iMaxStrLength: 0,
			iMinLength:    3,
			iMaxLength:    6,
			iCharset:      "abc",
		}.assertVector())
		t.Run(fmt.Sprintf("%d equal string lengths", i), vStringDirectiveTester{
			iMinStrLength: 6,
			iMaxStrLength: 6,
			iMinLength:    3,
			iMaxLength:    6,
			iCharset:      "abc",
		}.assertVector())
	}
	t.Run("invalid min length", vStringDirectiveTester{
		iMinLength: -1,
		iMaxLength: 6,
		oErr:       errors.New("freeformgen: vector cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("invalid max length", vStringDirectiveTester{
		iMinLength: 3,
		iMaxLength: -1,
		oErr:       errors.New("freeformgen: vector cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("invalid min string length", vStringDirectiveTester{
		iMinStrLength: -1,
		iMaxStrLength: 6,
		oErr:          errors.New("freeformgen: string cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("invalid max string length", vStringDirectiveTester{
		iMinStrLength: 3,
		iMaxStrLength: -1,
		oErr:          errors.New("freeformgen: string cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("min string length greater than max string length", vStringDirectiveTester{
		iMinStrLength: 6,
		iMaxStrLength: 3,
		oErr:          errors.New("freeformgen: min string length cannot exceed max string length"),
	}.assertMinGreaterThanMaxError())
	t.Run("min length greater than max length", vStringDirectiveTester{
		iMinLength: 6,
		iMaxLength: 3,
		oErr:       errors.New("freeformgen: min length cannot exceed max length"),
	}.assertMinGreaterThanMaxError())
	t.Run("empty charset", vStringDirectiveTester{
		iMinStrLength: 3,
		iMaxStrLength: 6,
		iMinLength:    3,
		iMaxLength:    6,
		iCharset:      "",
		oErr:          errors.New("freeformgen: charset cannot be empty"),
	}.assertEmptyCharsetError())
}

func TestVBooleanDirective(t *testing.T) {
	for i := 0; i < limit; i++ {
		t.Run(fmt.Sprintf("%d baseline", i), vBooleanDirectiveTester{
			iMinLength: 3,
			iMaxLength: 6,
		}.assertVector())
		t.Run(fmt.Sprintf("%d length of zero", i), vBooleanDirectiveTester{
			iMinLength: 0,
			iMaxLength: 0,
		}.assertVector())
		t.Run(fmt.Sprintf("%d equal lengths", i), vBooleanDirectiveTester{
			iMinLength: 6,
			iMaxLength: 6,
		}.assertVector())
	}
	t.Run("invalid min length", vBooleanDirectiveTester{
		iMinLength: -1,
		iMaxLength: 6,
		oErr:       errors.New("freeformgen: vector cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("invalid max length", vBooleanDirectiveTester{
		iMinLength: 3,
		iMaxLength: -1,
		oErr:       errors.New("freeformgen: vector cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("min length greater than max length", vBooleanDirectiveTester{
		iMinLength: 6,
		iMaxLength: 3,
		oErr:       errors.New("freeformgen: min length cannot exceed max length"),
	}.assertMinGreaterThanMaxError())
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

func TestMNullDirective(t *testing.T) {
	for i := 0; i < limit; i++ {
		t.Run(fmt.Sprintf("%d baseline", i), mNullDirectiveTester{
			iMinRows: 3,
			iMaxRows: 6,
			iMinCols: 3,
			iMaxCols: 6,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d row count of zero", i), mNullDirectiveTester{
			iMinRows: 0,
			iMaxRows: 0,
			iMinCols: 3,
			iMaxCols: 6,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d column count of zero", i), mNullDirectiveTester{
			iMinRows: 3,
			iMaxRows: 6,
			iMinCols: 0,
			iMaxCols: 0,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d square matrix", i), mNullDirectiveTester{
			iMinRows: 6,
			iMaxRows: 6,
			iMinCols: 6,
			iMaxCols: 6,
		}.assertMatrix())
	}
	t.Run("invalid min row count", mNullDirectiveTester{
		iMinRows: -1,
		iMaxRows: 6,
		iMinCols: 3,
		iMaxCols: 6,
		oErr:     errors.New("freeformgen: matrix cannot have a negative row count"),
	}.assertInvalidRowCountError())
	t.Run("invalid max row count", mNullDirectiveTester{
		iMinRows: 3,
		iMaxRows: -1,
		iMinCols: 3,
		iMaxCols: 6,
		oErr:     errors.New("freeformgen: matrix cannot have a negative row count"),
	}.assertInvalidRowCountError())
	t.Run("invalid min column count", mNullDirectiveTester{
		iMinRows: 3,
		iMaxRows: 6,
		iMinCols: -1,
		iMaxCols: 6,
		oErr:     errors.New("freeformgen: matrix cannot have a negative column count"),
	}.assertInvalidColumnCountError())
	t.Run("invalid max column count", mNullDirectiveTester{
		iMinRows: 3,
		iMaxRows: 6,
		iMinCols: 3,
		iMaxCols: -1,
		oErr:     errors.New("freeformgen: matrix cannot have a negative column count"),
	}.assertInvalidColumnCountError())
	t.Run("min row count greater than max row count", mNullDirectiveTester{
		iMinRows: 6,
		iMaxRows: 3,
		iMinCols: 3,
		iMaxCols: 6,
		oErr:     errors.New("freeformgen: min row count cannot exceed max row count"),
	}.assertMinRowCountGreaterThanMaxRowCountError())
	t.Run("min column count greater than max column count", mNullDirectiveTester{
		iMinRows: 3,
		iMaxRows: 6,
		iMinCols: 6,
		iMaxCols: 3,
		oErr:     errors.New("freeformgen: min column count cannot exceed max column count"),
	}.assertMinColumnCountGreaterThanMaxColumnCountError())
}
