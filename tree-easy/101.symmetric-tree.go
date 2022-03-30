package main

/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
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
func isMirror(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil {
		return false
	}

	return node1.Val == node2.Val && isMirror(node1.Right, node2.Left) && isMirror(node1.Left, node2.Right)
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

// @lc code=end
