package main

import (
	"fmt"
	"time"
)

func time_consuming_task(ch chan<- string) {
	time.Sleep(time.Second * 3)
	ch <- "Completed the task"
}

func main() {
	ch := make(chan string)
	go time_consuming_task(ch)
	select {
	case v := <-ch:
		fmt.Println(v)
	case <-time.After(time.Second * 5):
		fmt.Println("Timeout...")

	}
}
