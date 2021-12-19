package main

/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	pos := 0
	for i, j := 0, 1; j <= len(nums); j++ {
		if j == len(nums) || nums[j-1] != nums[j] {
			nums[pos] = nums[i]
			pos++
			i = j
		}
	}
	return pos
}

// @lc code=end
