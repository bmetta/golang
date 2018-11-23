package main

import "fmt"

func main() {
	// create an array a that will hold exactly 5 ints
	var a [5]int
	fmt.Println(a)

	// declare and initialize an array in one line.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// multi-dimensional array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

	// declare and initialize 2D array in one line.
	twoD1 := [2][3]int{
		{1, 2, 3},
		{3, 4, 5}}
	fmt.Println(twoD1)
}
