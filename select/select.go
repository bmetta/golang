package main

import "time"
import "fmt"

/*
 * select statement blocks the code and waits for
 * multiple channel operations simultaneously.
 *
 *
 */

func f1(channel1 chan string) {
	for {
		time.Sleep(1 * time.Second)
		channel1 <- "I will ping every 1 sec"
	}
}

func f2(channel2 chan string) {
	for {
		time.Sleep(2 * time.Second)
		channel2 <- "I will ping every 2 sec"
	}
}

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go f1(channel1)
	go f2(channel2)

	for {
		select {
		case msg1 := <-channel1:
			fmt.Println(msg1)
		case msg2 := <-channel2:
			fmt.Println(msg2)
		}
	}
}
