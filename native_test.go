package main

import (
	"errors"
	"fmt"
	"testing"
)

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

func (tester mNullDirectiveTester) assertMinGreaterThanMaxError() func(*testing.T) {
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

func (tester mNullDirectiveTester) assertInvalidColCountError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := mNullDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols)
		assertZeroed(t, oMatrix)
		assertInvalidColCountError(t, got, tester.oErr)
	}
}

type mIntegerDirectiveTester struct {
	iMin     int
	iMax     int
	iMinRows int
	iMaxRows int
	iMinCols int
	iMaxCols int
	oErr     error
}

func (tester mIntegerDirectiveTester) assertMatrix() func(*testing.T) {
	return func(t *testing.T) {
		got, oErr := mIntegerDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols, tester.iMin, tester.iMax)
		assertZeroed(t, oErr)
		assertWildNumberMatrix(t, got, tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols, tester.iMin, tester.iMax)
	}
}

func (tester mIntegerDirectiveTester) assertMinGreaterThanMaxError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := mIntegerDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols, tester.iMin, tester.iMax)
		assertZeroed(t, oMatrix)
		assertMinGreaterThanMaxError(t, got, tester.oErr)
	}
}

func (tester mIntegerDirectiveTester) assertInvalidRowCountError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := mIntegerDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols, tester.iMin, tester.iMax)
		assertZeroed(t, oMatrix)
		assertInvalidRowCountError(t, got, tester.oErr)
	}
}

func (tester mIntegerDirectiveTester) assertInvalidColCountError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := mIntegerDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols, tester.iMin, tester.iMax)
		assertZeroed(t, oMatrix)
		assertInvalidColCountError(t, got, tester.oErr)
	}
}

type mFloatDirectiveTester struct {
	iMin     float64
	iMax     float64
	iMinRows int
	iMaxRows int
	iMinCols int
	iMaxCols int
	oErr     error
}

func (tester mFloatDirectiveTester) assertMatrix() func(*testing.T) {
	return func(t *testing.T) {
		got, oErr := mFloatDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols, tester.iMin, tester.iMax)
		assertZeroed(t, oErr)
		assertWildNumberMatrix(t, got, tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols, tester.iMin, tester.iMax)
	}
}

func (tester mFloatDirectiveTester) assertMinGreaterThanMaxError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := mFloatDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols, tester.iMin, tester.iMax)
		assertZeroed(t, oMatrix)
		assertMinGreaterThanMaxError(t, got, tester.oErr)
	}
}

func (tester mFloatDirectiveTester) assertInvalidRowCountError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := mFloatDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols, tester.iMin, tester.iMax)
		assertZeroed(t, oMatrix)
		assertInvalidRowCountError(t, got, tester.oErr)
	}
}

func (tester mFloatDirectiveTester) assertInvalidColCountError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := mFloatDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols, tester.iMin, tester.iMax)
		assertZeroed(t, oMatrix)
		assertInvalidColCountError(t, got, tester.oErr)
	}
}

type mStringDirectiveTester struct {
	iMinStrLength int
	iMaxStrLength int
	iCharset      string
	iMinRows      int
	iMaxRows      int
	iMinCols      int
	iMaxCols      int
	oErr          error
}

func (tester mStringDirectiveTester) assertMatrix() func(*testing.T) {
	return func(t *testing.T) {
		got, oErr := mStringDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols, tester.iMinStrLength, tester.iMaxStrLength, tester.iCharset)
		assertZeroed(t, oErr)
		assertWildStringMatrix(t, got, tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols, tester.iMinStrLength, tester.iMaxStrLength)
	}
}

func (tester mStringDirectiveTester) assertMinGreaterThanMaxError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := mStringDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols, tester.iMinStrLength, tester.iMaxStrLength, tester.iCharset)
		assertZeroed(t, oMatrix)
		assertMinGreaterThanMaxError(t, got, tester.oErr)
	}
}

func (tester mStringDirectiveTester) assertInvalidRowCountError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := mStringDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols, tester.iMinStrLength, tester.iMaxStrLength, tester.iCharset)
		assertZeroed(t, oMatrix)
		assertInvalidRowCountError(t, got, tester.oErr)
	}
}

func (tester mStringDirectiveTester) assertInvalidColCountError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := mStringDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols, tester.iMinStrLength, tester.iMaxStrLength, tester.iCharset)
		assertZeroed(t, oMatrix)
		assertInvalidColCountError(t, got, tester.oErr)
	}
}

func (tester mStringDirectiveTester) assertInvalidStrLengthError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := mStringDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols, tester.iMinStrLength, tester.iMaxStrLength, tester.iCharset)
		assertZeroed(t, oMatrix)
		assertInvalidLengthError(t, got, tester.oErr)
	}
}

func (tester mStringDirectiveTester) assertEmptyCharsetError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := mStringDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols, tester.iMinStrLength, tester.iMaxStrLength, tester.iCharset)
		assertZeroed(t, oMatrix)
		assertEmptyCharsetError(t, got, tester.oErr)
	}
}

type mBooleanDirectiveTester struct {
	iMinRows int
	iMaxRows int
	iMinCols int
	iMaxCols int
	oErr     error
}

