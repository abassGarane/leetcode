package binarysearch

import (
	"math"
)

// Given two sorted arrays nums1 and nums2 of size m and n respectively,
// return the median of the two sorted arrays.

// The overall run time complexity should be O(log (m+n)).

// If the array is merged and sorted, if it is partitioned at the medium,
// then, max(of elements left of medium) <= min(elements right of medium)
//
// Partitioning
// if the first pointer p1 is at the 3rd element, and the medium of array is to be at 5,
// then the second pointer on the second array must be at the 2nd element.
//
// correct partition
// a partitioning is only correct if and only if the max(elements on left partition) is
// less than min(elements on the right partition)
// ie max(left elements) <= min(right elements)
//
// since elements in the same array are already sorted, do cross checking
// ie max(elements on the left of arr 1 ) <= min(elements on the right of arr2)
// since it is already sorted, the min of right is immediately after the medium
// if (n+m) is odd, the medium if the maximum of the last elements on the left partition
// ie if the nums1 = [1,2,3,4 | 5,6,7] and nums2 is at [1,2,3|4,5,6,6,8], then
// the medium is `min(4,3)` which is (3)
//
// if n+m is even, the medium is the max(last 2 on left) + min(the first two the right) / 2
//
//*Boundary cases*
//
// if one array is smaller than the other ie [1,2,3] && [2,5,6,7,8], during the comparisons of the p1 and p2
// if the nums1[p+1] is out of bounds, replace it with +infinity or maxint

// if p1 is int 0 element of the nums1, then nums1[p1] won't exist.
// replace it with -infinity or -maxint
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		return FindMedianSortedArrays(nums2, nums1)
	}
	low, high := 0, m
	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (m+n+1)/2 - partitionX

		maxLeftX := math.MinInt
		if partitionX != 0 {
			maxLeftX = nums1[partitionX-1]
		}

		minRightX := math.MaxInt
		if partitionX != m {
			minRightX = nums1[partitionX]
		}

		maxLeftY := math.MinInt
		if partitionY != 0 {
			maxLeftY = nums2[partitionY-1]
		}

		minRightY := math.MaxInt
		if partitionY != n {
			minRightY = nums2[partitionY]
		}
		// medium found
		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			if (m+n)%2 == 0 {
				return (float64(max(maxLeftX, maxLeftY)) + float64(min(minRightX, minRightY))) / 2
			} else {
				return float64(max(maxLeftX, maxLeftY))
			}
		} else if maxLeftX > minRightY { // medium not found
			high = partitionX - 1
		} else {
			low = partitionX + 1
		}
	}
	return 0.0
}
