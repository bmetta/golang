package main

import "fmt"

func main() {
	// don’t need parentheses around conditions in Go,
	// but that the braces are required.
	// basic example.
	i := 1
	if i < 1 {
		fmt.Println(i, "< 1")
	} else {
		fmt.Println(i, ">= 1")
	}

	// A statement can precede conditionals;
	// any variables declared in this statement
	// are available in all branches.
	if n := 9; n < 0 {
		fmt.Println(n, "is a negative number")
	} else if n < 10 {
		fmt.Println(n, "is a one digit number")
	} else {
		fmt.Println(n, "is a multi digit number")
	}
}
