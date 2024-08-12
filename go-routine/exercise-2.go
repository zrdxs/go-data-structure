package main

import (
	"fmt"
	"sync"
)

/*
 Create a worker pool that processes a list of jobs concurrently.
 Implement a function worker that takes an input channel and an output channel as arguments.
 The input channel will receive jobs, and each worker should read a job, perform some processing on it, and send the result to the output channel.
 Use a fixed number of worker Go routines to handle the jobs.
 Finally, read the results from the output channel and print them.

*/

func JobProcessor() {
	jobs := []func() interface{}{message}

	workersLimit := 5

	var wg sync.WaitGroup

	inputC := make(chan func() interface{})
	outputC := make(chan interface{})

	// ask this to chat
	for _, job := range jobs {
		go func(j func() interface{}) {
			inputC <- j
		}(job)
	}

	for i := 0; i < workersLimit; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(inputC, outputC)
		}()
	}

	go func() {
		wg.Wait()
		close(outputC)
	}()

	for out := range outputC {
		fmt.Println(out)
	}

}

func worker(input chan func() interface{}, output chan interface{}) {
	for {
		select {
		case funcT := <-input:
			if funcT != nil {
				o := funcT()
				output <- o
			} else {
				return
			}
		default:
			close(input)
		}
	}
}

/* func sum() interface{} {
	return 1 + 1
} */

func message() interface{} {
	return "Jesus is great"
}
