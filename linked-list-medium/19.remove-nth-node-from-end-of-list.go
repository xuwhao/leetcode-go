package main

import "leetcode-go/util"

/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	head = &ListNode{Val: 0, Next: head}
	p, q, t := head, head.Next, head.Next
	for ; n > 0; n-- {
		t = t.Next
	}
	for t != nil {
		p, q, t = p.Next, q.Next, t.Next
	}
	p.Next = q.Next
	return head.Next
}

// @lc code=end
