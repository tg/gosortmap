package sortmap

import "fmt"

func ExampleSortValInt() {
	m := map[string]int{"b": 2, "c": 3, "a": 1, "d": 4}
	for _, x := range SortValInt(m) {
		fmt.Printf("%v: %v\n", x.K, x.V)
	}
	// Output:
	// a: 1
	// b: 2
	// c: 3
	// d: 4
}

func ExampleSortValIntDesc() {
	m := map[string]int{"b": 2, "c": 3, "a": 1, "d": 4}
	for _, x := range SortValIntDesc(m) {
		fmt.Printf("%v: %v\n", x.K, x.V)
	}
	// Output:
	// d: 4
	// c: 3
	// b: 2
	// a: 1
}
