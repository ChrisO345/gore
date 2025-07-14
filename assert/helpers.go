package assert

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

// getLength returns the length of a collection
func getLength[T container](collection T) int {
	return len(collection)
}

// isContainer checks if a type is a container (slice, array, map, etc.)
func isContainer[T any](collection T) bool {
	_, ok := any(collection).([]any)
	return ok
}
