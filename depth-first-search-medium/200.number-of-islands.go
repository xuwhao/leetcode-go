package main

import "fmt"

/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 * 就是求图中连通分量的个数
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	n, m := len(grid), len(grid[0])

	if n == 0 {
		return 0
	}

	grids := make([][]byte, len(grid)+2)
	grids[0], grids[n+1] = make([]byte, m+2), make([]byte, m+2)
	for i := 1; i < n+1; i++ {
		grids[i] = append(append(append([]byte{}, '0'), grid[i-1]...), '0')
	}

	// ans := 0
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j > m+1 { // 越界，返回
			return 0
		}

		if grids[i][j] == '0' || grids[i][j] == 0 { // 如果是海水，返回
			return 0
		}

		grids[i][j] = '0' // 发现一块陆地，淹没

		dfs(i-1, j) // 不断淹没上下左右的陆地
		dfs(i, j-1)
		dfs(i, j+1)
		dfs(i+1, j)

		return 1 // 找到一块陆地，且淹没掉了
	}

	ans := 0
	for i := 0; i < n+2; i++ {
		for j := 0; j < m+2; j++ {
			ans += dfs(i, j)
		}
	}
	return ans
}

// @lc code=end

func testNumIslands() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(grid))
}
