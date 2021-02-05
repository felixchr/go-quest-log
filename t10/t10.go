package main

import (
	"fmt"
)

type st1 struct {
	i1  int
	f1  float32
	str string
}

// type Person struct {
// 	firstName	string
// 	lastName	string
// }

func main() {
	func1()
}

func func1() {
	ms := new(st1)
	ms.i1 = 10
	ms.f1 = 3.14
	ms.str = "haha"

	fmt.Printf("%d\t%f\t%s\n", ms.i1, ms.f1, ms.str)

	ms1 := &st1{5, 6.18, "hehe"}
	fmt.Printf("%d\t%f\t%s\n", ms1.i1, ms1.f1, ms1.str)
}

func upPerson