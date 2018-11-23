package main

import "fmt"
import "time"

func main() {
	// basic switch.
	// eg: 1 => one, 2 => two

	// i := 2
	switch i := 2; i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("greater than three")
	}

	// commas to separate multiple expressions in the same case statement.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	case time.Friday:
		fmt.Println("Before Weekend")
	default:
		fmt.Println("Weekday")
	}

	// switch without an expression is an alternate way to
	// express if/else logic.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Its before noon")
	default:
		fmt.Println("Its after noon")
	}

	// A type switch compares types instead of values.
	// You can use this to discover the type of an interface value.
	// In this example, the variable t will have the type
	// corresponding to its clause.

}
