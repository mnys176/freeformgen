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

func integerDirective(min, max int) (int, error) {
	if min > max {
		return 0, freeformgenError{errors.New("min cannot exceed max")}
	}
	if min == max {
		return min, nil
	}
	return rand.Intn(max-min+1) + min, nil
}

func floatDirective(min, max float64) (float64, error) {
	if min > max {
		return 0.0, freeformgenError{errors.New("min cannot exceed max")}
	}
	if min == max {
		return min, nil
	}
	return rand.Float64()*(float64(max)-float64(min)) + float64(min), nil
}

func stringDirective(minLength, maxLength int, charset string) (string, error) {
	if minLength < 0 || maxLength < 0 {
		return "", freeformgenError{errors.New("string cannot have a negative length")}
	}
	if minLength > maxLength {
		return "", freeformgenError{errors.New("min length cannot exceed max length")}
	}

	count, err := integerDirective(minLength, maxLength)
	if err != nil {
		return "", err
	}

	var b strings.Builder
	for n := 0; n < count; n++ {
		i := rand.Intn(utf8.RuneCountInString(charset))
		b.WriteRune([]rune(charset)[i])
	}
	return b.String(), nil
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
	var p any
	switch rand.Intn(5) {
	case 0:
		p, _ = integerDirective(minRandomInteger, maxRandomInteger)
	case 1:
		p, _ = floatDirective(minRandomFloat, maxRandomFloat)
	case 2:
		p, _ = stringDirective(0, maxRandomStringLength, randomCharset)
	case 3:
		p = booleanDirective()
	default:
		p = nullDirective()
	}
	return p
}

func vectorDirective(minLength, maxLength int) ([]any, error) {
	if minLength < 0 || maxLength < 0 {
		return nil, freeformgenError{errors.New("vector cannot have a negative length")}
	}
	if minLength > maxLength {
		return nil, freeformgenError{errors.New("min length cannot exceed max length")}
	}

	count, err := integerDirective(minLength, maxLength)
	if err != nil {
		return nil, err
	}

	vec := make([]any, 0)
	for i := 0; i < count; i++ {
		vec = append(vec, primitiveDirective())
	}
	return vec, nil
}

// func matrixDirective(typ string, minRows, maxRows, minCols, maxCols int) []any {
// 	if minRows < 0 || maxRows < 0 {
// 		panic(freeformgenError{errors.New("matrix cannot have a negative number of rows")})
// 	}
// 	if minCols < 0 || maxCols < 0 {
// 		panic(freeformgenError{errors.New("matrix cannot have a negative number of columns")})
// 	}

// arr := make([]any, 0)
//
//	for i := 0; i < integerDirective(minLength, maxLength); i++ {
//		switch typ {
//		case "integer":
//			arr = append(arr, integerDirective(minRandomInteger, maxRandomInteger))
//		case "float":
//			arr = append(arr, floatDirective(minRandomFloat, maxRandomFloat))
//		case "string":
//			arr = append(arr, stringDirective(0, maxRandomStringLength, randomCharset))
//		case "boolean":
//			arr = append(arr, booleanDirective())
//		case "null":
//			arr = append(arr, nullDirective())
//		case "":
//			arr = append(arr, primitiveDirective())
//		default:
//			panic(freeformgenError{fmt.Errorf("unknown primitive %q", typ)})
//		}
//	}
//
// return arr
// return nil, nil
// }
