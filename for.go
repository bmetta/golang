package main

import "fmt"

func main() {
	// for is Go’s only looping construct.
	// Here are three basic types of for loops.

	// The most basic type, with a single condition.
	for 0 < 1 { // for ; cond ;
		fmt.Println("true")
		break
	}

	// A classic initial/condition/after for loop.
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// for without a condition will loop repeatedly
	for { // for true
		fmt.Println("true")
		break
	}

	x := []int{3, 4, 5, 6}
	for i, val := range x {
		fmt.Println(i, val)
	}
}
