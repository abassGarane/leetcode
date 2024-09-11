package binarysearch

// Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:

// [4,5,6,7,0,1,2] if it was rotated 4 times.
// [0,1,2,4,5,6,7] if it was rotated 7 times.
// Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

// Given the sorted rotated array nums of unique elements, return the minimum element of this array.

// You must write an algorithm that runs in O(log n) time.
func FindMin(nums []int) int {
	res := nums[0]
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left] < nums[right] {
			res = min(res, nums[left])
			break
		}
		mid := (left + right) / 2
		res = min(res, nums[mid])
		if nums[mid] >= nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}

func FindMinAlt(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	lo, hi := 0, n-1
	if nums[lo] < nums[hi] {
		return nums[lo]
	}

	for lo <= hi {
		mid := (lo + hi) / 2
		if mid+1 < n && nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if mid > 0 && nums[mid] < nums[mid-1] {
			return nums[mid]
		}

		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return -1

}
