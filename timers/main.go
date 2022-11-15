package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 3)

	<-timer1.C
	fmt.Println("Timer 1 fired....")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired...")
	}()
	time.Sleep(time.Second * 3)
	stopped := timer2.Stop()
	if stopped {
		fmt.Println("Timer 2 stopped..")
	} else {
		fmt.Println("Timer already expired...")
	}
}