func (tester mBooleanDirectiveTester) assertMatrix() func(*testing.T) {
	return func(t *testing.T) {
		got, oErr := mBooleanDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols)
		assertZeroed(t, oErr)
		assertWildBooleanMatrix(t, got, tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols)
	}
}

func (tester mBooleanDirectiveTester) assertMinGreaterThanMaxError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := mBooleanDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols)
		assertZeroed(t, oMatrix)
		assertMinGreaterThanMaxError(t, got, tester.oErr)
	}
}

func (tester mBooleanDirectiveTester) assertInvalidRowCountError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := mBooleanDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols)
		assertZeroed(t, oMatrix)
		assertInvalidRowCountError(t, got, tester.oErr)
	}
}

func (tester mBooleanDirectiveTester) assertInvalidColCountError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := mBooleanDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols)
		assertZeroed(t, oMatrix)
		assertInvalidColCountError(t, got, tester.oErr)
	}
}

type mPrimitiveDirectiveTester struct {
	iMinRows int
	iMaxRows int
	iMinCols int
	iMaxCols int
	oErr     error
}

func (tester mPrimitiveDirectiveTester) assertMatrix() func(*testing.T) {
	return func(t *testing.T) {
		got, oErr := mPrimitiveDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols)
		assertZeroed(t, oErr)
		assertWildPrimitiveMatrix(t, got, tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols)
	}
}

func (tester mPrimitiveDirectiveTester) assertMinGreaterThanMaxError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := mPrimitiveDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols)
		assertZeroed(t, oMatrix)
		assertMinGreaterThanMaxError(t, got, tester.oErr)
	}
}

func (tester mPrimitiveDirectiveTester) assertInvalidRowCountError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := mPrimitiveDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols)
		assertZeroed(t, oMatrix)
		assertInvalidRowCountError(t, got, tester.oErr)
	}
}

func (tester mPrimitiveDirectiveTester) assertInvalidColCountError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := mPrimitiveDirective(tester.iMinRows, tester.iMaxRows, tester.iMinCols, tester.iMaxCols)
		assertZeroed(t, oMatrix)
		assertInvalidColCountError(t, got, tester.oErr)
	}
}

type vectorOfDirectiveTester struct {
	iTyp       string
	iMinLength int
	iMaxLength int
	iArgs      []any
	oErr       error
}

func (tester vectorOfDirectiveTester) assertVector() func(*testing.T) {
	return func(t *testing.T) {
		got, oErr := vectorOfDirective(tester.iTyp, tester.iMinLength, tester.iMaxLength, tester.iArgs...)
		assertZeroed(t, oErr)
		assertWildPrimitiveVector(t, got, tester.iMinLength, tester.iMaxLength)
	}
}

func (tester vectorOfDirectiveTester) assertInvalidLengthError() func(*testing.T) {
	return func(t *testing.T) {
		oVector, got := vectorOfDirective(tester.iTyp, tester.iMinLength, tester.iMaxLength, tester.iArgs...)
		assertZeroed(t, oVector)
		assertInvalidLengthError(t, got, tester.oErr)
	}
}

func (tester vectorOfDirectiveTester) assertMinGreaterThanMaxError() func(*testing.T) {
	return func(t *testing.T) {
		oVector, got := vectorOfDirective(tester.iTyp, tester.iMinLength, tester.iMaxLength, tester.iArgs...)
		assertZeroed(t, oVector)
		assertMinGreaterThanMaxError(t, got, tester.oErr)
	}
}

func (tester vectorOfDirectiveTester) assertIncorrectArgsError() func(*testing.T) {
	return func(t *testing.T) {
		oVector, got := vectorOfDirective(tester.iTyp, tester.iMinLength, tester.iMaxLength, tester.iArgs...)
		assertZeroed(t, oVector)
		assertIncorrectArgsError(t, got, tester.oErr)
	}
}

func (tester vectorOfDirectiveTester) assertEmptyCharsetError() func(*testing.T) {
	return func(t *testing.T) {
		oVector, got := vectorOfDirective(tester.iTyp, tester.iMinLength, tester.iMaxLength, tester.iArgs...)
		assertZeroed(t, oVector)
		assertEmptyCharsetError(t, got, tester.oErr)
	}
}

const limit int = 10

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
	}.assertInvalidColCountError())
	t.Run("invalid max column count", mNullDirectiveTester{
		iMinRows: 3,
		iMaxRows: 6,
		iMinCols: 3,
		iMaxCols: -1,
		oErr:     errors.New("freeformgen: matrix cannot have a negative column count"),
	}.assertInvalidColCountError())
	t.Run("min row count greater than max row count", mNullDirectiveTester{
		iMinRows: 6,
		iMaxRows: 3,
		iMinCols: 3,
		iMaxCols: 6,
		oErr:     errors.New("freeformgen: min row count cannot exceed max row count"),
	}.assertMinGreaterThanMaxError())
	t.Run("min column count greater than max column count", mNullDirectiveTester{
		iMinRows: 3,
		iMaxRows: 6,
		iMinCols: 6,
		iMaxCols: 3,
		oErr:     errors.New("freeformgen: min column count cannot exceed max column count"),
	}.assertMinGreaterThanMaxError())
}

