package main
// import (
// 	"fmt"
// )

func main() {
	str := "I am a string"
	s1, s2 := strSplit(str, 3)
	println(s1)
	println(s2)
	println(str[len(str)/2:] + str[:len(str)/2])
	println(strReverse(str))
}

func strSplit(s string, i int) (string, string) {
	newS := []byte(s)
	return string(newS[:i]), string(newS[i:])
}

func strReverse(s string) string {
	oldSlice := []byte(s)
	newSlice := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		newSlice[i] = oldSlice[len(s) - i - 1]
	}
	return string(newSlice)
}