package main

import "fmt"

/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */

// @lc code=start
// func findKthLargest(nums []int, k int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}

// 	ans := -1
// 	var search func(left, right, k int)
// 	search = func(left, right, k int) {
// 		if left == right {
// 			ans = left
// 			return
// 		}

// 		if left > right {
// 			return
// 		}

// 		base, i, j := nums[left], left, right
// 		for i < j {
// 			for nums[j] <= base && i < j {
// 				j--
// 			}
// 			for nums[i] >= base && i < j {
// 				i++
// 			}

// 			temp := nums[i]
// 			nums[i] = nums[j]
// 			nums[j] = temp
// 		}
// 		nums[left] = nums[i]
// 		nums[i] = base

// 		if i == left+k-1 {
// 			ans = i
// 			return
// 		}

// 		if i > left+k-1 {
// 			search(left, i-1, k)
// 		} else {
// 			search(i+1, right, k-(i-left+1))
// 		}
// 	}
// 	search(0, len(nums)-1, k)
// 	return nums[ans]
// }

func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	// for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
	// 	nums[0], nums[i] = nums[i], nums[0]
	// 	heapSize--
	// 	maxHeapify(nums, 0, heapSize)
	// }
	for ; k > 1; k-- {
		i := heapSize - 1
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeap(a []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(a, i, heapSize)
	}
}

func maxHeapify(a []int, i, heapSize int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < heapSize && a[l] > a[largest] {
		largest = l
	}
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, heapSize)
	}
}

// @lc code=end

func main() {
	nums := []int{7, 6, 5, 4, 3, 2, 1}
	fmt.Println(findKthLargest(nums, 5))
}
