package main

import "leetcode-go/util"

/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */

type ListNode = util.ListNode

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	h := &ListNode{Val: 0, Next: nil}
	for head != nil {
		t := head.Next
		head.Next = h.Next
		h.Next = head
		head = t
	}
	return h.Next
}

// @lc code=end
