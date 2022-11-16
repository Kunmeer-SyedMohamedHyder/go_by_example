package main

import (
	"fmt"
	"time"
)

func burst_limter() {
	limiter := make(chan time.Time, 3)
	requests := make(chan int, 10)

	for i := 0; i < 7; i++ {
		requests <- i
	}
	close(requests)

	for i := 0; i < 3; i++ {
		limiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Second) {
			limiter <- t
		}
	}()

	for r := range requests {
		fmt.Println("Completed req:", r, "at time", <-limiter)
	}

}
