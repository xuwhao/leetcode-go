package main

/*
 * @lc app=leetcode id=567 lang=golang
 *
 * [567] Permutation in String
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	s, p := s2, s1
	if len(s) < len(p) {
		return false
	}
	need, window, finish, i, j := map[byte]int{}, map[byte]int{}, 0, 0, 0

	for _, c := range p {
		need[byte(c)]++
	}

	for j = 0; j < len(p); j++ {
		c := s[j]
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				finish++
			}
		}
	}

	for j < len(s) {
		if finish == len(need) {
			return true
		}
		c := s[i]
		if _, ok := need[c]; ok {
			window[c]--
			if window[c]+1 == need[c] {
				finish--
			}
		}
		c = s[j]
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				finish++
			}
		}
		i, j = i+1, j+1
	}

	if finish == len(need) {
		return true
	}

	return false
}

// @lc code=end
