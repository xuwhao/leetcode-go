package main

import "fmt"

/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start

func binarySearch(nums []int, left, right, target int, isFirst bool) int {
	if left == right {
		return left
	}

	if left > right {
		return right
	}

	mid := (left + right) / 2
	if nums[mid] == target {
		if isFirst {
			return binarySearch(nums, left, mid-1, target, isFirst)
		} else {
			return binarySearch(nums, mid+1, right, target, isFirst)
		}
	} else if target < nums[mid] {
		return binarySearch(nums, left, mid-1, target, isFirst)
	} else {
		return binarySearch(nums, mid+1, right, target, isFirst)
	}
}

func searchRange(nums []int, target int) []int {
	first, last, ans := -1, -1, []int{-1, -1}

	if len(nums) == 0 {
		return ans
	}

	first = binarySearch(nums, 0, len(nums)-1, target, true)

	if first == len(nums)-1 {
		if nums[first] != target {
			first = -1
		}
	} else if first == -1 {
		if nums[first+1] == target {
			first++
		}
	} else {
		if nums[first] == target {

		} else if nums[first+1] == target {
			first++
		} else if nums[first] != target && nums[first+1] != target {
			first = -1
		}
	}

	if first != -1 {
		last = binarySearch(nums, first, len(nums)-1, target, false)

		if last == 0 {
			if nums[last] != target {
				last = first
			}
		} else {
			if nums[last] == target {

			} else if nums[last-1] == target {
				last--
			} else if nums[last-1] != target && nums[last] != target {
				last = first
			}
		}

		ans = []int{first, last}
	}

	return ans

}

// @lc code=end

func testSearchRange() {
	nums := []int{2, 2}
	target := 2
	fmt.Println(searchRange(nums, target))
}
