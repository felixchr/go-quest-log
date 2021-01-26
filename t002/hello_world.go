package main

import (
	"fmt"
)

var (
	d int
	e int
	f string
)

func main() {
	a, b, c := 5, 7, "test"
	d, e, f = a, b, c
	s := fmt.Sprintln("I got some values: ", a, b, c, d, e, f)
	fmt.Println(s)
}
