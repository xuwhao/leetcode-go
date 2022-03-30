package main

/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	ans := &ListNode{Val: 0, Next: head}
	h, p, q := ans, ans, ans
	k := 1
	q = p.Next

	for k < left {
		p = q
		q = q.Next
		k++
	}
	h = p
	tail := q

	for k < right+1 {
		t := q.Next
		q.Next = h.Next
		h.Next = q
		q = t
		k++
	}

	if tail != nil {
		tail.Next = q
	}

	return ans.Next
}

// @lc code=end
