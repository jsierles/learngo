package main

import (
	"fmt"
	"sync"
)

func doit(id int, result chan int, stop chan bool, waiter *sync.WaitGroup) {

	defer waiter.Done()
	for i := 1; i <= 1000; i++ {
		select {
		case <-stop:
			fmt.Printf("Stopping %d at %d invocations\n", id, i)
			return
		default:
			fmt.Printf("%d(%d): %d\n", id, i, id%i)
		}
	}
	result <- id
}

func main() {
	result := make(chan int, 3)
	stop := make(chan bool)
	waiter := new(sync.WaitGroup)

	for i := 0; i <= 2; i++ {
		waiter.Add(1)
		go doit(i, result, stop, waiter)
	}

	var winner int = <-result
	// This approach appears to behave similarly to just calling close(result) without the signal channel

	for i := 0; i < 2; i++ {
		stop <- true
	}

	waiter.Wait()
	fmt.Printf("Winner was %d\n", winner)
}
