package main

import "fmt"

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start

/*
  BRUTE FORCE
*/
// func maxArea(height []int) int {
// 	n, ans := len(height), 0
// 	for i := 0; i < n; i++ {
// 		for j := i; j < n; j++ {
// 			a := area(j-i, height[i], height[j])
// 			if a > ans {
// 				ans = a
// 			}
// 		}
// 	}
// 	return ans
// }

/*
	这题的问题在滑动窗口到底应该怎么滑，怎么判断坐标移动后面积会增大。

	area = min(height[i], height[j]) * (j-i);

	i 是左指针， j是右指针。 一开始令 j-i 最大，也就是 j0 = len(height)-1， i = 0

	不失一般性的，假设 height[i] < height[j0].

	这个时候如果我们移动高的那一侧，也就是说往左移动指针 j,  有两种情况：

	1. 移动之后 height[i] < height[j1]

	计算面积的时候仍然用的是 height[i], (j-i) 还变小了，面积肯定变小。

		(before) area1 = min(height[i], height[j0]) = height[i] * (j0-i);

		(after)  area2 = min(height[i], height[j1]) = height[i] * (j1-i) ;

		j1 < j0 ==> j1-i < j0-i ==> area2 < area1.

		2. 移动之后 height[i] >= height[j1].

		移动之前的面积

		(before) area1 = min(height[i], height[j0]) = height[i] * (j0-i);

		考虑下现在的面积。

		(after)  area2 = min(height[i], height[j1]) = height[j1] * (j1-i) ;

		由于 height[i] >= height[j1], (j0-i) > (j1-i), 所以 area2 < area1.


	所以我们发现移动高侧是没有任何意义的！只有移动矮的那一侧，面积才会有变大的可能性！

	那么算法就很简单，只需移动矮的那一侧，如果左边矮, i++, 右边矮, j--. 直到两个指针相撞，则遍历就结束了。
*/

func maxArea(height []int) int {
	ans := 0
	for i, j := 0, len(height)-1; i < j; {
		temp := area(j-i, height[i], height[j])
		if temp > ans {
			ans = temp
		}
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return ans
}

func area(dis int, y1 int, y2 int) int {
	return min(y1, y2) * dis
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

// @lc code=end

func main() {
	height := []int{1, 2, 1}
	fmt.Println(maxArea(height))
}
