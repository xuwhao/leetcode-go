package main

import "fmt"

/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */

// @lc code=start
func isHappy(n int) bool {
	sum := n
	counter := make(map[int]int)
	for sum != 1 {
		_, exists := counter[sum]
		if exists {
			break
		}
		counter[sum]++
		n = sum
		sum = 0
		for n != 0 {
			t := n % 10
			n = n / 10
			sum += t * t
		}

	}
	return sum == 1
}

// @lc code=end

func TestIsHappy() {
	n := 202
	fmt.Println(isHappy(n))
}
