package main

import (
	"errors"
	"testing"
)

type nullTester struct{}

func (tester nullTester) assertNil() func(*testing.T) {
	return func(t *testing.T) {
		got := null()
		assertNil(t, got)
	}
}

type numberTester[T int | float64] struct {
	iMin   T
	iMax   T
	oPanic error
}

func (tester numberTester[T]) assertNumber() func(*testing.T) {
	return func(t *testing.T) {
		got := number[T](tester.iMin, tester.iMax)
		assertNumber(t, got, tester.iMin, tester.iMax)
	}
}

func (tester numberTester[T]) assertMinGreaterThanMaxErrorPanic() func(*testing.T) {
	return func(t *testing.T) {
		defer assertMinGreaterThanMaxPanic(t, tester.oPanic)
		number[T](tester.iMin, tester.iMax)
	}
}

func TestNull(t *testing.T) {
	t.Run("baseline", nullTester{}.assertNil())
}

func TestNumber(t *testing.T) {
	t.Run("baseline", numberTester[float64]{
		iMin: 0.0,
		iMax: 1.0,
	}.assertNumber())
	t.Run("broad range", numberTester[float64]{
		iMin: -10.0,
		iMax: 10.0,
	}.assertNumber())
	t.Run("equal", numberTester[float64]{
		iMin: 0.0,
		iMax: 0.0,
	}.assertNumber())
	t.Run("min greater than max", numberTester[float64]{
		iMin:   10.0,
		iMax:   -10.0,
		oPanic: errors.New("freeformgen: min cannot exceed max"),
	}.assertMinGreaterThanMaxErrorPanic())
	t.Run("int baseline", numberTester[int]{
		iMin: 0,
		iMax: 1,
	}.assertNumber())
	t.Run("int broad range", numberTester[int]{
		iMin: -10,
		iMax: 10,
	}.assertNumber())
	t.Run("int equal", numberTester[int]{
		iMin: 0,
		iMax: 0,
	}.assertNumber())
	t.Run("int min greater than max", numberTester[int]{
		iMin:   1,
		iMax:   -1,
		oPanic: errors.New("freeformgen: min cannot exceed max"),
	}.assertMinGreaterThanMaxErrorPanic())
}
