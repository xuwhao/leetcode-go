package main

/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func count(head *ListNode) int {
	p, cnt := head, 0

	for p != nil {
		cnt++
		p = p.Next
	}

	return cnt
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	cntA, cntB := count(headA), count(headB)

	if cntA > cntB {
		headA, headB = headB, headA
		cntA, cntB = cntB, cntA
	}

	k := cntB - cntA
	p, q := headA, headB
	for ; k > 0; k-- {
		q = q.Next
	}

	for p != q {
		p = p.Next
		q = q.Next
	}

	return p
}

// @lc code=end
