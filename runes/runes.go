package main

import (
	"fmt"
)

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}

func main() {
	s := "สวัสดี"

	for ind, val := range s {
		fmt.Printf("%#U starts at %d\n", val, ind)
		examineRune(val)
	}
}
