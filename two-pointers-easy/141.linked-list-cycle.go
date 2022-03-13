package main

/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	var p, q *ListNode = head, nil
	if p != nil {
		q = p.Next
	}
	flag := false
	for q != nil && p != q {
		p, q = p.Next, q.Next
		if q != nil {
			q = q.Next
		}
	}
	if p == q && q != nil {
		flag = true
	}
	return flag
}

// @lc code=end
