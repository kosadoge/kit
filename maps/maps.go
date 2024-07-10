package maps

// Keys returns a slice of all keys in the input map.
// If the input map is nil, it returns nil.
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	if m == nil {
		return nil
	}

	result := make([]K, 0, len(m))
	for k := range m {
		result = append(result, k)
	}

	return result
}

// Values returns a slice of all values in the input map.
// If the input map is nil, it returns nil.
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	if m == nil {
		return nil
	}

	result := make([]V, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}

	return result
}

// ToSlice converts a map to a slice using a provided function.
// The function f takes a key-value pair from the map and returns an element for the slice.
// If the input map is nil, it returns nil.
func ToSlice[S []E, M ~map[K]V, K comparable, V, E any](m M, f func(K, V) E) S {
	if m == nil {
		return nil
	}

	result := make([]E, 0, len(m))
	for k, v := range m {
		result = append(result, f(k, v))
	}

	return result
}
