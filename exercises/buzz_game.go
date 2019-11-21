package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
		channel1 <- "Player 1 Buzzed"
	}()

	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
		channel2 <- "Player 2 Buzzed"
	}()

	//fmt.Println(<-channel1)
	//fmt.Println(<-channel2)
	select {
	case player1 := <-channel1:
		fmt.Println(player1)
	case player2 := <-channel2:
		fmt.Println(player2)
	}

}
