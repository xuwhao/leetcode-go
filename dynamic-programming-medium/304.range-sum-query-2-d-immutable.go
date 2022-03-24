package main

/*
 * @lc app=leetcode id=304 lang=golang
 *
 * [304] Range Sum Query 2D - Immutable
 */

// @lc code=start
type NumMatrix struct {
	data [][]int
	sum  [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	size := len(matrix)
	obj := NumMatrix{data: matrix, sum: make([][]int, size+1)}

	obj.sum[0] = make([]int, len(matrix[0])+1)
	for i := 0; i < size; i++ {
		obj.sum[i+1] = make([]int, len(matrix[i])+1)
		for j := 0; j < len(matrix[i]); j++ {
			obj.sum[i+1][j+1] = obj.data[i][j] + obj.sum[i][j+1] + obj.sum[i+1][j] - obj.sum[i][j]
		}
	}

	return obj
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sum[row2+1][col2+1] + this.sum[row1][col1] - this.sum[row1][col2+1] - this.sum[row2+1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end