func TestMIntegerDirective(t *testing.T) {
	for i := 0; i < limit; i++ {
		t.Run(fmt.Sprintf("%d baseline", i), mIntegerDirectiveTester{
			iMin:     0,
			iMax:     3,
			iMinRows: 3,
			iMaxRows: 6,
			iMinCols: 3,
			iMaxCols: 6,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d row count of zero", i), mIntegerDirectiveTester{
			iMin:     0,
			iMax:     3,
			iMinRows: 0,
			iMaxRows: 0,
			iMinCols: 3,
			iMaxCols: 6,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d column count of zero", i), mIntegerDirectiveTester{
			iMin:     0,
			iMax:     3,
			iMinRows: 3,
			iMaxRows: 6,
			iMinCols: 0,
			iMaxCols: 0,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d square matrix", i), mIntegerDirectiveTester{
			iMin:     0,
			iMax:     3,
			iMinRows: 6,
			iMaxRows: 6,
			iMinCols: 6,
			iMaxCols: 6,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d broad range", i), mIntegerDirectiveTester{
			iMin:     -10,
			iMax:     10,
			iMinRows: 3,
			iMaxRows: 6,
			iMinCols: 3,
			iMaxCols: 6,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d equal", i), mIntegerDirectiveTester{
			iMinRows: 3,
			iMaxRows: 6,
			iMinCols: 3,
			iMaxCols: 6,
		}.assertMatrix())
	}
	t.Run("invalid min row count", mIntegerDirectiveTester{
		iMin:     0,
		iMax:     3,
		iMinRows: -1,
		iMaxRows: 6,
		iMinCols: 3,
		iMaxCols: 6,
		oErr:     errors.New("freeformgen: matrix cannot have a negative row count"),
	}.assertInvalidRowCountError())
	t.Run("invalid max row count", mIntegerDirectiveTester{
		iMin:     0,
		iMax:     3,
		iMinRows: 3,
		iMaxRows: -1,
		iMinCols: 3,
		iMaxCols: 6,
		oErr:     errors.New("freeformgen: matrix cannot have a negative row count"),
	}.assertInvalidRowCountError())
	t.Run("invalid min column count", mIntegerDirectiveTester{
		iMin:     0,
		iMax:     3,
		iMinRows: 3,
		iMaxRows: 6,
		iMinCols: -1,
		iMaxCols: 6,
		oErr:     errors.New("freeformgen: matrix cannot have a negative column count"),
	}.assertInvalidColCountError())
	t.Run("invalid max column count", mIntegerDirectiveTester{
		iMin:     0,
		iMax:     3,
		iMinRows: 3,
		iMaxRows: 6,
		iMinCols: 3,
		iMaxCols: -1,
		oErr:     errors.New("freeformgen: matrix cannot have a negative column count"),
	}.assertInvalidColCountError())
	t.Run("min row count greater than max row count", mIntegerDirectiveTester{
		iMin:     0,
		iMax:     3,
		iMinRows: 6,
		iMaxRows: 3,
		iMinCols: 3,
		iMaxCols: 6,
		oErr:     errors.New("freeformgen: min row count cannot exceed max row count"),
	}.assertMinGreaterThanMaxError())
	t.Run("min column count greater than max column count", mIntegerDirectiveTester{
		iMin:     0,
		iMax:     3,
		iMinRows: 3,
		iMaxRows: 6,
		iMinCols: 6,
		iMaxCols: 3,
		oErr:     errors.New("freeformgen: min column count cannot exceed max column count"),
	}.assertMinGreaterThanMaxError())
	t.Run("min greater than max", mIntegerDirectiveTester{
		iMin: 1,
		iMax: -1,
		oErr: errors.New("freeformgen: min cannot exceed max"),
	}.assertMinGreaterThanMaxError())
}

