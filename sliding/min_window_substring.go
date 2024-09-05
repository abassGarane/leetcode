package sliding

import "math"

// Given two strings s and t of lengths m and n respectively, return the minimum window
// substring of s such that every character in t (including duplicates) is included in the window.
// If there is no such substring, return the empty string "".
// The testcases will be generated such that the answer is unique.
func MinWindow(s string, t string) string {
	// Create a map to count the frequency of characters in t
	mapT := make([]int, 128)
	for _, c := range t {
		mapT[c]++
	}

	// Initialize variables
	counter := len(t)
	begin, end := 0, 0
	d := math.MaxInt
	head := 0

	// Iterate through the string s
	for end < len(s) {
		if mapT[s[end]] > 0 {
			counter--
		}
		mapT[s[end]]--
		end++

		// Shrink the window from the left if all characters are matched
		for counter == 0 {
			if end-begin < d {
				d = end - begin
				head = begin
			}
			// add the left elem to the freq map
			mapT[s[begin]]++
			// if positive, increase the counter because we just lost an element that we need to be in substr
			if mapT[s[begin]] > 0 {
				counter++
			}
			begin++
		}
	}

	if d == math.MaxInt {
		return ""
	}
	return s[head : head+d]
}
