package sliding

import (
	"slices"
)

func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	if len(nums) <= k {
		return []int{slices.Max(nums)}
	}
	window, res := []int{}, []int{}
	for right := 0; right < len(nums); right++ {
		window = append(window, nums[right])
		if len(window) < k {
			continue
		}
		for len(window) > k {
			window = window[1:]
		}
		window_max := slices.Max(window)
		res = append(res, window_max)
	}
	return res
}