func TestMFloatDirective(t *testing.T) {
	for i := 0; i < limit; i++ {
		t.Run(fmt.Sprintf("%d baseline", i), mFloatDirectiveTester{
			iMin:     0.0,
			iMax:     1.0,
			iMinRows: 3,
			iMaxRows: 6,
			iMinCols: 3,
			iMaxCols: 6,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d row count of zero", i), mFloatDirectiveTester{
			iMin:     0.0,
			iMax:     1.0,
			iMinRows: 0,
			iMaxRows: 0,
			iMinCols: 3,
			iMaxCols: 6,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d column count of zero", i), mFloatDirectiveTester{
			iMin:     0.0,
			iMax:     1.0,
			iMinRows: 3,
			iMaxRows: 6,
			iMinCols: 0,
			iMaxCols: 0,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d square matrix", i), mFloatDirectiveTester{
			iMin:     0.0,
			iMax:     1.0,
			iMinRows: 6,
			iMaxRows: 6,
			iMinCols: 6,
			iMaxCols: 6,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d broad range", i), mFloatDirectiveTester{
			iMin:     -10.0,
			iMax:     10.0,
			iMinRows: 3,
			iMaxRows: 6,
			iMinCols: 3,
			iMaxCols: 6,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d equal", i), mFloatDirectiveTester{
			iMinRows: 3,
			iMaxRows: 6,
			iMinCols: 3,
			iMaxCols: 6,
		}.assertMatrix())
	}
	t.Run("invalid min row count", mFloatDirectiveTester{
		iMin:     0.0,
		iMax:     1.0,
		iMinRows: -1,
		iMaxRows: 6,
		iMinCols: 3,
		iMaxCols: 6,
		oErr:     errors.New("freeformgen: matrix cannot have a negative row count"),
	}.assertInvalidRowCountError())
	t.Run("invalid max row count", mFloatDirectiveTester{
		iMin:     0.0,
		iMax:     1.0,
		iMinRows: 3,
		iMaxRows: -1,
		iMinCols: 3,
		iMaxCols: 6,
		oErr:     errors.New("freeformgen: matrix cannot have a negative row count"),
	}.assertInvalidRowCountError())
	t.Run("invalid min column count", mFloatDirectiveTester{
		iMin:     0.0,
		iMax:     1.0,
		iMinRows: 3,
		iMaxRows: 6,
		iMinCols: -1,
		iMaxCols: 6,
		oErr:     errors.New("freeformgen: matrix cannot have a negative column count"),
	}.assertInvalidColCountError())
	t.Run("invalid max column count", mFloatDirectiveTester{
		iMin:     0.0,
		iMax:     1.0,
		iMinRows: 3,
		iMaxRows: 6,
		iMinCols: 3,
		iMaxCols: -1,
		oErr:     errors.New("freeformgen: matrix cannot have a negative column count"),
	}.assertInvalidColCountError())
	t.Run("min row count greater than max row count", mFloatDirectiveTester{
		iMin:     0.0,
		iMax:     1.0,
		iMinRows: 6,
		iMaxRows: 3,
		iMinCols: 3,
		iMaxCols: 6,
		oErr:     errors.New("freeformgen: min row count cannot exceed max row count"),
	}.assertMinGreaterThanMaxError())
	t.Run("min column count greater than max column count", mFloatDirectiveTester{
		iMin:     0.0,
		iMax:     1.0,
		iMinRows: 3,
		iMaxRows: 6,
		iMinCols: 6,
		iMaxCols: 3,
		oErr:     errors.New("freeformgen: min column count cannot exceed max column count"),
	}.assertMinGreaterThanMaxError())
	t.Run("min greater than max", mFloatDirectiveTester{
		iMin: 10.0,
		iMax: -10.0,
		oErr: errors.New("freeformgen: min cannot exceed max"),
	}.assertMinGreaterThanMaxError())
}

func TestMStringDirective(t *testing.T) {
	for i := 0; i < limit; i++ {
		t.Run(fmt.Sprintf("%d baseline", i), mStringDirectiveTester{
			iMinStrLength: 3,
			iMaxStrLength: 6,
			iCharset:      "abc",
			iMinRows:      3,
			iMaxRows:      6,
			iMinCols:      3,
			iMaxCols:      6,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d row count of zero", i), mStringDirectiveTester{
			iMinStrLength: 3,
			iMaxStrLength: 6,
			iCharset:      "abc",
			iMinRows:      0,
			iMaxRows:      0,
			iMinCols:      3,
			iMaxCols:      6,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d column count of zero", i), mStringDirectiveTester{
			iMinStrLength: 3,
			iMaxStrLength: 6,
			iCharset:      "abc",
			iMinRows:      3,
			iMaxRows:      6,
			iMinCols:      0,
			iMaxCols:      0,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d square matrix", i), mStringDirectiveTester{
			iMinStrLength: 3,
			iMaxStrLength: 6,
			iCharset:      "abc",
			iMinRows:      6,
			iMaxRows:      6,
			iMinCols:      6,
			iMaxCols:      6,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d emojis", i), mStringDirectiveTester{
			iMinStrLength: 3,
			iMaxStrLength: 6,
			iCharset:      "🔴🟡🟢",
			iMinRows:      3,
			iMaxRows:      6,
			iMinCols:      3,
			iMaxCols:      6,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d string length of zero", i), mStringDirectiveTester{
			iMinStrLength: 0,
			iMaxStrLength: 0,
			iCharset:      "abc",
			iMinRows:      3,
			iMaxRows:      6,
			iMinCols:      3,
			iMaxCols:      6,
		}.assertMatrix())
	}
	t.Run("invalid min row count", mStringDirectiveTester{
		iMinStrLength: 3,
		iMaxStrLength: 6,
		iCharset:      "abc",
		iMinRows:      -1,
		iMaxRows:      6,
		iMinCols:      3,
		iMaxCols:      6,
		oErr:          errors.New("freeformgen: matrix cannot have a negative row count"),
	}.assertInvalidRowCountError())
	t.Run("invalid max row count", mStringDirectiveTester{
		iMinStrLength: 3,
		iMaxStrLength: 6,
		iCharset:      "abc",
		iMinRows:      3,
		iMaxRows:      -1,
		iMinCols:      3,
		iMaxCols:      6,
		oErr:          errors.New("freeformgen: matrix cannot have a negative row count"),
	}.assertInvalidRowCountError())
	t.Run("invalid min column count", mStringDirectiveTester{
		iMinStrLength: 3,
		iMaxStrLength: 6,
		iCharset:      "abc",
		iMinRows:      3,
		iMaxRows:      6,
		iMinCols:      -1,
		iMaxCols:      6,
		oErr:          errors.New("freeformgen: matrix cannot have a negative column count"),
	}.assertInvalidColCountError())
	t.Run("invalid max column count", mStringDirectiveTester{
		iMinStrLength: 3,
		iMaxStrLength: 6,
		iCharset:      "abc",
		iMinRows:      3,
		iMaxRows:      6,
		iMinCols:      3,
		iMaxCols:      -1,
		oErr:          errors.New("freeformgen: matrix cannot have a negative column count"),
	}.assertInvalidColCountError())
	t.Run("min row count greater than max row count", mStringDirectiveTester{
		iMinStrLength: 3,
		iMaxStrLength: 6,
		iCharset:      "abc",
		iMinRows:      6,
		iMaxRows:      3,
		iMinCols:      3,
		iMaxCols:      6,
		oErr:          errors.New("freeformgen: min row count cannot exceed max row count"),
	}.assertMinGreaterThanMaxError())
	t.Run("min column count greater than max column count", mStringDirectiveTester{
		iMinStrLength: 3,
		iMaxStrLength: 6,
		iCharset:      "abc",
		iMinRows:      3,
		iMaxRows:      6,
		iMinCols:      6,
		iMaxCols:      3,
		oErr:          errors.New("freeformgen: min column count cannot exceed max column count"),
	}.assertMinGreaterThanMaxError())
	t.Run("invalid min string length", mStringDirectiveTester{
		iMinStrLength: -1,
		iMaxStrLength: 6,
		iCharset:      "abc",
		iMinRows:      3,
		iMaxRows:      6,
		iMinCols:      3,
		iMaxCols:      6,
		oErr:          errors.New("freeformgen: string cannot have a negative length"),
	}.assertInvalidStrLengthError())
	t.Run("invalid max string length", mStringDirectiveTester{
		iMinStrLength: 3,
		iMaxStrLength: -1,
		iCharset:      "abc",
		iMinRows:      3,
		iMaxRows:      6,
		iMinCols:      3,
		iMaxCols:      6,
		oErr:          errors.New("freeformgen: string cannot have a negative length"),
	}.assertInvalidStrLengthError())
	t.Run("min string length greater than max string length", mStringDirectiveTester{
		iMinStrLength: 6,
		iMaxStrLength: 3,
		iCharset:      "abc",
		iMinRows:      3,
		iMaxRows:      6,
		iMinCols:      3,
		iMaxCols:      6,
		oErr:          errors.New("freeformgen: min string length cannot exceed max string length"),
	}.assertMinGreaterThanMaxError())
	t.Run("empty charset", mStringDirectiveTester{
		iMinStrLength: 3,
		iMaxStrLength: 6,
		iCharset:      "",
		iMinRows:      3,
		iMaxRows:      6,
		iMinCols:      3,
		iMaxCols:      6,
		oErr:          errors.New("freeformgen: charset cannot be empty"),
	}.assertEmptyCharsetError())
}

