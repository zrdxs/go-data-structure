package main

import (
	"fmt"
	"time"
)

/*
	2)Create a program that demonstrates the use of a buffered channel for rate limiting.
	The program should have a producer goroutine that generates tasks at a faster rate than the consumer goroutine can process them.
	Use a buffered channel to limit the number of tasks that can be queued up, preventing the producer from overwhelming the consumer.
*/

func RateLimite() {
	limiter := 5
	producerRate := 100 * time.Millisecond
	consumerRate := 1000 * time.Millisecond

	tasks := make(chan int, limiter)

	go func() {
		for i := 0; i < 1000; i++ {
			select {
			case tasks <- i:
				fmt.Println("Sending task!!")
				time.Sleep(producerRate)
			default:
				fmt.Println("Rate limite achived, waiting!!")
				time.Sleep(producerRate)
			}
		}
	}()

	go func() {
		for task := range tasks {
			fmt.Printf("Task processed: %d \n", task)
			time.Sleep(consumerRate)
		}
	}()
	time.Sleep(10 * time.Second)
}
