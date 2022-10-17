package main

import "fmt"

func main() {
	d := make(map[string]int)

	d["a"] = 1
	d["b"] = 2
	d["c"] = 3
	fmt.Println(d)

	key, val := d["c"]
	fmt.Println(key, val)

	delete(d, "c")
	fmt.Println(d)

	_, val = d["c"]
	fmt.Println(val)
}
