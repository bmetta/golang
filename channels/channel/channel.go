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
	go sendValues(int_chan)

	//for i := 0; i < 6; i++ { // this will make dead lock
	for i := 0; i < 5; i++ {
		fmt.Println(<-int_chan) // receiving value
	}
}
