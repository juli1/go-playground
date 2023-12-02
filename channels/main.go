package main

import (
	"fmt"
	"time"
)

func simpleChannel() {
	// make one channel
	messages := make(chan string, 1)
	// write into the channel
	messages <- "foobar"
	// receive the value
	received := <-messages
	// print it
	println(received)
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func workers() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

func main() {
	simpleChannel()
	workers()
}
