package main

import (
	"fmt"
)

func main() {
	factorial(30)
}

func factorial(num int) (res int) {
	if num == 0 {
		res = 1
	} else {
		res = factorial(num-1) * num
	}
	fmt.Printf("Factorial of %d is: %d\n", num, res)
	return
}
