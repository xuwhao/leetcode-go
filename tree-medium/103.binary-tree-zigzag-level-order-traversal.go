package main

import "fmt"

/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func zigzagLevelOrder(root *TreeNode) [][]int {
	stk := [256]*TreeNode{}
	stack := [256]*TreeNode{}
	tail, top, flag := 0, 0, 0
	ans, inline := [][]int{}, []int{}

	if root == nil {
		return [][]int{}
	}

	stk[tail] = root
	tail++

	for tail > 0 || top > 0 {
		if flag == 0 { // using stk, left to right
			tail--
			node := stk[tail]

			inline = append(inline, node.Val)

			if node.Left != nil {
				stack[top] = node.Left
				top++
			}
			if node.Right != nil {
				stack[top] = node.Right
				top++
			}

			if tail == 0 {
				ans = append(ans, inline)
				inline = []int{}
				flag = 1
			}
		} else { // flag == 1, using stack, right to left
			top--
			node := stack[top]

			inline = append(inline, node.Val)

			if node.Right != nil {
				stk[tail] = node.Right
				tail++
			}

			if node.Left != nil {
				stk[tail] = node.Left
				tail++
			}

			if top == 0 {
				ans = append(ans, inline)
				inline = []int{}
				flag = 0
			}
		}
	}
	return ans
}

// @lc code=end

func testZigzagLevelOrder() {
	root := createTree()
	forder(root)
	println()

	ans := zigzagLevelOrder(root)
	println(ans)
}

func forder(node *TreeNode) {
	if node == nil {
		return
	}
	print(node.Val, " ")
	forder(node.Left)
	forder(node.Right)
}

func createTree() *TreeNode {
	var x int
	fmt.Scan(&x)
	if x == 101 {
		return nil
	}
	left := createTree()
	right := createTree()
	return &TreeNode{Val: x, Left: left, Right: right}
}
