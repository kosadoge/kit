# kit
[![Go Reference](https://pkg.go.dev/badge/github.com/kosadoge/kit.svg)](https://pkg.go.dev/github.com/kosadoge/kit)

A collection of utility functions that are frequently needed during development.

## Install
```
$ go get github.com/kosadoge/kit@latest
```

## Usage
`kit` is organized into different packages for various purposes.

### collect
A package offering functional-style operations on slices. It includes the following functions:

- **Map**: Transforms elements in a slice according to the provided function `f`.
- **Keep**: Retains elements in a slice where the function `f` returns true.
- **Discard**: Removes elements in a slice where the function `f` returns true.
- **Unique**: Removes duplicate elements in a slice.
