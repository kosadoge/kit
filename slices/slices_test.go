package slices_test

import (
	"strconv"
	"testing"

	"github.com/kosadoge/kit/internal/assert"
	"github.com/kosadoge/kit/slices"
)

func TestMap(t *testing.T) {
	t.Parallel()

	args := []int{1, 2, 3, 4, 5}

	expect := []string{"1", "2", "3", "4", "5"}
	actual := slices.Map(args, func(n int) string { return strconv.FormatInt(int64(n), 10) })

	assert.Equal(t, expect, actual)
}

func TestKeep(t *testing.T) {
	t.Parallel()

	args := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	expect := []int{2, 4, 6, 8, 10}
	actual := slices.Keep(args, func(n int) bool { return n%2 == 0 })

	assert.Equal(t, expect, actual)
}

func TestDiscard(t *testing.T) {
	t.Parallel()

	args := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	expect := []int{1, 3, 5, 7, 9}
	actual := slices.Discard(args, func(n int) bool { return n%2 == 0 })

	assert.Equal(t, expect, actual)
}

func TestUnique(t *testing.T) {
	t.Parallel()

	args := []string{"🍎", "🍌", "🍉", "🍏", "🍎", "🍎", "🍇", "🍉", "🍌"}

	expect := []string{"🍎", "🍌", "🍉", "🍏", "🍇"}
	actual := slices.Unique(args)

	assert.Equal(t, expect, actual)
}

func TestChunk(t *testing.T) {
	t.Parallel()

	args := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	expect := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9}}
	actual := slices.Chunk(args, 2)

	assert.Equal(t, expect, actual)
}

func TestToMap(t *testing.T) {
	t.Parallel()

	type User struct {
		ID   int
		Name string
	}

	args := []User{
		{ID: 1, Name: "alice"},
		{ID: 2, Name: "bob"},
		{ID: 3, Name: "charlie"},
		{ID: 4, Name: "david"},
	}

	expect := map[int]string{
		1: "alice",
		2: "bob",
		3: "charlie",
		4: "david",
	}
	actual := slices.ToMap(args, func(u User) (int, string) { return u.ID, u.Name })

	assert.Equal(t, expect, actual)
}
