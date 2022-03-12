package main

/*
 * @lc app=leetcode id=138 lang=golang
 *
 * [138] Copy List with Random Pointer
 */
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// @lc code=start

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	p, np := head.Next, &Node{Val: head.Val}
	np.Next = nil
	clone := np

	// clone attribute 'Next' only
	for p != nil {
		cur := &Node{Val: p.Val}
		np.Next = cur
		p = p.Next
		np = np.Next
	}

	// construct attribute 'Random'
	p, np = head, clone
	for p != nil {
		target := p.Random
		if target == nil {
			np.Random = nil
		} else {
			rh, rc := head, clone
			for rh != nil {
				if target == rh {
					np.Random = rc
					break
				}
				rh = rh.Next
				rc = rc.Next
			}
		}
		p = p.Next
		np = np.Next
	}
	return clone
}

// @lc code=end

func testopyRandomList() {

}
