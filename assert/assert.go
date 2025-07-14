package assert

import (
	"testing"
)

// Equal checks if two values are equal
func Equal[T comparable](t testing.TB, actual, expected T) {
	t.Helper()
	if expected != actual {
		t.Errorf("%s", getErrorMessage(expected, actual))
	}
}

// NotEqual checks if two values are not equal
func NotEqual[T comparable](t testing.TB, actual, unexpected T) {
	t.Helper()
	if unexpected == actual {
		t.Errorf("%s", getErrorMessage(unexpected, actual))
	}
}

// Greater checks if one value is greater than another
func Greater[T ordered](t testing.TB, actual, expected T) {
	t.Helper()
	if actual <= expected {
		t.Errorf("%s", getErrorMessage(expected, actual))
	}
}

// Less checks if one value is less than another
func Less[T ordered](t testing.TB, actual, expected T) {
	t.Helper()
	if actual >= expected {
		t.Errorf("%s", getErrorMessage(expected, actual))
	}
}

// GreaterOrEqual checks if one value is greater than or equal to another
func GreaterOrEqual[T ordered](t testing.TB, actual, expected T) {
	t.Helper()
	if actual < expected {
		t.Errorf("%s", getErrorMessage(expected, actual))
	}
}

// LessOrEqual checks if one value is less than or equal to another
func LessOrEqual[T ordered](t testing.TB, actual, expected T) {
	t.Helper()
	if actual > expected {
		t.Errorf("%s", getErrorMessage(expected, actual))
	}
}

// True checks if a condition is true
func True(t testing.TB, condition bool) {
	t.Helper()
	if !condition {
		t.Errorf("%s", getErrorMessage(true, condition))
	}
}

// False checks if a condition is false
func False(t testing.TB, condition bool) {
	t.Helper()
	if condition {
		t.Errorf("%s", getErrorMessage(false, condition))
	}
}

// Nil checks if a value is nil
func Nil(t testing.TB, value any) {
	t.Helper()
	if value != nil {
		t.Errorf("%s", getErrorMessage(nil, value))
	}
}

// NotNil checks if a value is not nil
func NotNil(t testing.TB, value any) {
	t.Helper()
	if value == nil {
		t.Errorf("%s", getErrorMessage("not nil", "nil"))
	}
}

// Contains checks if a container (string or slice) contains an item
func Contains[T comparable](t testing.TB, container []T, item T) {
	t.Helper()
	if !contains(container, item) {
		t.Errorf("%s", getErrorMessage(item, container))
	}
}

// Length checks if a collection has the expected length
func Length[T container](t testing.TB, collection T, expectedLength int) {
	t.Helper()
	actualLength := getLength(collection)
	if actualLength != expectedLength {
		t.Errorf("%s", getErrorMessage(expectedLength, actualLength))
	}
}

// Zero checks if a numeric value is zero
func Zero[T numeric](t testing.TB, v T) {
	t.Helper()
	if v != 0 {
		t.Errorf("%s", getErrorMessage(0, v))
	}
}

// NotZero checks if a numeric value is not zero
func NotZero[T numeric](t testing.TB, v T) {
	t.Helper()
	if v == 0 {
		t.Errorf("%s", getErrorMessage("not zero", v))
	}
}

// Panic checks if a function call results in a panic
func Panic(t testing.TB, f func()) {
	t.Helper()
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("%s", getErrorMessage("panic", "no panic occurred"))
		}
	}()
	f()
}

// Safe checks if a function call does not result in a panic
func Safe(t testing.TB, f func()) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("%s", getErrorMessage("no panic", "panic occurred"))
		}
	}()
	f()
}

// IsClose checks if a value is close to another value within a tolerance
func IsClose(t testing.TB, actual, expected, tolerance float64) {
	t.Helper()
	if (expected-actual) > tolerance || (actual-expected) > tolerance {
		t.Errorf("%s", getErrorMessage(expected, actual))
	}
}
