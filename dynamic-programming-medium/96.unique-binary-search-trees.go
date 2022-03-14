package main

import "fmt"

/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 */

// @lc code=start
// func numTrees(n int) int {
// 	var search func(left, right int) int
// 	ans := [20][20]int{}

// 	search = func(left, right int) int {
// 		if ans[left][right] != 0 {
// 			return ans[left][right]
// 		}
// 		if left == right {
// 			return 1
// 		}
// 		sum := 0
// 		for middle := left; middle <= right; middle++ {
// 			if left == middle {
// 				ans[middle+1][right] = search(middle+1, right)
// 				sum += ans[middle+1][right]
// 			} else if middle == right {
// 				ans[left][middle-1] = search(left, middle-1)
// 				sum += ans[left][middle-1]
// 			} else {
// 				ans[left][middle-1] = search(left, middle-1)
// 				ans[middle+1][right] = search(middle+1, right)
// 				sum += ans[left][middle-1] * ans[middle+1][right]
// 			}
// 		}
// 		ans[left][right] = sum
// 		return sum
// 	}
// 	return search(1, n)
// }

// func numTrees(n int) int {
// 	var search func(cnt int) int
// 	ans := make([]int, n+1)
// 	search = func(cnt int) int {
// 		if ans[cnt] != 0 {
// 			return ans[cnt]
// 		}
// 		if cnt == 0 || cnt == 1 {
// 			return 1
// 		}
// 		sum := 0
// 		for root := 1; root <= cnt; root++ { // 以 root 为根，左子树都小于 root，共 root-1 个
// 			sum += search(root-1) * search(cnt-root) // 右子树共 cnt-root 个，不用管具体数值，因为 2,3,4 和 1,2,3 在 BST 的结构上是一样的
// 		}
// 		ans[cnt] = sum
// 		return ans[cnt]
// 	}
// 	return search(n)
// }

func numTrees(n int) int {
	ans := [20]int{}
	ans[0], ans[1] = 1, 1
	for i := 2; i <= n; i++ {
		for root := 1; root <= i; root++ {
			ans[i] += ans[root-1] * ans[i-root]
		}
	}
	return ans[n]
}

// @lc code=end

func testNumTrees() {
	n := 19
	fmt.Println(numTrees(n))
}
