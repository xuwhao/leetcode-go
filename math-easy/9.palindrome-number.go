package main

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
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

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	y := reverse(x)
	return x == y
}

// @lc code=end

func testIsPalindrome() {
	x := 1213
	fmt.Println(isPalindrome(x))
}
