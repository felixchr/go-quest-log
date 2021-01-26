package main

import (
	"fmt"
)

func loop1() {
	for i := 0; i < 6; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("G")
		}
		fmt.Println()
	}
}

func loop2() {
	var s = ""
	for i := 0; i < 6; i++ {
		s = s + "G"
		fmt.Println(s)
	}
}

func loopRange() {
	var s = "Hi I'm very tired!"
	for pos, char := range s {
		fmt.Printf("Letter %d: %c\n", pos, char)
	}
}

func loop3() {
	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}
}

func main() {
	loop1()
	loop2()
	loopRange()
	loop3()
}
