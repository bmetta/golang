package main

import "fmt"

/*
 * 	create a channel
 *		channel = make(chan type)
 *
 *	send a value into a channel
 *		channel <-
 *
 *	receive a value from a channel
 *		<- channel
 */

func main() {
	message := make(chan string)

	// anonymous go routine
	go func() {
		message <- "ping"
	}()

	msg := <-message
	fmt.Println(msg)
}
