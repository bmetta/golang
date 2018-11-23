package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

// signal when goroutine finishes
var done chan bool

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// make a new channel
	done = make(chan bool)
	// print number of initial goroutines
	fmt.Printf("goroutines start: %d\n", runtime.NumGoroutine())

	// doSomething 10 times concurrently... think web handler
	// style where each requests kicks off another goroutine
	for i := 0; i < 100; i++ {
		go doSomething()
	}
	// wait for all the goroutines to finish, and return
	for i := 0; i < 100; i++ {
		<-done
	}

	fmt.Printf("goroutines end: %d\n", runtime.NumGoroutine())

	// dont exit from main
	<-done
}

func doSomething() {
	// signal we are done doing something
	defer func() { done <- true }()
	// perform a web request
	resp, err := http.Get("https://husobee.github.io/")
	if err != nil {
		log.Fatal(err)
	}

	// avoid memory leak
	defer resp.Body.Close() // close it on defer
	// check the status code of the response
	// if it returns ok, in our example,
	// we dont care about the body, but if
	// not okay, then we need to read the body
	if resp.StatusCode != http.StatusOK {
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		// causes memory leak
		//defer resp.Body.Close()
	}
}
