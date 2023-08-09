package main

import (
	"errors"
	"math/rand"
	"strings"
	"unicode/utf8"
)

func nullDirective() any {
	return nil
}

func integerDirective(min, max int) int {
	if min > max {
		panic(freeformgenError{errors.New("min cannot exceed max")})
	}
	if min == max {
		return min
	}
	return rand.Intn(max-min+1) + min
}

func floatDirective(min, max float64) float64 {
	if min > max {
		panic(freeformgenError{errors.New("min cannot exceed max")})
	}
	if min == max {
		return min
	}
	return rand.Float64()*(float64(max)-float64(min)) + float64(min)
}

func stringDirective(minLength, maxLength int, charset string) string {
	if minLength < 0 || maxLength < 0 {
		panic(freeformgenError{errors.New("string cannot have a negative length")})
	}
	if minLength > maxLength {
		panic(freeformgenError{errors.New("min length cannot exceed max length")})
	}

	var b strings.Builder
	for n := 0; n < integerDirective(minLength, maxLength); n++ {
		i := rand.Intn(utf8.RuneCountInString(charset))
		b.WriteRune([]rune(charset)[i])
	}
	return b.String()
}

func booleanDirective() bool {
	return rand.Intn(2) == 1
}

const (
	maxRandomInteger      int     = 0xFFFFFFFF
	minRandomInteger      int     = -maxRandomInteger
	maxRandomFloat        float64 = float64(maxRandomInteger)
	minRandomFloat        float64 = -maxRandomFloat
	maxRandomStringLength int     = 32
	randomCharset         string  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
)

func primitiveDirective() any {
	// Primitives include integers, floats, strings, booleans, and null.
	switch rand.Intn(5) {
	case 0:
		return integerDirective(minRandomInteger, maxRandomInteger)
	case 1:
		return floatDirective(minRandomFloat, maxRandomFloat)
	case 2:
		return stringDirective(0, maxRandomStringLength, randomCharset)
	case 3:
		return booleanDirective()
	default:
		return nil
	}
}

// func vectorDirective(typ string, minLength, maxLength int) []any {
// 	if minLength < 0 || maxLength < 0 {
// 		panic(freeformgenError{errors.New("vector cannot have a negative length")})
// 	}
// 	if minLength > maxLength {
// 		panic(freeformgenError{errors.New("min length cannot exceed max length")})
// 	}

// 	vec := make([]any, 0)
// 	for i := 0; i < integerDirective(minLength, maxLength); i++ {
// 		switch typ {
// 		case "integer":
// 			vec = append(vec, integerDirective(minRandomInteger, maxRandomInteger))
// 		case "float":
// 			vec = append(vec, floatDirective(minRandomFloat, maxRandomFloat))
// 		case "string":
// 			vec = append(vec, stringDirective(0, maxRandomStringLength, randomCharset))
// 		case "boolean":
// 			vec = append(vec, booleanDirective())
// 		case "null":
// 			vec = append(vec, nullDirective())
// 		case "":
// 			vec = append(vec, primitiveDirective())
// 		default:
// 			panic(freeformgenError{fmt.Errorf("unknown primitive %q", typ)})
// 		}
// 	}
// 	return vec
// }

// func matrixDirective(typ string, minRows, maxRows, minCols, maxCols int) []any {
// 	if minRows < 0 || maxRows < 0 {
// 		panic(freeformgenError{errors.New("matrix cannot have a negative number of rows")})
// 	}
// 	if minCols < 0 || maxCols < 0 {
// 		panic(freeformgenError{errors.New("matrix cannot have a negative number of columns")})
// 	}

// 	arr := make([]any, 0)
// 	for i := 0; i < integerDirective(minLength, maxLength); i++ {
// 		switch typ {
// 		case "integer":
// 			arr = append(arr, integerDirective(minRandomInteger, maxRandomInteger))
// 		case "float":
// 			arr = append(arr, floatDirective(minRandomFloat, maxRandomFloat))
// 		case "string":
// 			arr = append(arr, stringDirective(0, maxRandomStringLength, randomCharset))
// 		case "boolean":
// 			arr = append(arr, booleanDirective())
// 		case "null":
// 			arr = append(arr, nullDirective())
// 		case "":
// 			arr = append(arr, primitiveDirective())
// 		default:
// 			panic(freeformgenError{fmt.Errorf("unknown primitive %q", typ)})
// 		}
// 	}
// 	return arr
// }
