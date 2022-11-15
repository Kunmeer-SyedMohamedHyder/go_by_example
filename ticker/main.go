package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)

	done := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			select {
			case <-done:
				fmt.Println("Done ticking...")
				return
			case <-ticker.C:
				fmt.Println("Tick tick", i)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	ticker.Stop()
	done <- true
}
