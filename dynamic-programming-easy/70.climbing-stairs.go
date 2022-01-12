package main

import "fmt"

/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start

// func climbStairs(n int) int {
// 	var search func(x int) int
// 	var ways [46]int

// 	search = func(x int) int {
// 		if x == 1 {
// 			return 1
// 		}
// 		if x == 2 {
// 			return 2
// 		}
// 		i := ways[x-1]
// 		if i == 0 {
// 			i = search(x - 1)
// 			ways[x-1] = i
// 		}
// 		j := ways[x-2]
// 		if j == 0 {
// 			j = search(x - 2)
// 			ways[x-2] = j
// 		}
// 		return i + j
// 	}
// 	return search(n)
// }

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	x, y := 1, 2
	for i := 3; i <= n; i++ {
		x, y = y, x+y
	}
	return y
}

// @lc code=end

func testClimbStairs() {
	n := 45
	fmt.Println(climbStairs(n))
}
