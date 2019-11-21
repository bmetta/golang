package main

import "fmt"

func main() {
	// The following code causes dead lock. Sending and Receiving operations are blocking the code
	//int_chan := make(chan int)
	//int_chan <- 10
	//fmt.Println(<-int_chan)

	// Solution for above problem: wrap the one of them in go routine
	//int_chan := make(chan int)
	//go func() {
	//	int_chan <- 10
	//}()
	//fmt.Println(<-int_chan)

	// Buffered channel will solve this problem
	int_chan := make(chan int, 2)
	int_chan <- 10
	int_chan <- 20
	//int_chan <- 30 // makes the dead lock
	fmt.Println(<-int_chan)
}
