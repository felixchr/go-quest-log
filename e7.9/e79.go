package main

import (
	"fmt"
)

func main() {
	var slice = []int{1, 2, 3}
	slice = sliceMulti(slice, 5)
	fmt.Printf("%s\n", slice)
	for i := 0; i < len(slice); i++ {
		slice[i] = i + 2
	}
	slice = filter(slice, isEven)
	fmt.Printf("%s\n", slice)
	s1 := []byte("abcde")
	s2 := []byte("de")
	fmt.Printf("%s\n", insertStringSlice(s1, s2, 2))
}

func sliceMulti(slice []int, factor int) []int {
	newSlice := make([]int, len(slice) * factor)
	copy(newSlice, slice)
	return newSlice
}

func filter(slice []int, fn func(int) bool) []int {
	newSlice := make([]int, 0)
	for _, num := range(slice) {
		if fn(num) {
			newSlice = append(newSlice, num)
		}
	}
	return newSlice
}

func isEven(i int) bool {
	if i % 2 == 0 {
		return true
	} else {
		return false
	}
}

func insertStringSlice(slice []byte, sliceToInsert []byte, idx int) []byte {
	newS := append(slice[:idx], append(sliceToInsert, slice[idx:]...)...)
	return newS
}