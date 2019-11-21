package main

import "fmt"

/*
 *create a channel
 *	channel = make(chan type)
 *
 *send a value into a channel
 *	channel <-
 *
 *receive a value from a channel
 *	<- channel
 */
func sendValues(int_chan chan int) {
	for i := 0; i < 5; i++ {
		int_chan <- i // sending value
	}
}

func main() {
	int_chan := make(chan int)
	// it defers the execution of a function until the end of the surrounding function
	// Good practice to defer the closing of channels in the main program
	defer close(int_chan)

	go sendValues(int_chan)

	//for i := 0; i < 6; i++ { // this will make dead lock
	for i := 0; i < 5; i++ {
		fmt.Println(<-int_chan) // receiving value
	}
}
