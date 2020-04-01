package main

import (
	"fmt"
	"time"
)

func worker(work <-chan int, results chan<- int) {

	for x := range work {

		fmt.Printf("Work start on %d\n", x)

		// Wait for the timer to tick over to simulate a bunch of work.
		// The timer is also a channel!
		time.Sleep(5 * time.Second)

		fmt.Printf("Work end on %d\n", x)

		// Send the result into the results chan for printing.
		results <-x*2
	}
	fmt.Printf("Goroutine exiting.\n")
}

func main() {

	workChan := make(chan int)
	resultChan := make(chan int, 10)

	// Start our goroutines
	for i := 0; i < 5; i++ {
		go worker(workChan, resultChan)
	}

	// Prime the work channel with units of work.
	for i := 0; i < 10; i++ {
		workChan <- i
	}

	for i := 0; i < 10; i++ {
		x := <-resultChan
		fmt.Printf("Result: %d\n", x)
	}

	// Closing the work channel should kill all of our goroutines.
	close(workChan)

	// The runtime will kill any goroutines immediately if the program exits, give them time to
	// clean themselves up.

	// Close our result channel
	close(resultChan)
}
