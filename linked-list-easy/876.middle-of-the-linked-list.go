package main

import (
	"fmt"
	"leetcode-go/util"
)

/*
 * @lc app=leetcode id=876 lang=golang
 *
 * [876] Middle of the Linked List
 */
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	return slow
}

// @lc code=end

func testMiddleNode() {
	head := []int{4, 2}
	ans := util.List2Ints(middleNode(util.Ints2List(head)))
	fmt.Println(ans)
}
