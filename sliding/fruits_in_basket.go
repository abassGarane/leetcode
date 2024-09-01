package sliding

// You are visiting a farm that has a single row
// of fruit trees arranged from left to right.
// The trees are represented by an integer array fruits where
// fruits[i] is the type of fruit the ith tree produces

func TotalFruit(fruits []int) int {
	start, end, res := 0, 0, 0
	freq := map[int]int{}
	for end < len(fruits) {
		freq[fruits[end]] += 1

		for len(freq) > 2 {
			freq[fruits[start]] -= 1
			if freq[fruits[start]] == 0 {
				delete(freq, fruits[start])
			}
			start++
		}
		res = max(res, end-start+1)
		end++
	}
	return res
}
