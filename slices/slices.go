package slices

// Map applies a function f to each element of a slice s and returns a new slice of the results.
// If the input slice is nil, it returns nil.
func Map[S ~[]E, E, T any](s S, f func(E) T) []T {
	if s == nil {
		return nil
	}

	result := make([]T, len(s))
	for i := range s {
		result[i] = f(s[i])
	}

	return result
}

// Keep retains elements in the slice s that satisfy the predicate f.
// If the input slice is nil, it returns nil.
func Keep[S ~[]E, E any](s S, f func(E) bool) []E {
	if s == nil {
		return nil
	}

	result := make([]E, 0, len(s))
	for i := range s {
		if f(s[i]) {
			result = append(result, s[i])
		}
	}

	return result
}

// Discard removes elements in the slice s that satisfy the predicate f.
// If the input slice is nil, it returns nil.
func Discard[S ~[]E, E any](s S, f func(E) bool) []E {
	if s == nil {
		return nil
	}

	result := make([]E, 0, len(s))
	for i := range s {
		if !f(s[i]) {
			result = append(result, s[i])
		}
	}

	return result
}

// Unique returns a new slice with duplicate elements removed.
// If the input slice is nil, it returns nil.
func Unique[S ~[]E, E comparable](s S) []E {
	if s == nil {
		return nil
	}

	seen := make(map[E]struct{}, len(s))
	result := make([]E, 0, len(s))

	for i := range s {
		if _, exists := seen[s[i]]; !exists {
			seen[s[i]] = struct{}{}
			result = append(result, s[i])
		}

	}

	return result
}
