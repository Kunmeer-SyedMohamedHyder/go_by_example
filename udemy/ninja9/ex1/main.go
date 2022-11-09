package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("Hello from go routine1")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hello from go routine2")
		wg.Done()
	}()
	fmt.Println("Hello from main")
	wg.Wait()
}
