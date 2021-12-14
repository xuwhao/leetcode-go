package main

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */
// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	diff := 100001
	ans := 0
	for i := 0; i < len(nums); i++ {
		// // 情况 2
		// if nums[i] > target {
		// 	break
		// }
		// 情况 3, [a, b, c] 跳过 a 的重复
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 情况 1, [a, b, c] 确保 a < b < c
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			temp := int(math.Abs(float64(sum - target)))
			if temp == 0 {
				return sum
			}
			if temp < diff {
				diff = temp
				ans = sum
			}
			if sum < target {
				j++
			} else if sum > target {
				k--
			}
		}
	}
	return ans
}

// @lc code=end
