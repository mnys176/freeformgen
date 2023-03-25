package native

import (
	"math"
	"math/rand"
	"strings"
)

const (
	MaxGeneratedStringLength int = 64
	MaxGeneratedArrayLength  int = 12
	MaxGeneratedObjectLength int = 12
)

func generatePrimitive() any {
	selector := rand.Intn(5)
	switch selector {
	case 1:
		return IntegerDirective(math.MinInt32, math.MaxInt32)
	case 2:
		return FloatDirective(float64(math.MinInt32), float64(math.MaxInt32))
	case 3:
		return StringDirective(-1)
	case 4:
		return BooleanDirective()
	default:
		return nil
	}
}

func generateInternalArray(length int, nested bool) []any {
	if length < 0 {
		length = rand.Intn(MaxGeneratedArrayLength)
	}
	array := make([]any, 0, length)
	for i := 0; i < length; i++ {
		// Either 0 for a primitive or 1 for complex.
		selector := rand.Intn(2)
		switch selector {
		case 1:
			if !nested {
				array = append(array, generateInternalArray(-1, true))
			}
		case 2:
			if !nested {
				array = append(array, generateInternalObject(-1, true))
			}
		default:
			array = append(array, generatePrimitive())
		}
	}
	return array
}

func generateInternalObject(length int, nested bool) map[string]any {
	if length < 0 {
		length = rand.Intn(MaxGeneratedObjectLength)
	}
	object := make(map[string]any)
	for i := 0; i < length; i++ {
		key := StringDirective(-1)
		for _, ok := object[key]; ok; _, ok = object[key] {
			// Keep generating until a unique key is created...
		}

		// Either 0 for a primitive, 1 for array, or 2 for object.
		selector := rand.Intn(3)
		switch selector {
		case 1:
			if !nested {
				object[key] = generateInternalArray(-1, true)
			}
		case 2:
			if !nested {
				object[key] = generateInternalObject(-1, true)
			}
		default:
			object[key] = generatePrimitive()
		}
	}
	return object
}

func BooleanDirective() bool {
	return rand.Intn(2) == 1
}

func FloatDirective(min, max float64) float64 {
	// TODO: Panic if `max` is less than `min`.
	return rand.Float64()*(max-min) + min
}

func IntegerDirective(min, max int64) int64 {
	// TODO: Panic if `max` is less than `min`.
	return rand.Int63n(max-min+1) + min
}

func StringDirective(length int) string {
	if length < 0 {
		length = rand.Intn(MaxGeneratedStringLength)
	}
	var b strings.Builder
	b.Grow(length)
	for b.Len() != b.Cap() {
		// Any printable ASCII character.
		r := rand.Intn(95) + ' '
		b.WriteRune(rune(r))
	}
	return b.String()
}

func ArrayDirective(length int) []any {
	if length < 0 {
		length = rand.Intn(MaxGeneratedArrayLength)
	}

	array := make([]any, 0, length)
	for i := 0; i < length; i++ {
		// Either 0 for a primitive, 1 for array, or 2 for object.
		selector := rand.Intn(3)
		switch selector {
		case 1:
			array = append(array, generateInternalArray(-1, false))
		case 2:
			array = append(array, generateInternalObject(-1, false))
		default:
			array = append(array, generatePrimitive())
		}
	}
	return array
}

func ObjectDirective(length int) map[string]any {
	if length < 0 {
		length = rand.Intn(MaxGeneratedObjectLength)
	}

	object := make(map[string]any)
	for i := 0; i < length; i++ {
		key := StringDirective(-1)
		for _, ok := object[key]; ok; _, ok = object[key] {
			// Keep generating until a unique key is created...
		}

		// Either 0 for a primitive, 1 for array, or 2 for object.
		selector := rand.Intn(3)
		switch selector {
		case 1:
			object[key] = generateInternalArray(-1, false)
		case 2:
			object[key] = generateInternalObject(-1, false)
		default:
			object[key] = generatePrimitive()
		}
	}
	return object
}
