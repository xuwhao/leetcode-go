package main

import "fmt"

/*
 * @lc app=leetcode id=560 lang=golang
 *
 * [560] Subarray Sum Equals K
 */

// @lc code=start

type PreSum struct {
	data    []int
	sum     []int
	counter map[int]int
}

func Constructor(nums []int) PreSum {
	obj := PreSum{data: nums, sum: make([]int, len(nums)+1), counter: make(map[int]int)}
	obj.sum[0] = 0
	obj.counter[0]++

	for i := 0; i < len(nums); i++ {
		obj.sum[i+1] = obj.sum[i] + obj.data[i]
	}

	return obj
}

func (this *PreSum) getSum(idx int) int {
	return this.sum[idx+1]
}

func subarraySum(nums []int, k int) int {
	prev, ans := Constructor(nums), 0

	for i := 0; i < len(nums); i++ {
		sum := prev.getSum(i)
		minus := sum - k
		if _, ok := prev.counter[minus]; ok {
			ans += prev.counter[minus]
		}
		prev.counter[sum]++
	}

	return ans
}

// @lc code=end

func testSubArraySum() {
	nums := []int{1}
	k := 0
	fmt.Println(subarraySum(nums, k))
}
