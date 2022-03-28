package main

import "leetcode-go/util"

/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
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
		} else {
			if list1.Val > list2.Val {
				list1, list2 = list2, list1
			}
			t1 := list1.Next
			list1.Next = nil
			tail.Next = list1
			tail = tail.Next
			list1 = t1
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

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)

	if n == 0 {
		return nil
	}

	if n == 1 {
		return lists[0]
	}

	if n == 2 {
		return mergeTwoLists(lists[0], lists[1])
	}

	mid := n / 2
	return mergeTwoLists(mergeKLists(lists[:mid]), mergeKLists(lists[mid:n]))
}

// @lc code=end
