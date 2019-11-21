package main

import "fmt"

func main() {
	s := make([]int, 0, 10)
	fmt.Printf("%v\n", s)

	//s = append(s, 1)
	//fmt.Printf("%v\n", s)

	//s[0] = 2
	//fmt.Printf("%v\n", s)

	//s = append(s, 3, 4)
	//fmt.Printf("%v\n", s)

	s[1] = 2
	s[2] = 3
	fmt.Printf("%d\n", s[2])
	fmt.Printf("%v\n", s)
}
