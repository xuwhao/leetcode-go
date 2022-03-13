package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
type IntSlice [][]int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Less(i, j int) bool { return s[i][0] < s[j][0] }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n == 1 {
		return intervals
	}
	sort.Sort(IntSlice(intervals))
	i, j, ans := 0, 1, [][]int{}
	ans = append(ans, intervals[0])
	for j < len(intervals) {
		if ans[i][1] >= intervals[j][1] {
			//j++
		} else if ans[i][1] >= intervals[j][0] {
			merged := []int{ans[i][0], intervals[j][1]}
			ans = append(ans[:i], merged)
		} else if ans[i][1] < intervals[j][1] {
			ans = append(ans, intervals[j])
			i++
		}
		j++
	}
	// i, j, ans, cnt := 0, 1, make([][]int, len(intervals)), n
	// ans[i] = intervals[i]
	// for j < n {
	// 	if ans[i][1] >= intervals[j][1] {
	// 		//j++
	// 		cnt-- // 将两个数组合并成一个，数组数量 -1
	// 	} else if ans[i][1] >= intervals[j][0] {
	// 		merged := []int{ans[i][0], intervals[j][1]}
	// 		ans[i] = merged
	// 		cnt-- // 将两个数组合并成一个，数组数量 -1
	// 	} else {
	// 		i++
	// 		ans[i] = intervals[j]
	// 	}
	// 	j++
	// }
	// return ans[0:cnt]
	return ans
}

// @lc code=end

func testMerge() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))
}
