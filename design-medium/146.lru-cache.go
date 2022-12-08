package main

import "fmt"

/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// @lc code=start

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

type LRUCache struct {
	capacity int
	size     int
	cacheMap map[int]*Node
	head     *Node
	// tail     *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	head.Next = head
	head.Prev = head
	return LRUCache{capacity, 0, make(map[int]*Node), head}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cacheMap[key]; ok {
		this.moveNodeToEnd(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) moveNodeToEnd(node *Node) {
	// remove node from the list
	if node.Prev != nil && node.Next != nil {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}

	tail := this.head.Prev

	// Put node in the end of the list
	node.Next = tail.Next
	node.Prev = tail
	tail.Next.Prev = node
	tail.Next = node
}

func (this *LRUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}

	if node, ok := this.cacheMap[key]; ok {
		this.moveNodeToEnd(node)
		node.Val = value
		return
	}

	if this.size == this.capacity {
		tmp := this.head.Next

		tmp.Next.Prev = this.head
		this.head.Next = tmp.Next

		// help GC
		tmp.Next = nil
		tmp.Prev = nil

		delete(this.cacheMap, tmp.Key)
		this.size--
	}

	node := &Node{Key: key, Val: value}
	this.moveNodeToEnd(node)
	this.cacheMap[key] = node
	this.size++
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

func main() {
	var res int
	ans := []int{}
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1)    // cache is {1=1}
	lRUCache.Put(2, 2)    // cache is {1=1, 2=2}
	res = lRUCache.Get(1) // return 1
	ans = append(ans, res)
	lRUCache.Put(3, 3)    // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	res = lRUCache.Get(2) // returns -1 (not found)
	lRUCache.Put(4, 4)    // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	res = lRUCache.Get(1) // return -1 (not found)
	ans = append(ans, res)
	res = lRUCache.Get(3) // return 3
	ans = append(ans, res)
	res = lRUCache.Get(4) // return 4
	ans = append(ans, res)
	fmt.Println(ans)
}
