package main

import "fmt"

/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

// @lc code=start
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		// target 在有序区间
		if nums[left] <= nums[mid] && nums[left] <= target && target <= nums[mid] {
			right = mid - 1
			continue
		}

		if nums[mid+1] <= nums[right] && nums[mid+1] <= target && target <= nums[right] {
			left = mid + 1
			continue
		}

		// target 在不一定有序的区间
		if nums[left] <= nums[mid] == false {
			right = mid - 1
			continue
		}

		if nums[mid+1] <= nums[right] == false {
			left = mid + 1
			continue
		}

		// 所有区间都有序但是没找到
		return -1
	}
	return -1
}

// @lc code=end

func testSearch() {
	nums := []int{6, 7, 0, 1, 2, 4, 5}
	target := 3
	fmt.Println(search(nums, target))
}
