package main

/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */

// @lc code=start
func sortColors(nums []int) {
	p, q := 0, len(nums)-1
	for i := 0; i <= q; i++ {
		if nums[i] == 0 {
			nums[i], nums[p] = nums[p], nums[i]
			p++
		}
		if nums[i] == 2 {
			nums[i], nums[q] = nums[q], nums[i]
			q--
			if nums[i] != 1 { // 如果换出来的不是1，那换出来的数还是要被换
				i--
			}
		}
	}
}

// @lc code=end
