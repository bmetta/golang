package main

import "fmt"
import "time"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(1)
	}
}

func main() {
	f("main")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going on")

	fmt.Scanln()
	fmt.Println("done")
}
