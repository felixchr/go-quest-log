package main

import (
	"fmt"
	"regexp"

	"./pack1"
)

func main() {
	func1()
	func2()
}

func func1() {
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+"

	if ok, m := regexp.MatchString(pat, searchIn); ok {
		fmt.Println("Matched")
		fmt.Println(m)
	}

	re, _ := regexp.Compile(pat)
	for _, s := range re.FindAllString(searchIn, -1) {
		fmt.Println(s)
	}
}

func func2() {
	var test1 string
	test1 = pack1.ReturnStr()
	fmt.Printf("ReturnStr from package1: %s\n", test1)
	fmt.Printf("Integer from package1: %d\n", pack1.Pack1Int)
}
