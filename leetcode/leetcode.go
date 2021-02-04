package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{2,7,11,15}
	fmt.Println(twoSum(nums, 9))
	num := 12345678900
	fmt.Println(reverse(num))
}

func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int, len(nums))
    for i, num := range nums {
        complement := target - num
        idx, isPresent := numMap[complement]
        if isPresent {
            return []int{idx, i}
        }
        numMap[num] = i
    }
    return []int{0, 0}
}

func reverse(x int) int {
	var flag int
	if x < 0 {
		flag = -1
	} else {
		flag = 1
	}
	str := strconv.Itoa(x)
	slice := []byte(str)
	sl := len(slice)
	hl := sl / 2
	for i := 0; i < hl; i++ {
		slice[i], slice[sl-i-1] = slice[sl-i-1], slice[i]
	}
	str = string(slice)
	rx, _ := strconv.Atoi(str)
	return rx * flag
}