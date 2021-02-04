package main

import (
	"fmt"
)

func main() {
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	println(b[1:4])
	println(b[:2])
	println(b[2:])
	println(b[:])
	var arr = [5]int{0, 1, 2, 3, 4}
	println(fmt.Printf("%d\n", sum(arr[:])))

	var slice1 = make([]int, 5, 10)
	for i := 0; i < len(slice1); i++ {
		slice1[i] = 5 * i
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	var p *[]int = new([]int)
	println(p)

	s := make([]byte, 5)
	fmt.Printf("Len is %d, Cap is %d\n", len(s), cap(s))

	slice2 := make([]byte, 5, 10)
	slice3 := append(slice2, []byte{2, 3, 3})
	for i := 0; i < len(slice3); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice3[i])
	}
}

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

func append(slice, data []byte) []byte {
	// add an array to slice
	totalLen := len(slice) + len(data)
	var newSlice []byte
	lenSlice := len(slice)
	if totalLen > cap(slice) {
		newSlice = make([]byte, totalLen)
		for i := 0; i < lenSlice; i++ {
			newSlice[i] = slice[i]
		}
	} else {
		newSlice = slice[:totalLen]
	}
	for i := lenSlice; i < totalLen; i++ {
		newSlice[i] = data[i-lenSlice]
	}
	return newSlice
}

// func bufSplit(buffer bytes.Buffer, n int) ([]byte, []byte) {
// 	return buffer[:n], buffer[n:]
// }
