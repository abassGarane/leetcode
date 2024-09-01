package sliding

func MaxSlidingWindow(nums []int, k int) []int {
	deque := []int{}
	res := []int{}
	left, right := 0, 0
	for right < len(nums) {
		for len(deque) > 0 && nums[right] > nums[deque[len(deque)-1]] {
			// pop if smaller elem in deque
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, right)
		if left > deque[0] {
			deque = deque[1:]
		}
		if right+1 >= k {
			res = append(res, nums[deque[0]])
			left++
		}
		right++
	}
	return res
}
