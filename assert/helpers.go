package assert

import (
	"fmt"
	"slices"
	"strings"
)

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

// containsString checks if a string contains a substring.
func containsString(s, substring string) bool {
	return strings.Contains(s, substring)
}

// hasPrefix checks if a string starts with a specific prefix.
func hasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// hasSuffix checks if a string ends with a specific suffix.
func hasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
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
