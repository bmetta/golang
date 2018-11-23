package main

import "fmt"
import "time"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			val, more := <-jobs
			if !more {
				fmt.Println("received all jobs")
				done <- true
				return
			}
			fmt.Println("received jobId: ", val)
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job: ", j)
		time.Sleep(1 * time.Second)
	}

	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
