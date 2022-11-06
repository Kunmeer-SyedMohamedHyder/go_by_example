package main

import "fmt"

type user struct {
	Name string
	Age  int
}

func (u *user) notify() {
	u.Age = 10
}

type notifier interface {
	notify()
}

func send_notification(n notifier) {
	fmt.Println("Notifier interface...")
	n.notify()
}

func main() {

	u := &user{
		Name: "Syed",
		Age:  22,
	}
	fmt.Println("Notifying age : ", u.Age)
	u.notify()
	fmt.Println("Notifying age : ", u.Age)

}
