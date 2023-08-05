package main

import (
	"errors"
	"math/rand"
)

func null() any {
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
