package main

import "fmt"

func main() {
	var s []int
	fmt.Printf("%v\n", s)

	s = append(s, 1)
	fmt.Printf("%v\n", s)

	s = append(s, 2, 3, 4)
	fmt.Printf("%v\n", s)
}
