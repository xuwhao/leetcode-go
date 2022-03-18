package main

import (
	"fmt"
	"leetcode-go/util"
)

/*
 * @lc app=leetcode id=1721 lang=golang
 *
 * [1721] Swapping Nodes in a Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	head = &ListNode{Val: 0, Next: head}
	p, q, cnt := head, head.Next, 1
	for cnt < k && q != nil {
		p = p.Next
		q = q.Next
		cnt++
	}

	r, s, t := head, head.Next, q.Next
	// cnt = 0
	// for cnt < k {
	// 	t = t.Next
	// 	cnt++
	// }

	for t != nil {
		t = t.Next
		s = s.Next
		r = r.Next
	}

	// sn := s.Next
	// qn := q.Next

	// r.Next = q
	// q.Next = sn

	// p.Next = s
	// s.Next = qn
	q.Val, s.Val = s.Val, q.Val
	return head.Next
}

// @lc code=end

func testSwapNodes() {
	head := []int{7, 9, 6, 6, 7, 8, 3, 0, 9, 5}
	k := 5
	input := util.Ints2List(head)
	ans := util.List2Ints(swapNodes(input, k))
	fmt.Println(ans)
}
