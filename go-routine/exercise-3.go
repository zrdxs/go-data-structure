package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*

Exercise 3: Pipeline

Design a pipeline to process a large dataset. Create three stages: producer, processor, and consumer.
The producer stage should generate a stream of data and send it to the processor stage via a channel.
The processor stage should receive the data, perform some processing on it, and send the processed data to the consumer stage via another channel.
The consumer stage should receive the processed data and perform some final action, such as printing the result.

*/

func Pipeline() {

	pipe := make(chan interface{})
	process := make(chan interface{})

	var wg sync.WaitGroup

	wg.Add(1)
	go producer(pipe, &wg)

	go processor(pipe, process)
	go consumer(process, &wg)

	wg.Wait()
	close(process)
}

func producer(pipe chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 40; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		if r.Int()%2 == 0 {
			pipe <- "Belive in God"
		} else {
			pipe <- r.Int()
		}
		//time.Sleep(3 * time.Second)
	}
	close(pipe)
}

func processor(pipe, process chan interface{}) {

	for {
		select {
		case val, ok := <-pipe:
			if !ok {
				close(process)
				return
			}

			switch v := val.(type) {
			case int:
				msg := fmt.Sprintf("Oh I received an integer %d", v)
				process <- msg
			case string:
				msg := fmt.Sprintf("I received an important message %s", v)
				process <- msg
			}
		}
	}

}

func consumer(process chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range process {
		fmt.Printf("This is the message that i receive: %v \n", msg)
	}
}
