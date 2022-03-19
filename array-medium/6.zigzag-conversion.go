package main

import "fmt"

/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] Zigzag Conversion
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	n, asc, cnt := len(s), true, 0
	order, str, ans := make([]int, n), make([][]byte, numRows), make([]byte, n)
	for i := 0; i < n; i++ {
		order[i] = cnt
		if cnt == numRows-1 {
			asc = false
		} else if cnt == 0 {
			asc = true
		}
		if asc {
			cnt++
		} else {
			cnt--
		}

	}
	for i := 0; i < n; i++ {
		str[order[i]] = append(str[order[i]], s[i])
	}
	cnt = 0
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(str[i]); j++ {
			ans[cnt] = str[i][j]
			cnt++
		}
	}
	return string(ans)
}

// @lc code=end
func testConvert() {
	s, n := "AB", 1
	fmt.Println(convert(s, n))
}
