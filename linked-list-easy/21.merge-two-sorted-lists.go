package main

//type ListNode = util.ListNode

/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{Val: 0, Next: nil}
	tail := head
	for list1 != nil && list2 != nil {
		if list1.Val == list2.Val {
			t1, t2 := list1.Next, list2.Next
			list1.Next, list2.Next = list2, nil
			tail.Next = list1
			tail = tail.Next.Next
			list1, list2 = t1, t2
		} else if list1.Val < list2.Val {
			t1 := list1.Next
			list1.Next = nil
			tail.Next = list1
			tail = tail.Next
			list1 = t1
		} else {
			t1 := list2.Next
			list2.Next = nil
			tail.Next = list2
			tail = tail.Next
			list2 = t1
		}
	}
	if list1 == nil && list2 != nil {
		tail.Next = list2
	}
	if list1 != nil && list2 == nil {
		tail.Next = list1
	}
	return head.Next
}

// @lc code=end
