package main

import "fmt"

/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 */

// @lc code=start
// func firstUniqChar(s string) int {
// 	counter := map[rune]int{}
// 	for _, ch := range s {
// 		if cnt, exists := counter[ch]; exists {
// 			counter[ch] = cnt + 1
// 		} else {
// 			counter[ch] = 1
// 		}
// 	}
// 	for idx, ch := range s {
// 		if cnt, exists := counter[ch]; exists {
// 			if cnt == 1 {
// 				return idx
// 			}
// 		}
// 	}
// 	return -1
// }

func firstUniqChar(s string) int {
	result := make([]int, 26)
	for i := 0; i < len(s); i++ {
		result[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if result[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

// @lc code=end

func TestFirstUniqChar() {
	str := "loveleetcode"
	fmt.Println(firstUniqChar(str))
}
