package main

import (
	"fmt"
	"sync"
)

func FirstChan() {

	msg := make(chan string)
	var wg sync.WaitGroup

	wg.Add(1)

	go goroutine1(msg, &wg)
	go goroutine2(msg, &wg)

	go func() {
		msg <- "this is some message"
	}()

	wg.Wait()
	close(msg)
}

func goroutine1(pipe chan string, wg *sync.WaitGroup) {
	msg := <-pipe
	fmt.Printf("Eu a go routine 1 recebi a mensagem %s", msg)
	wg.Done()
}

func goroutine2(pipe chan string, wg *sync.WaitGroup) {
	msg := <-pipe
	fmt.Printf("Eu a go routine 2 recebi a mensagem %s", msg)
	wg.Done()
}
