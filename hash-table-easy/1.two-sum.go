package main

import "fmt"

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// func twoSum(nums []int, target int) []int {
// 	dict := make(map[int]int)
// 	var ans []int
// 	for index, num := range nums {
// 		v, exists := dict[target-num]
// 		if exists {
// 			if v < index {
// 				ans = []int{v, index}
// 			} else {
// 				ans = []int{index, v}
// 			}
// 		} else {
// 			dict[num] = index
// 		}
// 	}
// 	return ans
// }

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

// func twoSum(nums []int, target int) []int {
// 	m := make(map[int]int)
// 	for k, v := range nums {
// 		if idx, ok := m[target-v]; ok {
// 			return []int{idx, k}
// 		}
// 		m[v] = k
// 	}
// 	return nil
// }

// @lc code=end

func TestTwoSum() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
