# gosortmap [![GoDoc](https://godoc.org/github.com/tg/gosortmap?status.svg)](https://godoc.org/github.com/tg/gosortmap)
Get your maps sorted by keys, values or a custom comparator.
## Example
```go
m := map[string]int{"daikon": 2, "cabbage": 3, "banana": 1, "apple": 4}
for _, x := range sortmap.ByValDesc(m) {
	fmt.Printf("%v: %v\n", x.K, x.V)
}
// Output:
// apple: 4
// cabbage: 3
// daikon: 2
// banana: 1
```
