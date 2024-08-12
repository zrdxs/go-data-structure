package main

import (
	"fmt"
	"sync"
)

/*
	Fan-out/Fan-in Write a program that calculates the square of each number in a given list concurrently.
	Implement a function square that takes an input channel and an output channel as arguments.
	This function should read a number from the input channel, calculate its square, and send the result to the output channel.
	Use multiple Go routines to handle the calculation for each number.
	Finally, read the results from the output channel and print them.
*/

/*
	CHALLENGE - MAKE THIS FUNC IN LOTES WITH CONCURRENCY LIMIT
*/

func FanInFanOut() {
	numbers := []float64{2, 4, 5, 6, 8, 50, 10, 23}

	input := make(chan float64)
	output := make(chan float64)

	var wg sync.WaitGroup

	for i := 0; i < len(numbers); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			calculateNumberSquare(input, output)
		}()
	}

	go func() {
		for _, n := range numbers {
			input <- n
		}
		close(input)
	}()

	go func() {
		wg.Wait()
		close(output)
	}()

	for v := range output {
		fmt.Println(v)
	}

}

func calculateNumberSquare(input chan float64, output chan float64) {
	for num := range input {
		output <- num * num
	}
}
