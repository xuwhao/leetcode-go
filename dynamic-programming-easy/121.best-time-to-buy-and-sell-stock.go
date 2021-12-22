package main

/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	min, ans := 10001, 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			temp := prices[i] - min
			if temp > ans {
				ans = temp
			}
		}
	}
	return ans
}

// @lc code=end
