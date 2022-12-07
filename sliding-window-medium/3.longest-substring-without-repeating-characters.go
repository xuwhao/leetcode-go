package main

import "fmt"

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start

// func slide(s string, start int, windowSize int) int {
// 	ch := s[start+windowSize]
// 	for i := start; i < start+windowSize; i++ {
// 		if s[i] == ch {
// 			return i + 1
// 		}
// 	}
// 	return start
// }

// func isRepeat(s string, start int, windowSize int) bool {
// 	counter := map[int]int{}
// 	for i := start; i < start+windowSize; i++ {
// 		if counter[int(s[i]-'a')] == 1 {
// 			return true
// 		}
// 		counter[int(s[i]-'a')] += 1
// 	}
// 	return false
// }

// func lengthOfLongestSubstring(s string) int {

// 	if s == "" {
// 		return 0
// 	}

// 	windowSize := 1
// 	start := 0
// 	counter := map[int]int{}

// 	counter[int(s[start]-'a')] = 1
// 	for start+windowSize < len(s) {
// 		ch := s[start+windowSize]
// 		if counter[int(ch-'a')] == 0 {
// 			windowSize += 1
// 			counter[int(ch-'a')] += 1
// 		} else {
// 			start = slide(s, start, windowSize)
// 			for start+windowSize < len(s) && isRepeat(s, start, windowSize) {
// 				start += 1
// 			}
// 			if start+windowSize >= len(s) {
// 				break
// 			}
// 			counter = map[int]int{}
// 			for i := start; i < start+windowSize; i++ {
// 				counter[int(s[i]-'a')] = 1
// 			}
// 		}

// 	}
// 	return windowSize

// }

// func lengthOfLongestSubstring(s string) int {
// 	if len(s) == 0 {
// 		return 0
// 	}
// 	var freq [256]int
// 	result, left, right := 0, 0, -1

// 	for left < len(s) {
// 		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
// 			freq[s[right+1]-'a']++
// 			right++
// 		} else {
// 			freq[s[left]-'a']--
// 			left++
// 		}
// 		result = max(result, right-left+1)
// 	}
// 	return result
// }

// func max(a int, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	ans := 1
	left, right := 0, 1
	flag := [256]int{0}

	// s[0] 进入窗口，出现次数 + 1
	flag[s[left]-'a']++

	for left < len(s) {

		if right == len(s) {
			break
		}

		if flag[s[right]-'a'] != 0 {
			flag[s[left]-'a']--
			left++
		} else {
			flag[s[right]-'a']++
			right++
		}

		if right-left > ans {
			ans = right - left
		}

	}

	return ans

}

// @lc code=end

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
