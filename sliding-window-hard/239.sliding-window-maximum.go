package main

import "fmt"

/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 */

// @lc code=start

type MQueue struct {
	value []int
}

func (mq *MQueue) empty() bool {
	return len(mq.value) == 0
}

func (mq *MQueue) max() (int, bool) {
	if !mq.empty() {
		return mq.value[0], true
	}
	return 0, false
}

func (mq *MQueue) remove(x int) {
	if !mq.empty() {
		if mq.value[0] == x {
			mq.value = mq.value[1:]
		}
	}
}

func (mq *MQueue) insert(x int) {
	for !mq.empty() && mq.value[len(mq.value)-1] < x {
		mq.value = mq.value[:len(mq.value)-1]
	}
	mq.value = append(mq.value, x)
}

func maxSlidingWindow(nums []int, k int) []int {
	ans, window := []int{}, MQueue{value: []int{}}
	i, j := 0, 0

	for ; j < k-1; j++ { // k-1 是精髓，下一次就是从 k 开始, 等于 [i,)
		window.insert(nums[j])
	}

	for ; j < len(nums); i, j = i+1, j+1 { // [i, j) 左闭右开
		window.insert(nums[j])
		if max, ok := window.max(); ok {
			ans = append(ans, max)
		}
		window.remove(nums[i])
	}

	return ans
}

// @lc code=end
func testMaxSlidingWindow() {
	nums := []int{1}
	k := 1
	fmt.Println(maxSlidingWindow(nums, k))
}
