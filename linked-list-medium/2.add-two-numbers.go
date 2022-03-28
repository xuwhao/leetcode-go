package main

import (
	"fmt"
	"leetcode-go/util"
)

/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	t, p, q := &ListNode{Val: 0, Next: nil}, l1, l2
	head := &ListNode{Val: 0, Next: t}
	r := head
	carry := 0

	for p != nil || q != nil || carry != 0 {
		nump, numq := 0, 0

		if p != nil {
			nump = p.Val
		}
		if q != nil {
			numq = q.Val
		}

		cur := nump + numq + carry
		carry = cur / 10

		t.Val = cur % 10
		t.Next = &ListNode{}

		r = r.Next
		t = t.Next

		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}

	if carry != 0 {
		t.Val = carry
		t.Next = nil
	} else {
		r.Next = nil
	}

	return head.Next
}

// @lc code=end
func testAddTwoNumbers() {
	l1 := []int{0, 1}
	l2 := []int{0, 9, 9, 9, 9, 9, 9, 9, 9}

	ans := addTwoNumbers(util.Ints2List(l1), util.Ints2List(l2))
	fmt.Println(util.List2Ints(ans))
}
