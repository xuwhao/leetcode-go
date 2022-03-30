package main

/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
func moveZeroes(nums []int) {
	nonZeroIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[nonZeroIndex] = nums[nonZeroIndex], nums[i]
			nonZeroIndex++
		}
	}
}

// @lc code=end
