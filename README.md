# kit
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

### maps
A replacement for the standard library `maps` package, with additional functions:

- Keys: Extracts keys from a map and returns them as a slice
- Values: Extracts values from a map and returns them as a slice