package main

import (
	"fmt"
	"leetcode-go/util"
)

/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func sortList(head *ListNode) *ListNode {
// 	var sort func(head *ListNode, tail *ListNode) *ListNode
// 	sort = func(head *ListNode, tail *ListNode) *ListNode {
// 		if head == nil {
// 			return nil
// 		}
// 		if head.Next == tail {
// 			head.Next = nil
// 			return head
// 		}
// 		slow, fast := head, head.Next
// 		for fast != tail {
// 			slow = slow.Next
// 			fast = fast.Next
// 			if fast != tail {
// 				fast = fast.Next
// 			}
// 		}
// 		return mergeTwoLists(sort(head, slow), sort(slow, tail))
// 	}

// 	return sort(head, nil)
// }

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	if fast.Next != nil {
		for fast != nil {
			slow = slow.Next
			fast = fast.Next
			if fast != nil {
				fast = fast.Next
			}
		}
	}
	fast = slow.Next
	slow.Next = nil

	return mergeTwoLists(sortList(head), sortList(fast))
}

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

func testSortList() {
	head := []int{-1, 5, 3, 4, 0}
	ans := util.List2Ints(sortList(util.Ints2List(head)))
	fmt.Println(ans)
}
