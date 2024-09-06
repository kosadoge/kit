# kit
[![Go Reference](https://pkg.go.dev/badge/github.com/kosadoge/kit.svg)](https://pkg.go.dev/github.com/kosadoge/kit)

A collection of utility functions that are frequently needed during development.

## Install
```
$ go get github.com/kosadoge/kit@latest
```

## Usage
`kit` is organized into different packages for various purposes.

### slices
A replacement for the standard library `slices` package, with additional functions:

- Map: Transforms elements in a slice according to the provided function `f`
- Keep: Keeps elements in a slice where the function `f` returns true
- Discard: Removes elements in a slice where the function `f` returns true
- Unique: Removes duplicate elements in a slice
- ToMap: Converts a slice to a map using the provided function `f`

> Note: This package does not support Go 1.23's new functions: All, AppendSeq, Backward, Collect, Sorted, SortedFunc, SortedStableFunc, and Values.

### maps
A replacement for the standard library `maps` package, with additional functions:

- Keys: Extracts keys from a map and returns them as a slice (Note: differs from the Go 1.23 implementation.)
- Values: Extracts values from a map and returns them as a slice (Note: differs from the Go 1.23 implementation.)
- ToSlice: Converts a map to a slice using the provided function `f`

> Note: This package does not support Go 1.23's new functions: All, Collect, and Insert.