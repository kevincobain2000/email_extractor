package pkg

// IsEqualSlice checks if two slices of strings are equal.
func IsEqualSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// UniqueStrings returns a new slice with unique strings from the given slice.
func UniqueStrings(input []string) []string {
	seen := make(map[string]struct{}) // Using an empty struct{} to minimize memory usage
	var result []string

	for _, value := range input {
		if _, exists := seen[value]; !exists {
			seen[value] = struct{}{}
			result = append(result, value)
		}
	}

	return result
}

// check if string is in slice
func StringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
