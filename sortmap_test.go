package sortmap_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/tg/gosortmap"
)

func ExampleByKey() {
	m := map[string]int{"daikon": 2, "cabbage": 3, "banana": 1, "apple": 4}
	for _, x := range sortmap.ByKey(m) {
		fmt.Printf("%v: %v\n", x.K, x.V)
	}
	// Output:
	// apple: 4
	// banana: 1
	// cabbage: 3
	// daikon: 2
}

func ExampleByKeyDesc() {
	m := map[string]int{"daikon": 2, "cabbage": 3, "banana": 1, "apple": 4}
	for _, x := range sortmap.ByKeyDesc(m) {
		fmt.Printf("%v: %v\n", x.K, x.V)
	}
	// Output:
	// daikon: 2
	// cabbage: 3
	// banana: 1
	// apple: 4
}

func ExampleByVal() {
	m := map[string]int{"daikon": 2, "cabbage": 3, "banana": 1, "apple": 4}
	for _, x := range sortmap.ByVal(m) {
		fmt.Printf("%v: %v\n", x.K, x.V)
	}
	// Output:
	// banana: 1
	// daikon: 2
	// cabbage: 3
	// apple: 4
}

func ExampleByValDesc() {
	m := map[string]int{"daikon": 2, "cabbage": 3, "banana": 1, "apple": 4}
	for _, x := range sortmap.ByValDesc(m) {
		fmt.Printf("%v: %v\n", x.K, x.V)
	}
	// Output:
	// apple: 4
	// cabbage: 3
	// daikon: 2
	// banana: 1
}

var benchMap = func() map[int]int {
	m := make(map[int]int)
	for n := 0; n < 10000; n++ {
		m[rand.Int()] = rand.Int()
	}
	return m
}()

func BenchmarkSortNone(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sortmap.ByFunc(benchMap, func(x, y sortmap.KV) bool { return false })
	}
}

func BenchmarkSortFunc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sortmap.ByFunc(benchMap, func(x, y sortmap.KV) bool { return x.K.(int) < y.K.(int) })
	}
}

func BenchmarkSortKey(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sortmap.ByKey(benchMap)
	}
}
