package maps_test

import (
	"fmt"
	"sort"

	"github.com/kosadoge/kit/maps"
)

func ExampleKeys() {
	users := map[string]string{
		"alice":   "alice@mail.com",
		"bob":     "bob@mail.com",
		"charlie": "charlie@mail.com",
		"david":   "david@mail.com",
	}
	names := maps.Keys(users)

	// Since Go maps do not guarantee order, we sort the result for comparison.
	sort.Strings(names)

	fmt.Printf("%#v\n", names)
	// Output:
	// []string{"alice", "bob", "charlie", "david"}
}

func ExampleValues() {
	users := map[string]string{
		"alice":   "alice@mail.com",
		"bob":     "bob@mail.com",
		"charlie": "charlie@mail.com",
		"david":   "david@mail.com",
	}
	emails := maps.Values(users)

	// Since Go maps do not guarantee order, we sort the result for comparison.
	sort.Strings(emails)

	fmt.Printf("%#v\n", emails)
	// Output:
	// []string{"alice@mail.com", "bob@mail.com", "charlie@mail.com", "david@mail.com"}
}
