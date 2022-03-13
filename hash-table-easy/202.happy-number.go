package main

import "fmt"

/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */

// @lc code=start
// func isHappy(n int) bool {
// 	sum := n
// 	counter := make(map[int]int)
// 	for sum != 1 {
// 		_, exists := counter[sum]
// 		if exists {
// 			break
// 		}
// 		counter[sum]++
// 		n = sum
// 		sum = 0
// 		for n != 0 {
// 			t := n % 10
// 			n = n / 10
// 			sum += t * t
// 		}

// 	}
// 	return sum == 1
// }

func isHappy(n int) bool {
	slow, fast := n, step(n)
	for fast != 1 && slow != fast {
		slow = step(slow)
		fast = step(step(fast))
	}
	return fast == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

// @lc code=end

func TestIsHappy() {
	n := 202
	fmt.Println(isHappy(n))
}
