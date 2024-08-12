package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*

Write a program that demonstrates the use of the select statement with channels.
Create two channels, ch1 and ch2, and two Go routines that send values to these channels at different intervals.
Implement a select statement that waits for values from either channel and prints the received value.
The program should terminate after receiving a total of 10 values.

*/

func SelectorHat() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
			time.Sleep(time.Second * time.Duration(randSleep()))
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch2 <- i
			time.Sleep(time.Second * time.Duration(randSleep()))
		}
	}()

	go func() {
		count := 0
		for {
			select {
			case m1 := <-ch1:
				fmt.Println("Canal 1: ", m1)
				count++
			case m2 := <-ch2:
				fmt.Println("Canal 2: ", m2)
				count++
			}
			if count == 10 {
				wg.Done()
				break
			}
		}
	}()

	wg.Wait()
}

func randSleep() int {
	return rand.Intn(10)
}
