package main

import "fmt"

/*
 * @lc app=leetcode id=130 lang=golang
 *
 * [130] Surrounded Regions
 */

// @lc code=start
func solve(board [][]byte) {
	n, m := len(board), len(board[0])
	if n == 0 {
		return
	}

	grids, flags := make([][]byte, len(board)+2), make([][]byte, len(board)+2)
	grids[0], grids[n+1] = make([]byte, m+2), make([]byte, m+2)
	for i := 1; i < n+1; i++ {
		grids[i] = append(append(append([]byte{}, 'X'), board[i-1]...), 'X')
		flags[i] = make([]byte, m+2)
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if grids[i][j] == 'X' || grids[i][j] == 0 || flags[i][j] == 1 {
			return
		}

		if grids[i][j] == 'O' {
			flags[i][j] = 1
		}

		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if i == 1 || i == n || j == 1 || j == m {
				dfs(i, j)
			}
		}
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if flags[i][j] == 0 {
				board[i-1][j-1] = 'X'
			}
		}
	}
}

// @lc code=end
func testSolve() {
	// board := [][]byte{
	// 	{'X', 'X', 'X', 'X'},
	// 	{'X', 'O', 'O', 'X'},
	// 	{'X', 'X', 'O', 'X'},
	// 	{'X', 'O', 'X', 'X'},
	// }
	board := [][]byte{
		{'O', 'O'},
		{'O', 'O'},
	}
	solve(board)
	fmt.Println(board)
}
