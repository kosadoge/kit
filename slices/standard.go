package slices

import (
	"cmp"
	"slices"
)

// BinarySearch searches for target in a sorted slice and returns the position
// where target is found, or the position where target would appear in the
// sort order; it also returns a bool saying whether the target is really found
// in the slice. The slice must be sorted in increasing order.
func BinarySearch[S ~[]E, E cmp.Ordered](s S, target E) (int, bool) {
	return slices.BinarySearch(s, target)
}

// BinarySearchFunc works like [BinarySearch], but uses a custom comparison
// function. The slice must be sorted in increasing order, where "increasing"
// is defined by cmp. cmp should return 0 if the slice element matches
// the target, a negative number if the slice element precedes the target,
// or a positive number if the slice element follows the target.
// cmp must implement the same ordering as the slice, such that if
// cmp(a, t) < 0 and cmp(b, t) >= 0, then a must precede b in the slice.
func BinarySearchFunc[S ~[]E, E, T any](s S, target T, cmp func(E, T) int) (int, bool) {
	return slices.BinarySearchFunc(s, target, cmp)
}

// Clip removes unused capacity from the slice, returning s[:len(s):len(s)].
func Clip[S ~[]E, E any](s S) S {
	return slices.Clip(s)
}

// Clone returns a copy of the slice.
// The elements are copied using assignment, so this is a shallow clone.
func Clone[S ~[]E, E any](s S) S {
	return slices.Clone(s)
}

// Compact replaces consecutive runs of equal elements with a single copy.
// This is like the uniq command found on Unix.
// Compact modifies the contents of the slice s and returns the modified slice,
// which may have a smaller length.
// Compact zeroes the elements between the new length and the original length.
func Compact[S ~[]E, E comparable](s S) S {
	return slices.Compact(s)
}

// CompactFunc is like [Compact] but uses an equality function to compare elements.
// For runs of elements that compare equal, CompactFunc keeps the first one.
// CompactFunc zeroes the elements between the new length and the original length.
func CompactFunc[S ~[]E, E any](s S, eq func(E, E) bool) S {
	return slices.CompactFunc(s, eq)
}

// Compare compares the elements of s1 and s2, using [cmp.Compare] on each pair
// of elements. The elements are compared sequentially, starting at index 0,
// until one element is not equal to the other.
// The result of comparing the first non-matching elements is returned.
// If both slices are equal until one of them ends, the shorter slice is
// considered less than the longer one.
// The result is 0 if s1 == s2, -1 if s1 < s2, and +1 if s1 > s2.
func Compare[S ~[]E, E cmp.Ordered](s1, s2 S) int {
	return slices.Compare(s1, s2)
}

// CompareFunc is like [Compare] but uses a custom comparison function on each
// pair of elements.
// The result is the first non-zero result of cmp; if cmp always
// returns 0 the result is 0 if len(s1) == len(s2), -1 if len(s1) < len(s2),
// and +1 if len(s1) > len(s2).
func CompareFunc[S ~[]E, E any](s1, s2 S, cmp func(E, E) int) int {
	return slices.CompareFunc(s1, s2, cmp)
}

// Concat returns a new slice concatenating the passed in slices.
//
// added in go 1.22.0
func Concat[S ~[]E, E any](slices ...S) S {
	size := 0
	for _, s := range slices {
		size += len(s)
		if size < 0 {
			panic("len out of range")
		}
	}
	newslice := Grow[S](nil, size)
	for _, s := range slices {
		newslice = append(newslice, s...)
	}
	return newslice
}

// Contains reports whether v is present in s.
func Contains[S ~[]E, E comparable](s S, v E) bool {
	return slices.Contains(s, v)
}

// ContainsFunc reports whether at least one
// element e of s satisfies f(e).
func ContainsFunc[S ~[]E, E any](s S, f func(E) bool) bool {
	return slices.ContainsFunc(s, f)
}

// Delete removes the elements s[i:j] from s, returning the modified slice.
// Delete panics if j > len(s) or s[i:j] is not a valid slice of s.
// Delete is O(len(s)-i), so if many items must be deleted, it is better to
// make a single call deleting them all together than to delete one at a time.
// Delete zeroes the elements s[len(s)-(j-i):len(s)].
func Delete[S ~[]E, E any](s S, i, j int) S {
	return slices.Delete(s, i, j)
}

// DeleteFunc removes any elements from s for which del returns true,
// returning the modified slice.
// DeleteFunc zeroes the elements between the new length and the original length.
func DeleteFunc[S ~[]E, E any](s S, del func(E) bool) S {
	return slices.DeleteFunc(s, del)
}

// Equal reports whether two slices are equal: the same length and all
// elements equal. If the lengths are different, Equal returns false.
// Otherwise, the elements are compared in increasing index order, and the
// comparison stops at the first unequal pair.
// Floating point NaNs are not considered equal.
func Equal[S ~[]E, E comparable](s1, s2 S) bool {
	return slices.Equal(s1, s2)
}

