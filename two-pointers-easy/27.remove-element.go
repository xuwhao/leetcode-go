package main

import "fmt"

/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	i, j := 0, len(nums)-1

	for i < j {
		for i < j && nums[i] != val {
			i++
		}
		for i < j && nums[j] == val {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	if nums[i] != val {
		return i + 1
	}
	return i

}

// @lc code=end

func testRemoveElement() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println(removeElement(nums, val), nums)
}
