package main

import "fmt"

func main() {
	s := make([]string, 3, 5)
	fmt.Println(len(s), cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println(len(s), cap(s))

	s = append(s, "d")
	fmt.Println(len(s), cap(s))

	s = append(s, "e")
	fmt.Println(len(s), cap(s))

	s = append(s, "f")
	fmt.Println(len(s), cap(s))

	t := make([][]int, 3)
	for i := 0; i < len(t); i++ {
		inner_len := i + 1
		t[i] = make([]int, inner_len)
		for j := 0; j < inner_len; j++ {
			t[i][j] = i + j
		}
	}

	fmt.Println(len(t), cap(t))
	fmt.Println(t)

}
