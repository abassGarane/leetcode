package sliding

// Given an array of integers nums and an integer k, return the number of contiguous
// subarrays where the productuct of all the elements in the subarray is strictly less than k.
func NumSubarrayProductLessThanK(nums []int, k int) int {
	if len(nums) <= 1 || k <= 1 {
		return 0
	}
	result := 0
	product := 1
	right, left := 0, 0

	for right < len(nums) {
		product *= nums[right]
		for product >= k {
			product /= nums[left]
			left++
		}
		result += right - left + 1
		right++
	}
	return result
}
