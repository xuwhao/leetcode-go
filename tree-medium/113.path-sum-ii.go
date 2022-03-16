package main

import (
	"fmt"
	"leetcode-go/util"
)

type TreeNode = util.TreeNode

/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
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
func pathSum(root *TreeNode, targetSum int) [][]int {
	ans, cur := [][]int{}, []int{}
	var search func(node *TreeNode, sum int)
	search = func(node *TreeNode, sum int) {
		sum += node.Val
		cur = append(cur, node.Val)
		defer func() { cur = cur[:len(cur)-1] }()
		if node.Left == nil && node.Right == nil {
			if sum == targetSum {
				ans = append(ans, append([]int{}, cur...))
			}
		}

		if node.Left != nil {
			search(node.Left, sum)
		}

		if node.Right != nil {
			search(node.Right, sum)
		}
	}
	if root != nil {
		search(root, 0)
	}
	return ans
}

// @lc code=end

func testPathSum() {
	input := []int{5, 4, 8, 11, util.NULL, 13, 4, 7, 2, util.NULL, util.NULL, 5, 1}
	target := 22
	root := util.Ints2TreeNode(input)
	fmt.Println(pathSum(root, target))

	input = []int{1, 2, 3}
	target = 5
	root = util.Ints2TreeNode(input)
	fmt.Println(pathSum(root, target))
}
