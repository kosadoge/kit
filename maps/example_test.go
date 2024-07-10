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

	// Since Go maps do not guarantee order, we sort the result for display.
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

	// Since Go maps do not guarantee order, we sort the result for display.
	sort.Strings(emails)

	fmt.Printf("%#v\n", emails)
	// Output:
	// []string{"alice@mail.com", "bob@mail.com", "charlie@mail.com", "david@mail.com"}
}

func ExampleToSlice() {
	type User struct {
		ID   int
		Name string
	}

	users := map[int]User{
		1: {ID: 1, Name: "alice"},
		2: {ID: 2, Name: "bob"},
		3: {ID: 3, Name: "charlie"},
		4: {ID: 4, Name: "david"},
	}
	names := maps.ToSlice(users, func(_ int, u User) string { return u.Name })

	// Since Go maps do not guarantee order, we sort the result for display.
	sort.Strings(names)

	fmt.Printf("%#v\n", names)
	// Output:
	// []string{"alice", "bob", "charlie", "david"}
}
