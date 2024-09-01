package sliding

import (
	"slices"
)

func MaxSlidingWindow(nums []int, k int) []int {
	left := 0
	if len(nums) == 0 {
		return []int{}
	}
	if len(nums) <= k {
		return []int{slices.Max(nums)}
	}
	res := []int{}

	for right := k - 1; right < len(nums); right++ {
		for right-left+1 > k {
			left++
		}
		window_max := slices.Max(nums[left : right+1])
		res = append(res, window_max)
	}
	return res
}
