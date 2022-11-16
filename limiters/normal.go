package main

import (
	"fmt"
	"time"
)

func normal_limiter() {

	requests := make(chan int, 7)
	limiter := time.Tick(time.Second)

	for i := 0; i < 7; i++ {
		requests <- i
	}
	close(requests)

	for req := range requests {
		<-limiter
		fmt.Println("Completed request", req, "at", time.Now())
	}

}
