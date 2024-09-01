package sliding

// You are given a binary string s and an integer k.
// A binary string satisfies the k-constraint if either of the following conditions holds:
// The number of 0's in the string is at most k.
// The number of 1's in the string is at most k.
// Return an integer denoting the number of
// substrings of s that satisfy the k-constraint.
func CountKConstraintSubstrings(s string, k int) int {
	result := 0
	ones := 0
	zeros := 0
	for left, right := 0, 0; right < len(s); right++ {
		if s[right] == '0' {
			zeros++
		} else {
			ones++
		}
		for ones > k && zeros > k {
			if s[left] == '0' {
				zeros--
			} else {
				ones--
			}
			left++
		}
		result += right - left + 1
	}
	return result

}
