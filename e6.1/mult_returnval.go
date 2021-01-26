package main

import (
	"fmt"
)

func main() {
	var x, y = 9, 8
	s, p, d := sumProdDiff(x, y)
	fmt.Printf("%d+%d=%d, %dx%d=%d, %d-%d=%d\n",
		x, y, s, x, y, p, x, y, d)
	s, p, d = sumProdDiff2(x, y)
	fmt.Printf("%d+%d=%d, %dx%d=%d, %d-%d=%d\n",
		x, y, s, x, y, p, x, y, d)
}

func sumProdDiff(x int, y int) (int, int, int) {
	return x + y, x * y, x - y
}

func sumProdDiff2(x int, y int) (s int, p int, d int) {
	s = x + y
	p = x * y
	d = x - y
	return
}
