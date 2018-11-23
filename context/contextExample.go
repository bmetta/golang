package main

import "fmt"
import "context"

func main() {
	rootCtx := context.Background() // Returns a singleton
	// a is a child of rootCtx
	a := context.WithValue(rootCtx, "location", "San Francisco")
	// b is a child of a
	b := context.WithValue(a, "weather", "foggy")
	// c is also a child of a and is unrelated to b
	c := context.WithValue(a, "weather", "cloudy")
	// d is a child of c and overrides c's "weather" value
	d := context.WithValue(c, "weather", "sunny")

	fmt.Printf("location=%v weather=%v\n", a.Value("location"), a.Value("weather"))
	fmt.Printf("location=%v weather=%v\n", b.Value("location"), b.Value("weather"))
	fmt.Printf("location=%v weather=%v\n", c.Value("location"), c.Value("weather"))
	fmt.Printf("location=%v weather=%v\n", d.Value("location"), d.Value("weather"))

	a.String()

}
