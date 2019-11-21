package main

import "fmt"
import "sync"

/*
 * Waitgroup blocks program and waits for a set of goroutines
 * to finsh before moving to the next steps of execution.
 *
 * waitgroups has following functions
 * Add(int): wait for the number of go routines to get exected
 * Done(): go routine calls this when its execution finishes
 * Wait(): it blocks the program until number of specified go
 *         routines finishes their execution
 */
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Go Routine 1")
		wg.Done()
	}()
	go func() {
		fmt.Println("Go Routine 2")
		wg.Done()
	}()
	wg.Wait()
}
