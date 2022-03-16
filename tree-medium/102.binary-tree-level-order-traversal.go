package main

import (
	"fmt"
)

// type TreeNode util.TreeNode

/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans, queue, head, tail, cnt := [2001][]int{}, [2001]*TreeNode{}, 0, 0, 0
	queue[tail], tail = root, tail+1
	queue[tail], tail = nil, tail+1
	for head != tail {
		cur := queue[head]
		head++
		if cur != nil {
			ans[cnt] = append(ans[cnt], cur.Val)
			if cur.Left != nil {
				queue[tail], tail = cur.Left, tail+1
			}
			if cur.Right != nil {
				queue[tail], tail = cur.Right, tail+1
			}
		} else {
			if head != tail {
				queue[tail], tail = nil, tail+1
				cnt++
			}

		}
	}
	return ans[:cnt+1]
}

// @lc code=end

var NULL = -1 << 63

func testLevelOrder() {
	input := []int{3, 9, 20, NULL, NULL, 15, 7}
	root := Ints2TreeNode(input)
	fmt.Println(levelOrder(root))

}

func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: ints[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
