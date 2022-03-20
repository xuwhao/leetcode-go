package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 */

// @lc code=start

func nSumTarget(nums []int, n, start, target int) [][]int {
	sz := len(nums)

	ans := [][]int{}
	if n < 2 || sz < n {
		return ans
	}

	if n == 2 {
		lo, hi := start, sz-1
		for lo < hi {
			sum := nums[lo] + nums[hi]
			left, right := nums[lo], nums[hi]
			if sum < target {
				for lo < hi && nums[lo] == left {
					lo++
				}
			} else if sum > target {
				for lo < hi && nums[hi] == right {
					hi--
				}
			} else {
				ans = append(ans, []int{left, right})
				for lo < hi && nums[lo] == left {
					lo++
				}
				for lo < hi && nums[hi] == right {
					hi--
				}
			}
		}
	} else {
		for i := start; i < sz; i++ {
			sub := nSumTarget(nums, n-1, i+1, target-nums[i])
			for _, arr := range sub {
				arr = append(arr, nums[i])
				ans = append(ans, arr)
			}
			for i < sz-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}

	return ans
}

type IntSlice []int

func (a IntSlice) Len() int           { return len(a) }
func (a IntSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a IntSlice) Less(i, j int) bool { return a[i] < a[j] }

func fourSum(nums []int, target int) [][]int {
	sort.Sort(IntSlice(nums))
	return nSumTarget(nums, 4, 0, target)
}

// @lc code=end

func testFourSum() {
	nums := []int{2, 2, 2, 2, 2}
	target := 8
	fmt.Println(fourSum(nums, target))
}
