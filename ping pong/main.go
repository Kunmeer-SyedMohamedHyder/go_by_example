package main

import "fmt"

func ping(ping_ch chan<- string, msg string) {
	fmt.Println("Ping complete")
	ping_ch <- msg
}

func pong(ping_ch <-chan string, pong_ch chan<- string) {
	fmt.Println("Pong complete")
	pong_ch <- <-ping_ch
}

func main() {
	ping_ch := make(chan string)
	pong_ch := make(chan string)
	for i := 0; i < 10; i++ {
		go ping(ping_ch, fmt.Sprintf("Hello hello %d", i))
	}
	for i := 0; i < 10; i++ {
		go pong(ping_ch, pong_ch)
		fmt.Println(<-pong_ch)
	}
	// go ping(ping_ch, fmt.Sprintf("Hello hello %d", 10))
	// go pong(ping_ch, pong_ch)
	// fmt.Println(<-pong_ch)

}
