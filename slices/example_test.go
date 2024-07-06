package slices_test

import (
	"fmt"
	"strconv"

	"github.com/kosadoge/kit/slices"
)

func ExampleMap() {
	nums := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	strs := slices.Map(nums, func(n float64) string { return strconv.FormatFloat(n, 'f', -1, 64) })

	fmt.Printf("%#v\n", strs)
	// Output:
	// []string{"1.1", "2.2", "3.3", "4.4", "5.5"}
}

func ExampleMap_reuseFunc() {
	nums := []int{1, 2, 3, 4, 5}
	strs := slices.Map(nums, strconv.Itoa)

	fmt.Printf("%#v\n", strs)
	// Output:
	// []string{"1", "2", "3", "4", "5"}
}

func ExampleMap_extractStructFields() {
	type Box struct {
		ID    int64
		Value string
	}

	boxes := []Box{
		{ID: 1, Value: "apple"},
		{ID: 2, Value: "banana"},
		{ID: 3, Value: "cat"},
		{ID: 4, Value: "dog"},
	}

	ids := slices.Map(boxes, func(b Box) int64 { return b.ID })
	values := slices.Map(boxes, func(b Box) string { return b.Value })

	fmt.Printf("%#v\n", ids)
	fmt.Printf("%#v\n", values)
	// Output:
	// []int64{1, 2, 3, 4}
	// []string{"apple", "banana", "cat", "dog"}
}

func ExampleKeep() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	even := slices.Keep(nums, func(n int) bool { return n%2 == 0 })

	fmt.Printf("%#v\n", even)
	// Output:
	// []int{2, 4, 6, 8, 10}
}

func ExampleDiscard() {
	nums := []int{-5, -3, -1, 1, 3, 5}
	positive := slices.Discard(nums, func(n int) bool { return n < 0 })

	fmt.Printf("%#v\n", positive)
	// Output:
	// []int{1, 3, 5}
}
