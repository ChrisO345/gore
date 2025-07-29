package assert

import "testing"

// Assertion is a struct that holds a reference to the testing.T object
type Assertion struct {
	T testing.TB
}

// That creates a new Assertion instance for use in tests.
func That(t testing.TB) *Assertion {
	t.Helper()
	return &Assertion{T: t}
}

// Equal checks if two values are equal
func (a *Assertion) Equal(actual, expected any) {
	a.T.Helper()
	Equal(a.T, actual, expected)
}

// NotEqual checks if two values are not equal
func (a *Assertion) NotEqual(actual, unexpected any) {
	a.T.Helper()
	NotEqual(a.T, actual, unexpected)
}

// Greater checks if one value is greater than another
func (a *Assertion) Greater(actual, expected any) {
	a.T.Helper()
	a.T.Fatalf("Greater is not implemented in fluent style yet, use Greater directly")
}

// Less checks if one value is less than another
func (a *Assertion) Less(actual, expected any) {
	a.T.Helper()
	a.T.Fatalf("Less is not implemented in fluent style yet, use Less directly")
}

// GreaterOrEqual checks if one value is greater than or equal to another
func (a *Assertion) GreaterOrEqual(actual, expected any) {
	a.T.Helper()
	a.T.Fatalf("GreaterOrEqual is not implemented in fluent style yet, use GreaterOrEqual directly")
}

// LessOrEqual checks if one value is less than or equal to another
func (a *Assertion) LessOrEqual(actual, expected any) {
	a.T.Helper()
	a.T.Fatalf("LessOrEqual is not implemented in fluent style yet, use LessOrEqual directly")
}

// True checks if a condition is true
func (a *Assertion) True(condition bool) {
	a.T.Helper()
	True(a.T, condition)
}

// False checks if a condition is false
func (a *Assertion) False(condition bool) {
	a.T.Helper()
	False(a.T, condition)
}

// IsClose checks if two floating-point numbers are close to each other
func (a *Assertion) IsClose(actual, expected any, tolerance float64) {
	a.T.Helper()
	a.T.Fatalf("IsClose is not implemented in fluent style yet, use IsClose directly")
}

// Zero checks if a numeric value is zero
func (a *Assertion) Zero(value any) {
	a.T.Helper()
	a.T.Fatalf("Zero is not implemented in fluent style yet, use Zero directly")
}

// NotZero checks if a numeric value is not zero
func (a *Assertion) NotZero(value any) {
	a.T.Helper()
	a.T.Fatalf("NotZero is not implemented in fluent style yet, use NotZero directly")
}

// === ERRORS ===

// Nil checks if a value is nil
func (a *Assertion) Nil(value any) {
	a.T.Helper()
	Nil(a.T, value)
}

// NotNil checks if a value is not nil
func (a *Assertion) NotNil(value any) {
	a.T.Helper()
	NotNil(a.T, value)
}

// Panic checks if a function panics
func (a *Assertion) Panic(fn func()) {
	a.T.Helper()
	Panic(a.T, fn)
}

// Safe checks if a function does not panic
func (a *Assertion) Safe(fn func()) {
	a.T.Helper()
	Safe(a.T, fn)
}

// === CONTAINERS ===

// Contains checks if a slice contains a specific item
func (a *Assertion) Contains(slice []any, item any) {
	a.T.Helper()
	Contains(a.T, slice, item)
}

// NotContains checks if a slice does not contain a specific item
func (a *Assertion) NotContains(slice []any, item any) {
	a.T.Helper()
	NotContains(a.T, slice, item)
}

// Length checks the length of a collection
func (a *Assertion) Length(collection []any, expectedLength int) {
	a.T.Helper()
	Length(a.T, collection, expectedLength)
}

// === STRINGS ===

// StringContains checks if a string contains a substring
func (a *Assertion) StringContains(s, substring string) {
	a.T.Helper()
	if !containsString(s, substring) {
		a.T.Errorf("%s", getErrorMessage(substring, s))
	}
}

// NotStringContains checks if a string does not contain a substring
func (a *Assertion) NotStringContains(s, substring string) {
	a.T.Helper()
	if containsString(s, substring) {
		a.T.Errorf("%s", getErrorMessage("not contains", s))
	}
}

// PrefixString checks if a string starts with a specific prefix
func (a *Assertion) PrefixString(s, prefix string) {
	a.T.Helper()
	if !hasPrefix(s, prefix) {
		a.T.Errorf("%s", getErrorMessage(prefix, s))
	}
}

// NotPrefixString checks if a string does not start with a specific prefix
func (a *Assertion) NotPrefixString(s, prefix string) {
	a.T.Helper()
	if hasPrefix(s, prefix) {
		a.T.Errorf("%s", getErrorMessage("not starts with", s))
	}
}

// SuffixString checks if a string ends with a specific suffix
func (a *Assertion) SuffixString(s, suffix string) {
	a.T.Helper()
	if !hasSuffix(s, suffix) {
		a.T.Errorf("%s", getErrorMessage(suffix, s))
	}
}

// NotSuffixString checks if a string does not end with a specific suffix
func (a *Assertion) NotSuffixString(s, suffix string) {
	a.T.Helper()
	if hasSuffix(s, suffix) {
		a.T.Errorf("%s", getErrorMessage("not ends with", s))
	}
}
