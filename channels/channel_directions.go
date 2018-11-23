package main

import "fmt"

func ping(ping_ch chan<- string, msg string) {
	ping_ch <- msg
}

func pong(pong_ch chan<- string, ping_ch <-chan string) {
	msg := <-ping_ch
	pong_ch <- msg
}

func main() {
	ping_ch := make(chan string, 1)
	pong_ch := make(chan string, 1)

	ping(ping_ch, "pinging...")
	pong(pong_ch, ping_ch)
	fmt.Println(<-pong_ch)
}
