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

func (tester vectorOfDirectiveTester) assertInvalidTypeError() func(*testing.T) {
	return func(t *testing.T) {
		oVector, got := vectorOfDirective(tester.iTyp, tester.iMinLength, tester.iMaxLength, tester.iArgs...)
		assertZeroed(t, oVector)
		assertInvalidTypeError(t, got, tester.oErr)
	}
}

type matrixOfDirectiveTester struct {
	iTyp         string
	iMinRowCount int
	iMaxRowCount int
	iMinColCount int
	iMaxColCount int
	iArgs        []any
	oErr         error
}

func (tester matrixOfDirectiveTester) assertMatrix() func(*testing.T) {
	return func(t *testing.T) {
		got, oErr := matrixOfDirective(tester.iTyp, tester.iMinRowCount, tester.iMaxRowCount, tester.iMinColCount, tester.iMaxColCount, tester.iArgs...)
		assertZeroed(t, oErr)
		assertWildPrimitiveMatrix(t, got, tester.iMinRowCount, tester.iMaxRowCount, tester.iMinColCount, tester.iMaxColCount)
	}
}

func (tester matrixOfDirectiveTester) assertInvalidRowCountError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := matrixOfDirective(tester.iTyp, tester.iMinRowCount, tester.iMaxRowCount, tester.iMinColCount, tester.iMaxColCount, tester.iArgs...)
		assertZeroed(t, oMatrix)
		assertInvalidRowCountError(t, got, tester.oErr)
	}
}

func (tester matrixOfDirectiveTester) assertInvalidColCountError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := matrixOfDirective(tester.iTyp, tester.iMinRowCount, tester.iMaxRowCount, tester.iMinColCount, tester.iMaxColCount, tester.iArgs...)
		assertZeroed(t, oMatrix)
		assertInvalidColCountError(t, got, tester.oErr)
	}
}

func (tester matrixOfDirectiveTester) assertInvalidLengthError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := matrixOfDirective(tester.iTyp, tester.iMinRowCount, tester.iMaxRowCount, tester.iMinColCount, tester.iMaxColCount, tester.iArgs...)
		assertZeroed(t, oMatrix)
		assertInvalidLengthError(t, got, tester.oErr)
	}
}

func (tester matrixOfDirectiveTester) assertMinGreaterThanMaxError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := matrixOfDirective(tester.iTyp, tester.iMinRowCount, tester.iMaxRowCount, tester.iMinColCount, tester.iMaxColCount, tester.iArgs...)
		assertZeroed(t, oMatrix)
		assertMinGreaterThanMaxError(t, got, tester.oErr)
	}
}

func (tester matrixOfDirectiveTester) assertIncorrectArgsError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := matrixOfDirective(tester.iTyp, tester.iMinRowCount, tester.iMaxRowCount, tester.iMinColCount, tester.iMaxColCount, tester.iArgs...)
		assertZeroed(t, oMatrix)
		assertIncorrectArgsError(t, got, tester.oErr)
	}
}

func (tester matrixOfDirectiveTester) assertEmptyCharsetError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := matrixOfDirective(tester.iTyp, tester.iMinRowCount, tester.iMaxRowCount, tester.iMinColCount, tester.iMaxColCount, tester.iArgs...)
		assertZeroed(t, oMatrix)
		assertEmptyCharsetError(t, got, tester.oErr)
	}
}

func (tester matrixOfDirectiveTester) assertInvalidTypeError() func(*testing.T) {
	return func(t *testing.T) {
		oMatrix, got := matrixOfDirective(tester.iTyp, tester.iMinRowCount, tester.iMaxRowCount, tester.iMinColCount, tester.iMaxColCount, tester.iArgs...)
		assertZeroed(t, oMatrix)
		assertInvalidTypeError(t, got, tester.oErr)
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
		iArgs:      []any{0, 3, 24},
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
		iArgs:      []any{0.0, 1.0, 24.0},
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
	}.assertInvalidTypeError())
}

