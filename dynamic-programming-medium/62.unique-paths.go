package main

import "fmt"

/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */

// @lc code=start

func uniquePaths(m int, n int) int {
	// 多开一层, 从 1 开始算，这样计算方便点，下标不会越界
	var path [101][101]int
	path[1][1] = 1
	for i := 2; i < m+n; i++ { // 共有 m+n-1 条对角线
		for j := 1; j <= m; j++ {
			if j > i { // 超出了行高，因为不一定是正方形
				continue
			}
			col := i - j + 1
			if col > n { // 超出了列宽，也不用计算，越界
				continue
			}
			path[j][col] = path[j-1][col] + path[j][col-1]
		}
	}
	return path[m][n]
}

// func uniquePaths(m int, n int) int {
// 	var search func(i, j int) int
// 	search = func(i, j int) int {
// 		if i == m-1 && j == n-1 {
// 			return 1
// 		}
// 		if i > m-1 || j > n-1 {
// 			return 0
// 		}
// 		return search(i+1, j) + search(i, j+1)
// 	}
// 	return search(0, 0)
// }

// func uniquePaths(m int, n int) int {
// 	var path [101][101]int
// 	var search func(i, j int) int
// 	search = func(i, j int) int {
// 		if i == m-1 && j == n-1 {
// 			return 1
// 		}
// 		if i > m-1 || j > n-1 {
// 			return 0
// 		}
// 		x := path[i+1][j]
// 		if x == 0 {
// 			x = search(i+1, j)
// 			path[i+1][j] = x
// 		}
// 		y := path[i][j+1]
// 		if y == 0 {
// 			y = search(i, j+1)
// 			path[i][j+1] = y
// 		}
// 		return x + y
// 	}
// 	return search(0, 0)
// }

// @lc code=end

func TestUniquePaths() {
	m, n := 50, 9
	fmt.Println(uniquePaths(m, n))
}
