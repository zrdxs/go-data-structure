package main

import (
	"fmt"
	"sync"
)

/*
Implement a worker pool using goroutines and channels.
The worker pool should have a configurable number of workers that can process tasks concurrently.
The tasks should be submitted to the worker pool, and the results should be returned through a channel.
*/

type Task func(input any) (output any)

func Interview() {

	workerLimit := 5
	var wg sync.WaitGroup

	input := make(chan Task)
	output := make(chan any)

	for i := 0; i < workerLimit; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Worker(input, output)
		}()
	}

	go func() {
		for i := 0; i < 10; i++ {
			input <- someTask
		}
		close(input)
	}()

	go func() {
		wg.Wait()
		close(output)
	}()

	// log the result
	for result := range output {
		wg.Add(1)
		go func() {
			wg.Done()
			if val, ok := result.(string); ok {
				fmt.Println(val)
			}
		}()
	}
}

func Worker(input <-chan Task, output chan<- any) {
	for task := range input {
		result := task("Hey!! \n")
		output <- result
	}
}

func someTask(input any) (output any) {
	if val, ok := input.(string); ok {
		fmt.Printf("Received: %s \n", val)
	}
	return "Task Done!!"
}
