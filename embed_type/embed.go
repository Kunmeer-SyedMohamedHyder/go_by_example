package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("DESCRIBING.....%v", b.num)
}

type container struct {
	base
	msg string
}

func main() {
	cnt := container{
		base: base{num: 10},
		msg:  "Hello world",
	}

	fmt.Println(cnt.describe())
	fmt.Println(cnt.base, cnt.base.num, cnt.msg)

	type describer interface {
		describe() string
	}

	var d describer = cnt
	fmt.Println(d.describe())

}