func TestMBooleanDirective(t *testing.T) {
	for i := 0; i < limit; i++ {
		t.Run(fmt.Sprintf("%d baseline", i), mBooleanDirectiveTester{
			iMinRows: 3,
			iMaxRows: 6,
			iMinCols: 3,
			iMaxCols: 6,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d row count of zero", i), mBooleanDirectiveTester{
			iMinRows: 0,
			iMaxRows: 0,
			iMinCols: 3,
			iMaxCols: 6,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d column count of zero", i), mBooleanDirectiveTester{
			iMinRows: 3,
			iMaxRows: 6,
			iMinCols: 0,
			iMaxCols: 0,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d square matrix", i), mBooleanDirectiveTester{
			iMinRows: 6,
			iMaxRows: 6,
			iMinCols: 6,
			iMaxCols: 6,
		}.assertMatrix())
	}
	t.Run("invalid min row count", mBooleanDirectiveTester{
		iMinRows: -1,
		iMaxRows: 6,
		iMinCols: 3,
		iMaxCols: 6,
		oErr:     errors.New("freeformgen: matrix cannot have a negative row count"),
	}.assertInvalidRowCountError())
	t.Run("invalid max row count", mBooleanDirectiveTester{
		iMinRows: 3,
		iMaxRows: -1,
		iMinCols: 3,
		iMaxCols: 6,
		oErr:     errors.New("freeformgen: matrix cannot have a negative row count"),
	}.assertInvalidRowCountError())
	t.Run("invalid min column count", mBooleanDirectiveTester{
		iMinRows: 3,
		iMaxRows: 6,
		iMinCols: -1,
		iMaxCols: 6,
		oErr:     errors.New("freeformgen: matrix cannot have a negative column count"),
	}.assertInvalidColCountError())
	t.Run("invalid max column count", mBooleanDirectiveTester{
		iMinRows: 3,
		iMaxRows: 6,
		iMinCols: 3,
		iMaxCols: -1,
		oErr:     errors.New("freeformgen: matrix cannot have a negative column count"),
	}.assertInvalidColCountError())
	t.Run("min row count greater than max row count", mBooleanDirectiveTester{
		iMinRows: 6,
		iMaxRows: 3,
		iMinCols: 3,
		iMaxCols: 6,
		oErr:     errors.New("freeformgen: min row count cannot exceed max row count"),
	}.assertMinGreaterThanMaxError())
	t.Run("min column count greater than max column count", mBooleanDirectiveTester{
		iMinRows: 3,
		iMaxRows: 6,
		iMinCols: 6,
		iMaxCols: 3,
		oErr:     errors.New("freeformgen: min column count cannot exceed max column count"),
	}.assertMinGreaterThanMaxError())
}

