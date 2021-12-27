package main

import "fmt"

/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray

	只需要考虑当前第 i 个要不要取。 如果加上 nums[i], 和变成负数了，那么后面任何数加上这段负数区间，
	相比于不加，都是变小的，所以到 i 这个位置，直接另起一段了，前面是负的了都不要了。
	然后每次计算的时候保存下，如果比当前的最大值大就更新一下最大值。

*/

// @lc code=start
func maxSubArray(nums []int) int {
	sum, ans := 0, -10001
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > ans {
			ans = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return ans
}

// @lc code=end

func TestMaxSubArray() {
	nums := []int{-6}
	fmt.Println(maxSubArray(nums))
}
