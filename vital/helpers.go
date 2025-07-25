// Code generated by go generate; DO NOT EDIT.

package vital

import "slices"

import "fmt"

// getErrorMessage formats the error message for assertion failures
func getErrorMessage(expected, actual any) string {
	return fmt.Sprintf(
		"\nAssertion failed:\nExpected:\n\t%v\nGot:\n\t%v\n",
		expected, actual,
	)
}

// contains checks if a slice contains a specific item
func contains[T comparable](slice []T, item T) bool {
	return slices.Contains(slice, item)
}

// stringContains checks if a string contains a substring
func containsString(s, substring string) bool {
	return slices.Contains([]rune(s), rune(substring[0]))
}

// getLength returns the length of a collection
func getLength[T container](collection T) int {
	return len(collection)
}

type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

type container interface {
	~string | ~[]byte | ~[]rune | ~[]any |
		~[]int | ~[]int8 | ~[]int16 | ~[]int64 |
		~[]uint | ~[]uint16 | ~[]uint32 | ~[]uint64 | ~[]uintptr
}

type numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

