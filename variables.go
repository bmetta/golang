package main

import "fmt"

func main() {
	// variable declarations
	var i int
	var u, v, w float64
	var k = 0
	var x, y float32 = -1, -2
	fmt.Println(i, u, v, w, k, x, y)

	var (
		i1      int
		U, V, S = 2.0, 3.0, "bar"
	)
	fmt.Println(i1, U, V, S)

	// Short variable declarations
	j, k := 0, 10
	f := func() int { return 7 }
	fmt.Println(j, k, f())
}