func TestMPrimitiveDirective(t *testing.T) {
	for i := 0; i < limit; i++ {
		t.Run(fmt.Sprintf("%d baseline", i), mPrimitiveDirectiveTester{
			iMinRows: 3,
			iMaxRows: 6,
			iMinCols: 3,
			iMaxCols: 6,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d row count of zero", i), mPrimitiveDirectiveTester{
			iMinRows: 0,
			iMaxRows: 0,
			iMinCols: 3,
			iMaxCols: 6,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d column count of zero", i), mPrimitiveDirectiveTester{
			iMinRows: 3,
			iMaxRows: 6,
			iMinCols: 0,
			iMaxCols: 0,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d square matrix", i), mPrimitiveDirectiveTester{
			iMinRows: 6,
			iMaxRows: 6,
			iMinCols: 6,
			iMaxCols: 6,
		}.assertMatrix())
	}
	t.Run("invalid min row count", mPrimitiveDirectiveTester{
		iMinRows: -1,
		iMaxRows: 6,
		iMinCols: 3,
		iMaxCols: 6,
		oErr:     errors.New("freeformgen: matrix cannot have a negative row count"),
	}.assertInvalidRowCountError())
	t.Run("invalid max row count", mPrimitiveDirectiveTester{
		iMinRows: 3,
		iMaxRows: -1,
		iMinCols: 3,
		iMaxCols: 6,
		oErr:     errors.New("freeformgen: matrix cannot have a negative row count"),
	}.assertInvalidRowCountError())
	t.Run("invalid min column count", mPrimitiveDirectiveTester{
		iMinRows: 3,
		iMaxRows: 6,
		iMinCols: -1,
		iMaxCols: 6,
		oErr:     errors.New("freeformgen: matrix cannot have a negative column count"),
	}.assertInvalidColCountError())
	t.Run("invalid max column count", mPrimitiveDirectiveTester{
		iMinRows: 3,
		iMaxRows: 6,
		iMinCols: 3,
		iMaxCols: -1,
		oErr:     errors.New("freeformgen: matrix cannot have a negative column count"),
	}.assertInvalidColCountError())
	t.Run("min row count greater than max row count", mPrimitiveDirectiveTester{
		iMinRows: 6,
		iMaxRows: 3,
		iMinCols: 3,
		iMaxCols: 6,
		oErr:     errors.New("freeformgen: min row count cannot exceed max row count"),
	}.assertMinGreaterThanMaxError())
	t.Run("min column count greater than max column count", mPrimitiveDirectiveTester{
		iMinRows: 3,
		iMaxRows: 6,
		iMinCols: 6,
		iMaxCols: 3,
		oErr:     errors.New("freeformgen: min column count cannot exceed max column count"),
	}.assertMinGreaterThanMaxError())
}

