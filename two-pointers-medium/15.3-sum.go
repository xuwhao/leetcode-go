package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */
/*
	几个剪枝的 tricks.
	1. 排序后的结果一定是升序排列，所以不用考虑诸如 [-1,0,1] 和 [1,0,-1] 这样子的重复，在搜索 [1, *, *] 的时候只需要从 1 后面的数开始。
	2. 假设目标和为 target, 则排序后大于 target 的区间都不用搜索，升序排列下，后面的数怎么加都不可能比 target 小
	3. 找到一组后要把所有和这组一样的情况都消去。因为是升序排列，一样的数肯定在一起，所以当找到一组数之后直接跳过所有一样的三个数即可。
*/
// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	const target = 0
	for i := 0; i < len(nums); i++ {
		// 情况 2
		if nums[i] > target {
			break
		}
		// 情况 3, [a, b, c] 跳过 a 的重复
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 情况 1, [a, b, c] 确保 a < b < c
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				// 情况 3, [a, b, c] 跳过 b, c 的重复
				t1, t2 := nums[j], nums[k]
				for nums[j] == t1 && j < k {
					j++
				}
				for nums[k] == t2 && j < k {
					k--
				}
			}
			if sum < target {
				j++
			} else if sum > target {
				k--
			}
		}
	}
	return ans
}

// @lc code=end

// func main() {
// nums := []int{0, 0, 0, 0}
// fmt.Println(threeSum(nums))
// }
