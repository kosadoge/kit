package maps

import "maps"

// Clone returns a copy of m.  This is a shallow clone:
// the new keys and values are set using ordinary assignment.
func Clone[M ~map[K]V, K comparable, V any](m M) M {
	return maps.Clone(m)
}

// Copy copies all key/value pairs in src adding them to dst.
// When a key in src is already present in dst,
// the value in dst will be overwritten by the value associated
// with the key in src.
func Copy[M ~map[K]V, K comparable, V any](dst, src M) {
	maps.Copy(dst, src)
}

// DeleteFunc deletes any key/value pairs from m for which del returns true.
func DeleteFunc[M ~map[K]V, K comparable, V any](m M, del func(K, V) bool) {
	maps.DeleteFunc(m, del)
}

// Equal reports whether two maps contain the same key/value pairs.
// Values are compared using ==.
func Equal[M ~map[K]V, K, V comparable](m1, m2 M) bool {
	return maps.Equal(m1, m2)
}

// EqualFunc is like Equal, but compares values using eq.
// Keys are still compared with ==.
func EqualFunc[M ~map[K]V, K comparable, V any](m1, m2 M, eq func(V, V) bool) bool {
	return maps.EqualFunc(m1, m2, eq)
}
