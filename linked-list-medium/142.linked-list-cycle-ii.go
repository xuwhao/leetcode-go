package main

import (
	"fmt"
	"leetcode-go/util"
)

/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
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
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head.Next

	for fast != nil && slow != fast {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}

	if fast == nil {
		return nil
	}

	slow = head
	fast = fast.Next
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return fast
}

// @lc code=end

func testDetectCycle() {
	list := util.Ints2List([]int{1, 2})
	tail := list.GetNodeWith(2)
	head := list.GetNodeWith(1)
	tail.Next = head

	ans := detectCycle(list)
	if ans == nil {
		fmt.Println(-1)
	} else {
		fmt.Println(ans.Val)
	}

}
