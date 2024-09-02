package sliding

import "math"

// Given two strings s and t of lengths m and n respectively, return the minimum window
// substring of s such that every character in t (including duplicates) is included in the window.
// If there is no such substring, return the empty string "".
// The testcases will be generated such that the answer is unique.
func MinWindow(s string, t string) string {
	var answer = ""
	if len(s) < len(t) {
		return answer
	}
	window := [128]int{}
	for _, c := range t {
		window[c]++
	}
	min_length := math.MaxInt64
	left, right := 0, 0
	counter := 0
	for right < len(s) {
		if left >= len(s) {
			return answer
		}
		curr := s[right]
		window[curr]--
		if window[curr] >= 0 {
			counter++
		}
		for counter == len(t) {
			curr_min := right - left + 1
			if curr_min < min_length {
				min_length = curr_min
				answer = s[left : right+1]
			}
			v := window[s[left]]
			if v >= 0 {
				counter--
			}
			left++
		}
		right++
	}
	return answer
}
