package main

import "time"
import "fmt"

func f1(c1 chan string) {
	for {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}
}

func f2(c2 chan string) {
	for {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go f1(c1)
	go f2(c2)

	//for i := 0; i < 2; i++ {
	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
