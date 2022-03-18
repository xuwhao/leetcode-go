package main

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */

// @lc code=start
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	need, window := map[byte]int{}, map[byte]int{}
	l, r, cnt, left, right := 0, 0, 0, 0, math.MaxInt32

	for _, c := range t {
		need[byte(c)]++
	}

	for r < len(s) {
		// 字符进入窗口
		c := s[r]
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] { // 某种字符集齐了
				cnt++
			}
		}
		r++ // 扩展窗口

		// 收缩窗口
		if cnt == len(need) { // 所有字符都集齐了
			for l < r { // 收缩窗口
				remove := s[l]
				if (r - l) <= (right - left) {
					left, right = l, r
				}
				l++                            // 抛掉字符 remove
				if _, ok := need[remove]; ok { // 如果 remove 在 t 中，更新窗口内的数据
					window[remove]--
					if window[remove] < need[remove] {
						cnt--
						break
					}
				}

			}
		}
	}

	if right != math.MaxInt32 { // 完全不符合，防止没更新过默认输出
		return s[left:right]
	}
	return ""
}

// @lc code=end

func testMinWindow() {
	s, t := "abc", "abc"
	fmt.Println(minWindow(s, t))
}
