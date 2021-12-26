package main

import "fmt"

/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */

// @lc code=start

func minSubArrayLen(target int, nums []int) int {
	const MAXINT int = 10001
	sum, ans := 0, MAXINT
	for i, j := 0, 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			if j-i+1 < ans {
				ans = j - i + 1
			}
			sum -= nums[i]
			i += 1
		}
	}
	if ans == MAXINT {
		ans = 0
	}
	return ans
}

// func minSubArrayLen(target int, nums []int) int {
// 	const MAXINT int = 10001
// 	ans := MAXINT
// 	length := len(nums)
// 	for i := 0; i < length; i++ {
// 		sum := 0
// 		for j := i; j < length; j++ {
// 			sum += nums[j]
// 			if sum >= target {
// 				if j-i+1 < ans {
// 					ans = j - i + 1
// 				}
// 				break
// 			}
// 		}
// 	}
// 	if ans == MAXINT {
// 		ans = 0
// 	}
// 	return ans
// }

// @lc code=end

func testMinSubArrayLen() {
	target := 11
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println(minSubArrayLen(target, nums))
}
