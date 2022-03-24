package main

/*
 * @lc app=leetcode id=303 lang=golang
 *
 * [303] Range Sum Query - Immutable
 */

// @lc code=start
type NumArray struct {
	data []int
	sum  []int
}

func Constructor(nums []int) NumArray {
	obj := NumArray{data: nums, sum: make([]int, len(nums))}
	obj.sum[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		obj.sum[i] = obj.sum[i-1] + nums[i]
	}

	return obj
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sum[right] - this.sum[left] + this.data[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end
