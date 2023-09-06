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
	if charset == "" {
		return "", freeformgenError{errors.New("charset cannot be empty")}
	}

	count, _ := integerDirective(minLength, maxLength)
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
	minRandomStringLength int     = 0
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
		p, _ = stringDirective(minRandomStringLength, maxRandomStringLength, randomCharset)
	case 3:
		p = booleanDirective()
	default:
		p = nullDirective()
	}
	return p
}

func vNullDirective(minLength, maxLength int) ([]any, error) {
	if minLength < 0 || maxLength < 0 {
		return nil, freeformgenError{errors.New("vector cannot have a negative length")}
	}
	if minLength > maxLength {
		return nil, freeformgenError{errors.New("min length cannot exceed max length")}
	}

	count, _ := integerDirective(minLength, maxLength)
	vec := make([]any, count)
	for i := 0; i < count; i++ {
		vec[i] = nullDirective()
	}
	return vec, nil
}

func vIntegerDirective(minLength, maxLength, min, max int) ([]int, error) {
	if minLength < 0 || maxLength < 0 {
		return nil, freeformgenError{errors.New("vector cannot have a negative length")}
	}
	if minLength > maxLength {
		return nil, freeformgenError{errors.New("min length cannot exceed max length")}
	}
	if min > max {
		return nil, freeformgenError{errors.New("min cannot exceed max")}
	}

	count, _ := integerDirective(minLength, maxLength)
	vec := make([]int, count)
	for i := 0; i < count; i++ {
		v, _ := integerDirective(min, max)
		vec[i] = v
	}
	return vec, nil
}

func vFloatDirective(minLength, maxLength int, min, max float64) ([]float64, error) {
	if minLength < 0 || maxLength < 0 {
		return nil, freeformgenError{errors.New("vector cannot have a negative length")}
	}
	if minLength > maxLength {
		return nil, freeformgenError{errors.New("min length cannot exceed max length")}
	}
	if min > max {
		return nil, freeformgenError{errors.New("min cannot exceed max")}
	}

	count, _ := integerDirective(minLength, maxLength)
	vec := make([]float64, count)
	for i := 0; i < count; i++ {
		v, _ := floatDirective(min, max)
		vec[i] = v
	}
	return vec, nil
}

func vStringDirective(minLength, maxLength, minStrLength, maxStrLength int, charset string) ([]string, error) {
	if minLength < 0 || maxLength < 0 {
		return nil, freeformgenError{errors.New("vector cannot have a negative length")}
	}
	if minLength > maxLength {
		return nil, freeformgenError{errors.New("min length cannot exceed max length")}
	}
	if minStrLength < 0 || maxStrLength < 0 {
		return nil, freeformgenError{errors.New("string cannot have a negative length")}
	}
	if minStrLength > maxStrLength {
		return nil, freeformgenError{errors.New("min string length cannot exceed max string length")}
	}
	if charset == "" {
		return nil, freeformgenError{errors.New("charset cannot be empty")}
	}

	count, _ := integerDirective(minLength, maxLength)
	vec := make([]string, count)
	for i := 0; i < count; i++ {
		v, _ := stringDirective(minStrLength, maxStrLength, charset)
		vec[i] = v
	}
	return vec, nil
}

func vBooleanDirective(minLength, maxLength int) ([]bool, error) {
	if minLength < 0 || maxLength < 0 {
		return nil, freeformgenError{errors.New("vector cannot have a negative length")}
	}
	if minLength > maxLength {
		return nil, freeformgenError{errors.New("min length cannot exceed max length")}
	}

	count, _ := integerDirective(minLength, maxLength)
	vec := make([]bool, count)
	for i := 0; i < count; i++ {
		vec[i] = booleanDirective()
	}
	return vec, nil
}

func vPrimitiveDirective(minLength, maxLength int) ([]any, error) {
	if minLength < 0 || maxLength < 0 {
		return nil, freeformgenError{errors.New("vector cannot have a negative length")}
	}
	if minLength > maxLength {
		return nil, freeformgenError{errors.New("min length cannot exceed max length")}
	}

	count, _ := integerDirective(minLength, maxLength)
	vec := make([]any, count)
	for i := 0; i < count; i++ {
		vec[i] = primitiveDirective()
	}
	return vec, nil
}

func mNullDirective(minRowCount, maxRowCount, minColCount, maxColCount int) ([][]any, error) {
	if minRowCount < 0 || maxRowCount < 0 {
		return nil, freeformgenError{errors.New("matrix cannot have a negative row count")}
	}
	if minColCount < 0 || maxColCount < 0 {
		return nil, freeformgenError{errors.New("matrix cannot have a negative column count")}
	}
	if minRowCount > maxRowCount {
		return nil, freeformgenError{errors.New("min row count cannot exceed max row count")}
	}
	if minColCount > maxColCount {
		return nil, freeformgenError{errors.New("min column count cannot exceed max column count")}
	}

	rowCount, _ := integerDirective(minRowCount, maxRowCount)
	mat := make([][]any, rowCount)
	for r := 0; r < rowCount; r++ {
		vec, _ := vNullDirective(minColCount, maxColCount)
		mat[r] = vec
	}
	return mat, nil
}

func mIntegerDirective(minRowCount, maxRowCount, minColCount, maxColCount, min, max int) ([][]int, error) {
	if minRowCount < 0 || maxRowCount < 0 {
		return nil, freeformgenError{errors.New("matrix cannot have a negative row count")}
	}
	if minColCount < 0 || maxColCount < 0 {
		return nil, freeformgenError{errors.New("matrix cannot have a negative column count")}
	}
	if minRowCount > maxRowCount {
		return nil, freeformgenError{errors.New("min row count cannot exceed max row count")}
	}
	if minColCount > maxColCount {
		return nil, freeformgenError{errors.New("min column count cannot exceed max column count")}
	}
	if min > max {
		return nil, freeformgenError{errors.New("min cannot exceed max")}
	}

	rowCount, _ := integerDirective(minRowCount, maxRowCount)
	mat := make([][]int, rowCount)
	for r := 0; r < rowCount; r++ {
		vec, _ := vIntegerDirective(minColCount, maxColCount, min, max)
		mat[r] = vec
	}
	return mat, nil
}