func TestMatrixOfDirective(t *testing.T) {
	for i := 0; i < limit; i++ {
		t.Run(fmt.Sprintf("%d null baseline", i), matrixOfDirectiveTester{
			iTyp:         "null",
			iMinRowCount: 3,
			iMaxRowCount: 6,
			iMinColCount: 3,
			iMaxColCount: 6,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d null row count of zero", i), matrixOfDirectiveTester{
			iTyp:         "null",
			iMinRowCount: 0,
			iMaxRowCount: 0,
			iMinColCount: 3,
			iMaxColCount: 6,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d null column count of zero", i), matrixOfDirectiveTester{
			iTyp:         "null",
			iMinRowCount: 3,
			iMaxRowCount: 6,
			iMinColCount: 0,
			iMaxColCount: 0,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d null square matrix", i), matrixOfDirectiveTester{
			iTyp:         "null",
			iMinRowCount: 6,
			iMaxRowCount: 6,
			iMinColCount: 6,
			iMaxColCount: 6,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d int baseline", i), matrixOfDirectiveTester{
			iTyp:         "int",
			iMinRowCount: 3,
			iMaxRowCount: 6,
			iMinColCount: 3,
			iMaxColCount: 6,
			iArgs:        []any{0, 3},
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d int row count of zero", i), matrixOfDirectiveTester{
			iTyp:         "int",
			iMinRowCount: 0,
			iMaxRowCount: 0,
			iMinColCount: 3,
			iMaxColCount: 6,
			iArgs:        []any{0, 3},
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d int column count of zero", i), matrixOfDirectiveTester{
			iTyp:         "int",
			iMinRowCount: 3,
			iMaxRowCount: 6,
			iMinColCount: 0,
			iMaxColCount: 0,
			iArgs:        []any{0, 3},
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d int square matrix", i), matrixOfDirectiveTester{
			iTyp:         "int",
			iMinRowCount: 6,
			iMaxRowCount: 6,
			iMinColCount: 6,
			iMaxColCount: 6,
			iArgs:        []any{0, 3},
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d int broad range", i), matrixOfDirectiveTester{
			iTyp:         "int",
			iMinRowCount: 3,
			iMaxRowCount: 6,
			iMinColCount: 3,
			iMaxColCount: 6,
			iArgs:        []any{-10, 10},
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d int equal", i), matrixOfDirectiveTester{
			iTyp:         "int",
			iMinRowCount: 3,
			iMaxRowCount: 6,
			iMinColCount: 3,
			iMaxColCount: 6,
			iArgs:        []any{0, 0},
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d float baseline", i), matrixOfDirectiveTester{
			iTyp:         "float",
			iMinRowCount: 3,
			iMaxRowCount: 6,
			iMinColCount: 3,
			iMaxColCount: 6,
			iArgs:        []any{0.0, 1.0},
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d float row count of zero", i), matrixOfDirectiveTester{
			iTyp:         "float",
			iMinRowCount: 0,
			iMaxRowCount: 0,
			iMinColCount: 3,
			iMaxColCount: 6,
			iArgs:        []any{0.0, 1.0},
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d float column count of zero", i), matrixOfDirectiveTester{
			iTyp:         "float",
			iMinRowCount: 3,
			iMaxRowCount: 6,
			iMinColCount: 0,
			iMaxColCount: 0,
			iArgs:        []any{0.0, 1.0},
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d float square matrix", i), matrixOfDirectiveTester{
			iTyp:         "float",
			iMinRowCount: 6,
			iMaxRowCount: 6,
			iMinColCount: 6,
			iMaxColCount: 6,
			iArgs:        []any{0.0, 1.0},
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d float broad range", i), matrixOfDirectiveTester{
			iTyp:         "float",
			iMinRowCount: 3,
			iMaxRowCount: 6,
			iMinColCount: 3,
			iMaxColCount: 6,
			iArgs:        []any{-10.0, 10.0},
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d float equal", i), matrixOfDirectiveTester{
			iTyp:         "float",
			iMinRowCount: 3,
			iMaxRowCount: 6,
			iMinColCount: 3,
			iMaxColCount: 6,
			iArgs:        []any{0.0, 0.0},
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d string baseline", i), matrixOfDirectiveTester{
			iTyp:         "string",
			iMinRowCount: 3,
			iMaxRowCount: 6,
			iMinColCount: 3,
			iMaxColCount: 6,
			iArgs:        []any{3, 6, "abc"},
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d string row count of zero", i), matrixOfDirectiveTester{
			iTyp:         "string",
			iMinRowCount: 0,
			iMaxRowCount: 0,
			iMinColCount: 3,
			iMaxColCount: 6,
			iArgs:        []any{3, 6, "abc"},
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d string column count of zero", i), matrixOfDirectiveTester{
			iTyp:         "string",
			iMinRowCount: 3,
			iMaxRowCount: 6,
			iMinColCount: 0,
			iMaxColCount: 0,
			iArgs:        []any{3, 6, "abc"},
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d string square matrix", i), matrixOfDirectiveTester{
			iTyp:         "string",
			iMinRowCount: 6,
			iMaxRowCount: 6,
			iMinColCount: 6,
			iMaxColCount: 6,
			iArgs:        []any{3, 6, "abc"},
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d string emojis", i), matrixOfDirectiveTester{
			iTyp:         "string",
			iMinRowCount: 3,
			iMaxRowCount: 6,
			iMinColCount: 3,
			iMaxColCount: 6,
			iArgs:        []any{3, 6, "🔴🟡🟢"},
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d string string length of zero", i), matrixOfDirectiveTester{
			iTyp:         "string",
			iMinRowCount: 3,
			iMaxRowCount: 6,
			iMinColCount: 3,
			iMaxColCount: 6,
			iArgs:        []any{0, 0, "abc"},
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d string equal string lengths", i), matrixOfDirectiveTester{
			iTyp:         "string",
			iMinRowCount: 3,
			iMaxRowCount: 6,
			iMinColCount: 3,
			iMaxColCount: 6,
			iArgs:        []any{6, 6, "abc"},
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d bool baseline", i), matrixOfDirectiveTester{
			iTyp:         "bool",
			iMinRowCount: 3,
			iMaxRowCount: 6,
			iMinColCount: 3,
			iMaxColCount: 6,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d bool row count of zero", i), matrixOfDirectiveTester{
			iTyp:         "bool",
			iMinRowCount: 0,
			iMaxRowCount: 0,
			iMinColCount: 3,
			iMaxColCount: 6,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d bool column count of zero", i), matrixOfDirectiveTester{
			iTyp:         "bool",
			iMinRowCount: 3,
			iMaxRowCount: 6,
			iMinColCount: 0,
			iMaxColCount: 0,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d bool square matrix", i), matrixOfDirectiveTester{
			iTyp:         "bool",
			iMinRowCount: 6,
			iMaxRowCount: 6,
			iMinColCount: 6,
			iMaxColCount: 6,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d primitive baseline", i), matrixOfDirectiveTester{
			iTyp:         "primitive",
			iMinRowCount: 3,
			iMaxRowCount: 6,
			iMinColCount: 3,
			iMaxColCount: 6,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d primitive row count of zero", i), matrixOfDirectiveTester{
			iTyp:         "primitive",
			iMinRowCount: 0,
			iMaxRowCount: 0,
			iMinColCount: 3,
			iMaxColCount: 6,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d primitive column count of zero", i), matrixOfDirectiveTester{
			iTyp:         "primitive",
			iMinRowCount: 3,
			iMaxRowCount: 6,
			iMinColCount: 0,
			iMaxColCount: 0,
		}.assertMatrix())
		t.Run(fmt.Sprintf("%d primitive square matrix", i), matrixOfDirectiveTester{
			iTyp:         "primitive",
			iMinRowCount: 6,
			iMaxRowCount: 6,
			iMinColCount: 6,
			iMaxColCount: 6,
		}.assertMatrix())
	}
	t.Run("null invalid min row count", matrixOfDirectiveTester{
		iTyp:         "null",
		iMinRowCount: -1,
		iMaxRowCount: 6,
		iMinColCount: 3,
		iMaxColCount: 6,
		oErr:         errors.New("freeformgen: matrix cannot have a negative row count"),
	}.assertInvalidRowCountError())
	t.Run("null invalid max row count", matrixOfDirectiveTester{
		iTyp:         "null",
		iMinRowCount: 3,
		iMaxRowCount: -1,
		iMinColCount: 3,
		iMaxColCount: 6,
		oErr:         errors.New("freeformgen: matrix cannot have a negative row count"),
	}.assertInvalidRowCountError())
	t.Run("null invalid min column count", matrixOfDirectiveTester{
		iTyp:         "null",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: -1,
		iMaxColCount: 6,
		oErr:         errors.New("freeformgen: matrix cannot have a negative column count"),
	}.assertInvalidColCountError())
	t.Run("null invalid max column count", matrixOfDirectiveTester{
		iTyp:         "null",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: 3,
		iMaxColCount: -1,
		oErr:         errors.New("freeformgen: matrix cannot have a negative column count"),
	}.assertInvalidColCountError())
	t.Run("null min row count greater than max row count", matrixOfDirectiveTester{
		iTyp:         "null",
		iMinRowCount: 6,
		iMaxRowCount: 3,
		iMinColCount: 3,
		iMaxColCount: 6,
		oErr:         errors.New("freeformgen: min row count cannot exceed max row count"),
	}.assertMinGreaterThanMaxError())
	t.Run("null min column count greater than max column count", matrixOfDirectiveTester{
		iTyp:         "null",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: 6,
		iMaxColCount: 3,
		oErr:         errors.New("freeformgen: min column count cannot exceed max column count"),
	}.assertMinGreaterThanMaxError())
	t.Run("int invalid min row count", matrixOfDirectiveTester{
		iTyp:         "int",
		iMinRowCount: -1,
		iMaxRowCount: 6,
		iMinColCount: 3,
		iMaxColCount: 6,
		iArgs:        []any{0, 3},
		oErr:         errors.New("freeformgen: matrix cannot have a negative row count"),
	}.assertInvalidRowCountError())
	t.Run("int invalid max row count", matrixOfDirectiveTester{
		iTyp:         "int",
		iMinRowCount: 3,
		iMaxRowCount: -1,
		iMinColCount: 3,
		iMaxColCount: 6,
		iArgs:        []any{0, 3},
		oErr:         errors.New("freeformgen: matrix cannot have a negative row count"),
	}.assertInvalidRowCountError())
	t.Run("int invalid min column count", matrixOfDirectiveTester{
		iTyp:         "int",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: -1,
		iMaxColCount: 6,
		iArgs:        []any{0, 3},
		oErr:         errors.New("freeformgen: matrix cannot have a negative column count"),
	}.assertInvalidColCountError())
	t.Run("int invalid min column count", matrixOfDirectiveTester{
		iTyp:         "int",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: 3,
		iMaxColCount: -1,
		iArgs:        []any{0, 3},
		oErr:         errors.New("freeformgen: matrix cannot have a negative column count"),
	}.assertInvalidColCountError())
	t.Run("int min row count greater than max row count", matrixOfDirectiveTester{
		iTyp:         "int",
		iMinRowCount: 6,
		iMaxRowCount: 3,
		iMinColCount: 3,
		iMaxColCount: 6,
		iArgs:        []any{0, 3},
		oErr:         errors.New("freeformgen: min row count cannot exceed max row count"),
	}.assertMinGreaterThanMaxError())
	t.Run("int min column count greater than max column count", matrixOfDirectiveTester{
		iTyp:         "int",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: 6,
		iMaxColCount: 3,
		iArgs:        []any{0, 3},
		oErr:         errors.New("freeformgen: min column count cannot exceed max column count"),
	}.assertMinGreaterThanMaxError())
	t.Run("int min greater than max", matrixOfDirectiveTester{
		iTyp:         "int",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: 3,
		iMaxColCount: 6,
		iArgs:        []any{1, -1},
		oErr:         errors.New("freeformgen: min cannot exceed max"),
	}.assertMinGreaterThanMaxError())
	t.Run("int incorrect args", matrixOfDirectiveTester{
		iTyp:         "int",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: 3,
		iMaxColCount: 6,
		iArgs:        []any{0, 3, 24},
		oErr:         errors.New("freeformgen: wrong number of args"),
	}.assertIncorrectArgsError())
	t.Run("float invalid min row count", matrixOfDirectiveTester{
		iTyp:         "float",
		iMinRowCount: -1,
		iMaxRowCount: 6,
		iMinColCount: 3,
		iMaxColCount: 6,
		iArgs:        []any{0.0, 1.0},
		oErr:         errors.New("freeformgen: matrix cannot have a negative row count"),
	}.assertInvalidRowCountError())
	t.Run("float invalid max row count", matrixOfDirectiveTester{
		iTyp:         "float",
		iMinRowCount: 3,
		iMaxRowCount: -1,
		iMinColCount: 3,
		iMaxColCount: 6,
		iArgs:        []any{0.0, 1.0},
		oErr:         errors.New("freeformgen: matrix cannot have a negative row count"),
	}.assertInvalidRowCountError())
	t.Run("float invalid min column count", matrixOfDirectiveTester{
		iTyp:         "float",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: -1,
		iMaxColCount: 6,
		iArgs:        []any{0.0, 1.0},
		oErr:         errors.New("freeformgen: matrix cannot have a negative column count"),
	}.assertInvalidColCountError())
	t.Run("float invalid min column count", matrixOfDirectiveTester{
		iTyp:         "float",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: 3,
		iMaxColCount: -1,
		iArgs:        []any{0.0, 1.0},
		oErr:         errors.New("freeformgen: matrix cannot have a negative column count"),
	}.assertInvalidColCountError())
	t.Run("float min row count greater than max row count", matrixOfDirectiveTester{
		iTyp:         "float",
		iMinRowCount: 6,
		iMaxRowCount: 3,
		iMinColCount: 3,
		iMaxColCount: 6,
		iArgs:        []any{0.0, 1.0},
		oErr:         errors.New("freeformgen: min row count cannot exceed max row count"),
	}.assertMinGreaterThanMaxError())
	t.Run("float min column count greater than max column count", matrixOfDirectiveTester{
		iTyp:         "float",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: 6,
		iMaxColCount: 3,
		iArgs:        []any{0.0, 1.0},
		oErr:         errors.New("freeformgen: min column count cannot exceed max column count"),
	}.assertMinGreaterThanMaxError())
	t.Run("float min greater than max", matrixOfDirectiveTester{
		iTyp:         "float",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: 3,
		iMaxColCount: 6,
		iArgs:        []any{1.0, -1.0},
		oErr:         errors.New("freeformgen: min cannot exceed max"),
	}.assertMinGreaterThanMaxError())
	t.Run("float incorrect args", matrixOfDirectiveTester{
		iTyp:         "float",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: 3,
		iMaxColCount: 6,
		iArgs:        []any{0.0, 1.0, 24.0},
		oErr:         errors.New("freeformgen: wrong number of args"),
	}.assertIncorrectArgsError())
	t.Run("string invalid min row count", matrixOfDirectiveTester{
		iTyp:         "string",
		iMinRowCount: -1,
		iMaxRowCount: 6,
		iMinColCount: 3,
		iMaxColCount: 6,
		iArgs:        []any{3, 6, "abc"},
		oErr:         errors.New("freeformgen: matrix cannot have a negative row count"),
	}.assertInvalidRowCountError())
	t.Run("string invalid max row count", matrixOfDirectiveTester{
		iTyp:         "string",
		iMinRowCount: 3,
		iMaxRowCount: -1,
		iMinColCount: 3,
		iMaxColCount: 6,
		iArgs:        []any{3, 6, "abc"},
		oErr:         errors.New("freeformgen: matrix cannot have a negative row count"),
	}.assertInvalidRowCountError())
	t.Run("string invalid min column count", matrixOfDirectiveTester{
		iTyp:         "string",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: -1,
		iMaxColCount: 6,
		iArgs:        []any{3, 6, "abc"},
		oErr:         errors.New("freeformgen: matrix cannot have a negative column count"),
	}.assertInvalidColCountError())
	t.Run("string invalid max column count", matrixOfDirectiveTester{
		iTyp:         "string",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: 3,
		iMaxColCount: -1,
		iArgs:        []any{3, 6, "abc"},
		oErr:         errors.New("freeformgen: matrix cannot have a negative column count"),
	}.assertInvalidColCountError())
	t.Run("string min row count greater than max row count", matrixOfDirectiveTester{
		iTyp:         "string",
		iMinRowCount: 6,
		iMaxRowCount: 3,
		iMinColCount: 3,
		iMaxColCount: 6,
		iArgs:        []any{3, 6, "abc"},
		oErr:         errors.New("freeformgen: min row count cannot exceed max row count"),
	}.assertMinGreaterThanMaxError())
	t.Run("string min column count greater than max column count", matrixOfDirectiveTester{
		iTyp:         "string",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: 6,
		iMaxColCount: 3,
		iArgs:        []any{3, 6, "abc"},
		oErr:         errors.New("freeformgen: min column count cannot exceed max column count"),
	}.assertMinGreaterThanMaxError())
	t.Run("string invalid min string length", matrixOfDirectiveTester{
		iTyp:         "string",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: 3,
		iMaxColCount: 6,
		iArgs:        []any{-1, 6, "abc"},
		oErr:         errors.New("freeformgen: string cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("string invalid max string length", matrixOfDirectiveTester{
		iTyp:         "string",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: 3,
		iMaxColCount: 6,
		iArgs:        []any{3, -1, "abc"},
		oErr:         errors.New("freeformgen: string cannot have a negative length"),
	}.assertInvalidLengthError())
	t.Run("string min string length greater than max string length", matrixOfDirectiveTester{
		iTyp:         "string",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: 3,
		iMaxColCount: 6,
		iArgs:        []any{6, 3, "abc"},
		oErr:         errors.New("freeformgen: min string length cannot exceed max string length"),
	}.assertMinGreaterThanMaxError())
	t.Run("string empty charset", matrixOfDirectiveTester{
		iTyp:         "string",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: 3,
		iMaxColCount: 6,
		iArgs:        []any{3, 6, ""},
		oErr:         errors.New("freeformgen: charset cannot be empty"),
	}.assertEmptyCharsetError())
	t.Run("string incorrect args", matrixOfDirectiveTester{
		iTyp:         "string",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: 3,
		iMaxColCount: 6,
		iArgs:        []any{3, 6, "", "foo"},
		oErr:         errors.New("freeformgen: wrong number of args"),
	}.assertIncorrectArgsError())
	t.Run("bool invalid min row count", matrixOfDirectiveTester{
		iTyp:         "bool",
		iMinRowCount: -1,
		iMaxRowCount: 6,
		iMinColCount: 3,
		iMaxColCount: 6,
		oErr:         errors.New("freeformgen: matrix cannot have a negative row count"),
	}.assertInvalidRowCountError())
	t.Run("bool invalid max row count", matrixOfDirectiveTester{
		iTyp:         "bool",
		iMinRowCount: 3,
		iMaxRowCount: -1,
		iMinColCount: 3,
		iMaxColCount: 6,
		oErr:         errors.New("freeformgen: matrix cannot have a negative row count"),
	}.assertInvalidRowCountError())
	t.Run("bool invalid min column count", matrixOfDirectiveTester{
		iTyp:         "bool",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: -1,
		iMaxColCount: 6,
		oErr:         errors.New("freeformgen: matrix cannot have a negative column count"),
	}.assertInvalidColCountError())
	t.Run("bool invalid max column count", matrixOfDirectiveTester{
		iTyp:         "bool",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: 3,
		iMaxColCount: -1,
		oErr:         errors.New("freeformgen: matrix cannot have a negative column count"),
	}.assertInvalidColCountError())
	t.Run("bool min row count greater than max row count", matrixOfDirectiveTester{
		iTyp:         "bool",
		iMinRowCount: 6,
		iMaxRowCount: 3,
		iMinColCount: 3,
		iMaxColCount: 6,
		oErr:         errors.New("freeformgen: min row count cannot exceed max row count"),
	}.assertMinGreaterThanMaxError())
	t.Run("bool min column count greater than max column count", matrixOfDirectiveTester{
		iTyp:         "bool",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: 6,
		iMaxColCount: 3,
		oErr:         errors.New("freeformgen: min column count cannot exceed max column count"),
	}.assertMinGreaterThanMaxError())
	t.Run("primitive invalid min row count", matrixOfDirectiveTester{
		iTyp:         "primitive",
		iMinRowCount: -1,
		iMaxRowCount: 6,
		iMinColCount: 3,
		iMaxColCount: 6,
		oErr:         errors.New("freeformgen: matrix cannot have a negative row count"),
	}.assertInvalidRowCountError())
	t.Run("primitive invalid max row count", matrixOfDirectiveTester{
		iTyp:         "primitive",
		iMinRowCount: 3,
		iMaxRowCount: -1,
		iMinColCount: 3,
		iMaxColCount: 6,
		oErr:         errors.New("freeformgen: matrix cannot have a negative row count"),
	}.assertInvalidRowCountError())
	t.Run("primitive invalid min column count", matrixOfDirectiveTester{
		iTyp:         "primitive",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: -1,
		iMaxColCount: 6,
		oErr:         errors.New("freeformgen: matrix cannot have a negative column count"),
	}.assertInvalidColCountError())
	t.Run("primitive invalid max column count", matrixOfDirectiveTester{
		iTyp:         "primitive",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: 3,
		iMaxColCount: -1,
		oErr:         errors.New("freeformgen: matrix cannot have a negative column count"),
	}.assertInvalidColCountError())
	t.Run("primitive min row count greater than max row count", matrixOfDirectiveTester{
		iTyp:         "primitive",
		iMinRowCount: 6,
		iMaxRowCount: 3,
		iMinColCount: 3,
		iMaxColCount: 6,
		oErr:         errors.New("freeformgen: min row count cannot exceed max row count"),
	}.assertMinGreaterThanMaxError())
	t.Run("primitive min column count greater than max column count", matrixOfDirectiveTester{
		iTyp:         "primitive",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: 6,
		iMaxColCount: 3,
		oErr:         errors.New("freeformgen: min column count cannot exceed max column count"),
	}.assertMinGreaterThanMaxError())
	t.Run("invalid type", matrixOfDirectiveTester{
		iTyp:         "foo",
		iMinRowCount: 3,
		iMaxRowCount: 6,
		iMinColCount: 3,
		iMaxColCount: 6,
		oErr:         errors.New(`freeformgen: invalid type "foo"`),
	}.assertInvalidTypeError())
}
