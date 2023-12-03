package main

import "sync"

// when you send to a channel, you put the argument such as
// <argument-name> <- type
func producer(messages chan<- int, done chan<- bool) {
	for i := 0; i < 10; i++ {
		messages <- i
	}

	done <- true
	println("producer exited")
}

func consumer(messages chan int, done chan bool) {
	val := 0
	for {
		select {
		case val = <-messages:
			println("received", val)
		case <-done:
			println("consumer exited")
			return

		}
	}
}

func main() {
	var waitGroup sync.WaitGroup
	// done is here for the producer to tell the consumer the work is done
	done := make(chan bool)

	// all the messages being exchanged between producer and consumer
	msgs := make(chan int)

	// we are going to wait two tasks
	waitGroup.Add(2)

	// launch the first task, once the task is done, we notify the wait group
	go func() {
		defer waitGroup.Done()
		consumer(msgs, done)
	}()

	// launch the producer, and we notify the wait group once done
	go func() {
		defer waitGroup.Done()
		go producer(msgs, done)
	}()

	// wait for all tasks to complete
	waitGroup.Wait()
	println("main program exited")
}
