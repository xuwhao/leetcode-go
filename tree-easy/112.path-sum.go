package main

import (
	"fmt"
	"leetcode-go/util"
)

/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	var search func(node *TreeNode, sum int) bool
	search = func(node *TreeNode, sum int) bool {
		sum += node.Val
		if node.Left == nil && node.Right == nil {
			if sum == targetSum {
				return true
			}
			return false
		}
		flag := false
		if node.Left != nil {
			flag = search(node.Left, sum)
		}
		if flag {
			return true
		}
		if node.Right != nil {
			flag = search(node.Right, sum)
		}
		return flag
	}
	if root == nil {
		return false
	}
	return search(root, 0)
}

// @lc code=end
func testHasPathSum() {
	input := []int{5, 4, 8, 11, util.NULL, 13, 4, 7, 2, util.NULL, util.NULL, util.NULL, 1}
	target := 22
	root := util.Ints2TreeNode(input)
	fmt.Println(hasPathSum(root, target) == true)

	input = []int{1, 2, 3}
	target = 5
	root = util.Ints2TreeNode(input)
	fmt.Println(hasPathSum(root, target) == false)
}
