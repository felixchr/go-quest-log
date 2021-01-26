package main

import (
	"fmt"
	"time"
)

var d time.Duration

func main() {
	t := time.Now()

	fmt.Println(t.Format("20201125102334"))
	d = 1e9 * 60 * 60 * 24 * 7
	fmt.Println(d)
}
