package main

import "fmt"

/*
 * @lc app=leetcode id=695 lang=golang
 *
 * [695] Max Area of Island
 */

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	n, m := len(grid), len(grid[0])

	if n == 0 {
		return 0
	}

	grids, flags := make([][]int, len(grid)+2), make([][]int, len(grid)+2)
	grids[0], grids[n+1] = make([]int, m+2), make([]int, m+2)
	for i := 1; i < n+1; i++ {
		grids[i] = append(append(append([]int{}, 0), grid[i-1]...), 0)
		flags[i] = make([]int, m+2)
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j > m+1 { // 越界，返回
			return 0
		}

		if grids[i][j] == 0 || flags[i][j] == 1 { // 如果找过了
			return 0
		}

		flags[i][j] = 1

		return 1 + dfs(i-1, j) + dfs(i, j-1) + dfs(i, j+1) + dfs(i+1, j)
	}

	max := 0
	for i := 0; i < n+2; i++ {
		for j := 0; j < m+2; j++ {
			ans := dfs(i, j)
			if ans > max {
				max = ans
			}
		}
	}

	return max
}

// @lc code=end

func testMaxAreaOfIsland() {
	grid := [][]int{
		{0, 0, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 1},
		{0, 0, 0, 1},
	}

	fmt.Println(maxAreaOfIsland(grid))
}
