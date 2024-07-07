package maps_test

import (
	"sort"
	"testing"

	"github.com/kosadoge/kit/internal/assert"
	"github.com/kosadoge/kit/maps"
)

func TestKeys(t *testing.T) {
	t.Parallel()

	args := map[string]string{
		"a": "apple",
		"b": "banana",
		"c": "cat",
		"d": "dog",
		"e": "egg",
	}

	expect := []string{"a", "b", "c", "d", "e"}
	actual := maps.Keys(args)

	// Since Go maps do not guarantee order, we sort the result for comparison.
	sort.Strings(actual)

	assert.Equal(t, expect, actual)
}

func TestValues(t *testing.T) {
	t.Parallel()

	args := map[string]string{
		"a": "apple",
		"b": "banana",
		"c": "cat",
		"d": "dog",
		"e": "egg",
	}

	expect := []string{"apple", "banana", "cat", "dog", "egg"}
	actual := maps.Values(args)

	// Since Go maps do not guarantee order, we sort the result for comparison.
	sort.Strings(actual)

	assert.Equal(t, expect, actual)
}
