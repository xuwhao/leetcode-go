package main

import (
	"fmt"
)

/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start

func Factorials(x int) int {
	sum := 1
	for i := 2; i <= x; i++ {
		sum *= i
	}
	return sum
}

func permute(nums []int) [][]int {
	n := len(nums)
	v := make([]bool, n)
	var ans [][]int
	inline := []int{}
	var search func(i int)
	search = func(i int) {
		if i == n {
			temp := make([]int, len(inline))
			copy(temp, inline)
			ans = append(ans, temp)
			return
		}
		for j := 0; j < n; j++ {
			if !v[j] {
				v[j] = true
				inline = append(inline, nums[j])
				search(i + 1)
				v[j] = false
				inline = inline[:len(inline)-1]
			}
		}
	}
	search(0)
	return ans
}

// @lc code=end

func testPermute() {
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(permute(nums))
}
