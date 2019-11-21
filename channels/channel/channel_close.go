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
	close(int_chan) // closing the channel after all our send operations
}

func main() {
	int_chan := make(chan int)
	go sendValues(int_chan)

	//for i := 0; i < 5; i++ {
	//	fmt.Println(<-int_chan) // receiving value
	//}

	// receive operation returns 0 without blocking on the 6th iteration upon closing the channel by sender
	//for i := 0; i < 6; i++ {
	//	fmt.Println(<-int_chan) // receiving value
	//}

	// To check if the channel is open
	//for i := 0; i < 6; i++ {
	//	value, open := <-int_chan
	//	if !open {
	//		break
	//	}
	//	fmt.Println(value)
	//}

	// another way to implement is using range
	for value := range int_chan {
		fmt.Println(value)
	}
}
