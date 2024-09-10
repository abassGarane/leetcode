package binarysearch

import (
	"math"
	"slices"
)

// Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas.
// The guards have gone and will come back in h hours.
// Koko can decide her bananas-per-hour eating speed of k.
// Each hour, she chooses some pile of bananas and eats k bananas from that pile.
// If the pile has less than k bananas, she eats all of them instead and will not
// eat any more bananas during this hour.
// Koko likes to eat slowly but still wants to finish eating all
// the bananas before the guards return.
// Return the minimum integer k such that she can eat all the bananas within h hours.

func MinEatingSpeed(piles []int, h int) int {
	left, right := 0, slices.Max(piles)
	res := right // max pile
	for left <= right {
		mid := (left + right) / 2
		total_time := 0.0
		for _, v := range piles {
			total_time += math.Ceil(float64(v) / float64(mid))
		}
		if total_time <= float64(h) {
			res = min(res, mid)
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return int(res)
}