// EqualFunc reports whether two slices are equal using an equality
// function on each pair of elements. If the lengths are different,
// EqualFunc returns false. Otherwise, the elements are compared in
// increasing index order, and the comparison stops at the first index
// for which eq returns false.
func EqualFunc[S ~[]E, E any](s1, s2 S, eq func(E, E) bool) bool {
	return slices.EqualFunc(s1, s2, eq)
}

// Grow increases the slice's capacity, if necessary, to guarantee space for
// another n elements. After Grow(n), at least n elements can be appended
// to the slice without another allocation. If n is negative or too large to
// allocate the memory, Grow panics.
func Grow[S ~[]E, E any](s S, n int) S {
	return slices.Grow(s, n)
}

// Index returns the index of the first occurrence of v in s,
// or -1 if not present.
func Index[S ~[]E, E comparable](s S, v E) int {
	return slices.Index(s, v)
}

// IndexFunc returns the first index i satisfying f(s[i]),
// or -1 if none do.
func IndexFunc[S ~[]E, E any](s S, f func(E) bool) int {
	return slices.IndexFunc(s, f)
}

// Insert inserts the values v... into s at index i,
// returning the modified slice.
// The elements at s[i:] are shifted up to make room.
// In the returned slice r, r[i] == v[0],
// and r[i+len(v)] == value originally at r[i].
// Insert panics if i is out of range.
// This function is O(len(s) + len(v)).
func Insert[S ~[]E, E any](s S, i int, v ...E) S {
	return slices.Insert(s, i, v...)
}

// IsSorted reports whether s is sorted in ascending order.
func IsSorted[S ~[]E, E cmp.Ordered](s S) bool {
	return slices.IsSorted(s)
}

// IsSortedFunc reports whether x is sorted in ascending order, with cmp as the
// comparison function as defined by [SortFunc].
func IsSortedFunc[S ~[]E, E any](s S, cmp func(E, E) int) bool {
	return slices.IsSortedFunc(s, cmp)
}

// Max returns the maximal value in s. It panics if s is empty.
// For floating-point E, Max propagates NaNs (any NaN value in s
// forces the output to be NaN).
func Max[S ~[]E, E cmp.Ordered](s S) E {
	return slices.Max(s)
}

// MaxFunc returns the maximal value in s, using cmp to compare elements.
// It panics if s is empty. If there is more than one maximal element
// according to the cmp function, MaxFunc returns the first one.
func MaxFunc[S ~[]E, E cmp.Ordered](s S, cmp func(E, E) int) E {
	return slices.MaxFunc(s, cmp)
}

// Min returns the minimal value in s. It panics if s is empty.
// For floating-point numbers, Min propagates NaNs (any NaN value in s
// forces the output to be NaN).
func Min[S ~[]E, E cmp.Ordered](s S) E {
	return slices.Min(s)
}

// MinFunc returns the minimal value in s, using cmp to compare elements.
// It panics if s is empty. If there is more than one minimal element
// according to the cmp function, MinFunc returns the first one.
func MinFunc[S ~[]E, E cmp.Ordered](s S, cmp func(E, E) int) E {
	return slices.MinFunc(s, cmp)
}

// Replace replaces the elements s[i:j] by the given v, and returns the
// modified slice.
// Replace panics if j > len(s) or s[i:j] is not a valid slice of s.
// When len(v) < (j-i), Replace zeroes the elements between the new length and the original length.
func Replace[S ~[]E, E any](s S, i, j int, v ...E) S {
	return slices.Replace(s, i, j, v...)
}

// Reverse reverses the elements of the slice in place.
func Reverse[S ~[]E, E any](s S) {
	slices.Reverse(s)
}

// Sort sorts a slice of any ordered type in ascending order.
// When sorting floating-point numbers, NaNs are ordered before other values.
func Sort[S ~[]E, E cmp.Ordered](s S) {
	slices.Sort(s)
}

// SortFunc sorts the slice s in ascending order as determined by the cmp
// function. This sort is not guaranteed to be stable.
// cmp(a, b) should return a negative number when a < b, a positive number when
// a > b and zero when a == b.
//
// SortFunc requires that cmp is a strict weak ordering.
// See https://en.wikipedia.org/wiki/Weak_ordering#Strict_weak_orderings.
func SortFunc[S ~[]E, E any](s S, cmp func(E, E) int) {
	slices.SortFunc(s, cmp)
}

// SortStableFunc sorts the slice s while keeping the original order of equal
// elements, using cmp to compare elements in the same way as [SortFunc].
func SortStableFunc[S ~[]E, E any](s S, cmp func(E, E) int) {
	slices.SortStableFunc(s, cmp)
}
