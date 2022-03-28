package main

import "fmt"

/*
 * @lc app=leetcode id=704 lang=golang
 *
 * [704] Binary Search
 */

// @lc code=start
func search(nums []int, target int) int {
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

	return -1
}

// @lc code=end
func testSearch() {
	nums := []int{}
	target := 2
	fmt.Println(search(nums, target))
}
