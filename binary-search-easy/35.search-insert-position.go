package main

import "fmt"

/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if target == nums[mid] {
			return mid
		}
		if target < nums[mid] {
			right = mid - 1
		}
		if target > nums[mid] {
			left = mid + 1
		}
	}

	return left
}

// @lc code=end
func testSearchInsert() {
	nums := []int{1, 3, 5, 7, 8, 9}
	target := 6
	fmt.Println(searchInsert(nums, target))
}
