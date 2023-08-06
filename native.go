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

func number[T int | float64](min, max T) T {
	if min > max {
		panic(freeformgenError{errors.New("min cannot exceed max")})
	}
	if min == max {
		return min
	}
	num := rand.Float64()*(float64(max)-float64(min)) + float64(min)
	return T(num)
}

func stringDirective(minLength, maxLength int, charset string) string {
	if minLength < 0 || maxLength < 0 {
		panic(freeformgenError{errors.New("string cannot have a negative length")})
	}
	if minLength > maxLength {
		panic(freeformgenError{errors.New("min length cannot exceed max length")})
	}

	var count int
	if minLength == maxLength {
		count = minLength
	} else {
		count = maxLength - minLength
	}

	var b strings.Builder
	for n := 0; n <= count; n++ {
		i := rand.Intn(utf8.RuneCountInString(charset))
		b.WriteRune([]rune(charset)[i])
	}
	return b.String()
}

func booleanDirective() bool {
	return rand.Intn(2) == 1
}
