package main

import "fmt"

/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */

// @lc code=start

// 单调栈想法，有待改进，不用每次都往后找最接近 height[left] 的,直接填平坑
// func trap(height []int) int {
// 	ans, left, right, n := 0, 0, 0, len(height)
// 	height = append(height, -1)
// 	for right < n {
// 		j, maxIdx, minus, temp := right+1, n, 0, 0
// 		for j < n {
// 			if height[j] >= height[left] {
// 				maxIdx = j
// 				minus = temp
// 				break
// 			}
// 			if height[j] > height[maxIdx] {
// 				maxIdx = j
// 				minus = temp
// 			}
// 			temp += height[j]
// 			j++
// 		}
// 		if maxIdx == n {
// 			break
// 		}
// 		min := Min(height[left], height[maxIdx])
// 		ans += min*(maxIdx-left-1) - minus
// 		//ans += area
// 		left, right = maxIdx, maxIdx
// 	}

// 	return ans
// }

// func Min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }

func trap(height []int) int {
	ans, left, right := 0, 0, len(height)-1
	for left != right {
		i, j := left, right
		if left < right && height[left] <= height[right] { // 左指针小或相等，动左指针
			i = left + 1
			biasLeft := 0
			for i < right && height[i] < height[right] { // 向右探测
				if height[i] > height[left] {
					break
				}
				biasLeft += height[i]
				i++
			}
			ans += (i-left-1)*height[left] - biasLeft
		}
		if left < right && height[right] < height[left] {
			j = right - 1
			biasRight := 0
			for j > left && height[j] < height[left] {
				if height[j] > height[right] {
					break
				}
				biasRight += height[j]
				j--
			}
			ans += (right-j-1)*height[right] - biasRight
		}
		left, right = i, j
	}
	return ans
}

// @lc code=end
func testTrap() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}
