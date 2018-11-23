package main

import (
	"context"
	"fmt"
	//"log"
	//"net/http"
	//_ "net/http/pprof"
	"runtime"
	"time"
)

func main() {
	//go func() {
	//	log.Println(http.ListenAndServe("localhost:6061", nil))
	//}()

	//done := make(chan bool, 1)

	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				//fmt.Println("for: waiting")
				select {
				case <-ctx.Done():
					fmt.Println("ctx done return")
					return // returning not to leak the goroutine
				case dst <- n:
					n++
					//case <-time.After(2 * time.Second):
					//	fmt.Println("timeout return")
					//	return
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	fmt.Printf("#goroutines : %d\n", runtime.NumGoroutine())
	<-time.After(5 * time.Second)
	fmt.Printf("#goroutines : %d\n", runtime.NumGoroutine())

	//<-done
}
