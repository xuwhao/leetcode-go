package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for index, num := range nums {
		v, exists := dict[target-num]
		if exists {
			return []int{v, index}
		} else {
			dict[num] = index
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
