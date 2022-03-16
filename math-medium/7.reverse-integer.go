package main

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start
func reverse(x int) int {
	ans := 0
	for x != 0 {
		if ans > math.MaxInt32/10 || ans < math.MinInt32/10 {
			return 0
		}
		temp := x % 10
		ans = ans*10 + temp
		x = x / 10
	}
	return ans
}

// @lc code=end

func testReverse() {
	x := -2147447412
	fmt.Println(reverse(x))
}
