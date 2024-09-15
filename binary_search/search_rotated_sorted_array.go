package binarysearch

// There is an integer array nums sorted in ascending order (with distinct values).

// Prior to being passed to your function, nums is possibly rotated at
// an unknown pivot index k (1 <= k < nums.length) such that
// the resulting array
// is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
// For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

// Given the array nums after the possible rotation and an integer target,
// return the index of target if it is in nums, or -1 if it is not in nums.

// You must write an algorithm with O(log n) runtime complexity.

func SearchRotated(nums []int, target int) int {
	if len(nums) < 2 && nums[0] != target {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		}
		if nums[left] <= nums[mid] {
			if target > nums[mid] || target < nums[left] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if target < nums[mid] || target > nums[right] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}
func SearchAlt(nums []int, target int) int {
	k := 0
	if nums[0] > nums[len(nums)-1] {
		k = findK(nums, 0, len(nums)-1)
	}
	if target < nums[k] || target > nums[(len(nums)+k-1)%len(nums)] {
		return -1
	}
	return findTarget(nums, 0, len(nums)-1, k, target)
}

func findK(nums []int, start int, end int) int {
	if (end - start) <= 1 {
		return end
	}
	mid := (end + start) / 2

	if nums[mid] > nums[end] {
		return findK(nums, mid, end)
	} else {
		return findK(nums, start, mid)
	}
}

func findTarget(nums []int, start int, end int, k int, target int) int {
	mid := (end + start) / 2

	if nums[(start+k)%len(nums)] > target || nums[(end+k)%len(nums)] < target {
		return -1
	}
	if nums[(mid+k)%len(nums)] > target {
		return findTarget(nums, start, mid-1, k, target)
	} else if nums[(mid+k)%len(nums)] == target {
		return (mid + k) % len(nums)
	} else {
		return findTarget(nums, mid+1, end, k, target)
	}
}