func TestVectorOfDirective(t *testing.T) {
	for i := 0; i < limit; i++ {
		t.Run(fmt.Sprintf("%d null baseline", i), vectorOfDirectiveTester{
			iTyp:       "null",
			iMinLength: 3,
			iMaxLength: 6,
		}.assertVector())
		t.Run(fmt.Sprintf("%d null length of zero", i), vectorOfDirectiveTester{
			iTyp:       "null",
			iMinLength: 0,
			iMaxLength: 0,
		}.assertVector())
		t.Run(fmt.Sprintf("%d null equal lengths", i), vectorOfDirectiveTester{
			iTyp:       "null",
			iMinLength: 6,
			iMaxLength: 6,
		}.assertVector())
		t.Run(fmt.Sprintf("%d int baseline", i), vectorOfDirectiveTester{
			iTyp:       "int",
			iMinLength: 3,
			iMaxLength: 6,
			iArgs:      []any{0, 3},
		}.assertVector())
		t.Run(fmt.Sprintf("%d int length of zero", i), vectorOfDirectiveTester{
			iTyp:       "int",
			iMinLength: 0,
			iMaxLength: 0,
			iArgs:      []any{0, 3},
		}.assertVector())
		t.Run(fmt.Sprintf("%d int equal lengths", i), vectorOfDirectiveTester{
			iTyp:       "int",
			iMinLength: 6,
			iMaxLength: 6,
			iArgs:      []any{0, 3},
		}.assertVector())
		t.Run(fmt.Sprintf("%d int broad range", i), vectorOfDirectiveTester{
			iTyp:       "int",
			iMinLength: 3,
			iMaxLength: 6,
			iArgs:      []any{-10, 10},
		}.assertVector())
		t.Run(fmt.Sprintf("%d int equal", i), vectorOfDirectiveTester{
			iTyp:       "int",
			iMinLength: 3,
			iMaxLength: 6,
			iArgs:      []any{0, 0},
		}.assertVector())
		t.Run(fmt.Sprintf("%d float baseline", i), vectorOfDirectiveTester{
			iTyp:       "float",
			iMinLength: 3,
			iMaxLength: 6,
			iArgs:      []any{0.0, 1.0},
		}.assertVector())
		t.Run(fmt.Sprintf("%d float length of zero", i), vectorOfDirectiveTester{
			iTyp:       "float",
			iMinLength: 0,
			iMaxLength: 0,
			iArgs:      []any{0.0, 1.0},
		}.assertVector())
		t.Run(fmt.Sprintf("%d float equal lengths", i), vectorOfDirectiveTester{
			iTyp:       "float",
			iMinLength: 6,
			iMaxLength: 6,
			iArgs:      []any{0.0, 1.0},
		}.assertVector())
		t.Run(fmt.Sprintf("%d float broad range", i), vectorOfDirectiveTester{
			iTyp:       "float",
			iMinLength: 3,
			iMaxLength: 6,
			iArgs:      []any{-10.0, 10.0},
		}.assertVector())
		t.Run(fmt.Sprintf("%d float equal", i), vectorOfDirectiveTester{
			iTyp:       "float",
			iMinLength: 3,
			iMaxLength: 6,
			iArgs:      []any{0.0, 0.0},
		}.assertVector())
		t.Run(fmt.Sprintf("%d string baseline", i), vectorOfDirectiveTester{
			iTyp:       "string",
			iMinLength: 3,
			iMaxLength: 6,
			iArgs:      []any{3, 6, "abc"},
		}.assertVector())
		t.Run(fmt.Sprintf("%d string length of zero", i), vectorOfDirectiveTester{
			iTyp:       "string",
			iMinLength: 0,
			iMaxLength: 0,
			iArgs:      []any{3, 6, "abc"},
		}.assertVector())
		t.Run(fmt.Sprintf("%d string equal lengths", i), vectorOfDirectiveTester{
			iTyp:       "string",
			iMinLength: 6,
			iMaxLength: 6,
			iArgs:      []any{3, 6, "abc"},
		}.assertVector())
		t.Run(fmt.Sprintf("%d string emojis", i), vectorOfDirectiveTester{
			iTyp:       "string",
			iMinLength: 3,
			iMaxLength: 6,
			iArgs:      []any{3, 6, "🔴🟡🟢"},
		}.assertVector())
		t.Run(fmt.Sprintf("%d string string length of zero", i), vectorOfDirectiveTester{
			iTyp:       "string",
			iMinLength: 3,
			iMaxLength: 6,
			iArgs:      []any{0, 0, "abc"},
		}.assertVector())
		t.Run(fmt.Sprintf("%d string equal string lengths", i), vectorOfDirectiveTester{
			iTyp:       "string",
			iMinLength: 3,
			iMaxLength: 6,
			iArgs:      []any{6, 6, "abc"},
		}.assertVector())
		t.Run(fmt.Sprintf("%d boolean baseline", i), vectorOfDirectiveTester{
			iTyp:       "bool",
			iMinLength: 3,
			iMaxLength: 6,
		}.assertVector())
		t.Run(fmt.Sprintf("%d bool length of zero", i), vectorOfDirectiveTester{
			iTyp:       "bool",
			iMinLength: 0,
			iMaxLength: 0,
		}.assertVector())
		t.Run(fmt.Sprintf("%d bool equal lengths", i), vectorOfDirectiveTester{
			iTyp:       "bool",
			iMinLength: 6,
			iMaxLength: 6,
		}.assertVector())
		t.Run(fmt.Sprintf("%d primitive baseline", i), vectorOfDirectiveTester{
			iTyp:       "primitive",
			iMinLength: 3,
			iMaxLength: 6,
		}.assertVector())
		t.Run(fmt.Sprintf("%d primitive length of zero", i), vectorOfDirectiveTester{
			iTyp:       "primitive",
			iMinLength: 0,
			iMaxLength: 0,
		}.assertVector())
		t.Run(fmt.Sprintf("%d primitive equal lengths", i), vectorOfDirectiveTester{
			iTyp:       "primitive",
			iMinLength: 6,
			iMaxLength: 6,
		}.assertVector())
	}
	t.Run("null invalid min length", vectorOfDirectiveTester{
		iTyp:       "null",
		iMinLength: -1,
		iMaxLength: 6,
		oErr:       errors.New("freeformgen: vector cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("null invalid max length", vectorOfDirectiveTester{
		iTyp:       "null",
		iMinLength: 3,
		iMaxLength: -1,
		oErr:       errors.New("freeformgen: vector cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("null min length greater than max length", vectorOfDirectiveTester{
		iTyp:       "null",
		iMinLength: 6,
		iMaxLength: 3,
		oErr:       errors.New("freeformgen: min length cannot exceed max length"),
	}.assertMinGreaterThanMaxError())
	t.Run("int invalid min length", vectorOfDirectiveTester{
		iTyp:       "int",
		iMinLength: -1,
		iMaxLength: 6,
		iArgs:      []any{0, 3},
		oErr:       errors.New("freeformgen: vector cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("int invalid max length", vectorOfDirectiveTester{
		iTyp:       "int",
		iMinLength: 3,
		iMaxLength: -1,
		iArgs:      []any{0, 3},
		oErr:       errors.New("freeformgen: vector cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("int min length greater than max length", vectorOfDirectiveTester{
		iTyp:       "int",
		iMinLength: 6,
		iMaxLength: 3,
		iArgs:      []any{0, 3},
		oErr:       errors.New("freeformgen: min length cannot exceed max length"),
	}.assertMinGreaterThanMaxError())
	t.Run("int min greater than max", vectorOfDirectiveTester{
		iTyp:       "int",
		iMinLength: 3,
		iMaxLength: 6,
		iArgs:      []any{1, -1},
		oErr:       errors.New("freeformgen: min cannot exceed max"),
	}.assertMinGreaterThanMaxError())
	t.Run("int incorrect args", vectorOfDirectiveTester{
		iTyp:       "int",
		iMinLength: 6,
		iMaxLength: 3,
		iArgs:      []any{1, -1, 24},
		oErr:       errors.New("freeformgen: wrong number of args"),
	}.assertIncorrectArgsError())
	t.Run("float invalid min length", vectorOfDirectiveTester{
		iTyp:       "float",
		iMinLength: -1,
		iMaxLength: 6,
		iArgs:      []any{0.0, 1.0},
		oErr:       errors.New("freeformgen: vector cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("float invalid max length", vectorOfDirectiveTester{
		iTyp:       "float",
		iMinLength: 3,
		iMaxLength: -1,
		iArgs:      []any{0.0, 1.0},
		oErr:       errors.New("freeformgen: vector cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("float min length greater than max length", vectorOfDirectiveTester{
		iTyp:       "float",
		iMinLength: 6,
		iMaxLength: 3,
		iArgs:      []any{0.0, 1.0},
		oErr:       errors.New("freeformgen: min length cannot exceed max length"),
	}.assertMinGreaterThanMaxError())
	t.Run("float min greater than max", vectorOfDirectiveTester{
		iTyp:       "float",
		iMinLength: 3,
		iMaxLength: 6,
		iArgs:      []any{1.0, -1.0},
		oErr:       errors.New("freeformgen: min cannot exceed max"),
	}.assertMinGreaterThanMaxError())
	t.Run("float incorrect args", vectorOfDirectiveTester{
		iTyp:       "float",
		iMinLength: 3,
		iMaxLength: 6,
		iArgs:      []any{1.0, -1.0, 24.0},
		oErr:       errors.New("freeformgen: wrong number of args"),
	}.assertIncorrectArgsError())
	t.Run("string invalid min length", vectorOfDirectiveTester{
		iTyp:       "string",
		iMinLength: -1,
		iMaxLength: 6,
		iArgs:      []any{3, 6, "abc"},
		oErr:       errors.New("freeformgen: vector cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("string invalid max length", vectorOfDirectiveTester{
		iTyp:       "string",
		iMinLength: 3,
		iMaxLength: -1,
		iArgs:      []any{3, 6, "abc"},
		oErr:       errors.New("freeformgen: vector cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("string min length greater than max length", vectorOfDirectiveTester{
		iTyp:       "string",
		iMinLength: 6,
		iMaxLength: 3,
		iArgs:      []any{3, 6, "abc"},
		oErr:       errors.New("freeformgen: min length cannot exceed max length"),
	}.assertMinGreaterThanMaxError())
	t.Run("string invalid min string length", vectorOfDirectiveTester{
		iTyp:       "string",
		iMinLength: 3,
		iMaxLength: 6,
		iArgs:      []any{-1, 6, "abc"},
		oErr:       errors.New("freeformgen: string cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("string invalid max string length", vectorOfDirectiveTester{
		iTyp:       "string",
		iMinLength: 3,
		iMaxLength: 6,
		iArgs:      []any{3, -1, "abc"},
		oErr:       errors.New("freeformgen: string cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("string min string length greater than max string length", vectorOfDirectiveTester{
		iTyp:       "string",
		iMinLength: 3,
		iMaxLength: 6,
		iArgs:      []any{6, 3, "abc"},
		oErr:       errors.New("freeformgen: min string length cannot exceed max string length"),
	}.assertMinGreaterThanMaxError())
	t.Run("string empty charset", vectorOfDirectiveTester{
		iTyp:       "string",
		iMinLength: 3,
		iMaxLength: 6,
		iArgs:      []any{3, 6, ""},
		oErr:       errors.New("freeformgen: charset cannot be empty"),
	}.assertEmptyCharsetError())
	t.Run("string incorrect args", vectorOfDirectiveTester{
		iTyp:       "string",
		iMinLength: 3,
		iMaxLength: 6,
		iArgs:      []any{3, 6, "", "foo"},
		oErr:       errors.New("freeformgen: wrong number of args"),
	}.assertIncorrectArgsError())
	t.Run("bool invalid min length", vectorOfDirectiveTester{
		iTyp:       "bool",
		iMinLength: -1,
		iMaxLength: 6,
		oErr:       errors.New("freeformgen: vector cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("bool invalid max length", vectorOfDirectiveTester{
		iTyp:       "bool",
		iMinLength: 3,
		iMaxLength: -1,
		oErr:       errors.New("freeformgen: vector cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("bool min length greater than max length", vectorOfDirectiveTester{
		iTyp:       "bool",
		iMinLength: 6,
		iMaxLength: 3,
		oErr:       errors.New("freeformgen: min length cannot exceed max length"),
	}.assertMinGreaterThanMaxError())
	t.Run("primitive invalid min length", vectorOfDirectiveTester{
		iTyp:       "primitive",
		iMinLength: -1,
		iMaxLength: 6,
		oErr:       errors.New("freeformgen: vector cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("primitive invalid max length", vectorOfDirectiveTester{
		iTyp:       "primitive",
		iMinLength: 3,
		iMaxLength: -1,
		oErr:       errors.New("freeformgen: vector cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("primitive min length greater than max length", vectorOfDirectiveTester{
		iTyp:       "primitive",
		iMinLength: 6,
		iMaxLength: 3,
		oErr:       errors.New("freeformgen: min length cannot exceed max length"),
	}.assertMinGreaterThanMaxError())
	t.Run("invalid type", vectorOfDirectiveTester{
		iTyp:       "foo",
		iMinLength: 6,
		iMaxLength: 3,
		oErr:       errors.New(`freeformgen: invalid type "foo"`),
	}.assertMinGreaterThanMaxError())
}
