package stack

func LargestRectangleArea(heights []int) int {
	//---------
	// Scenarios
	//-----------
	// 1. Rectange can be extended if :-
	// - if the rects have the same height
	// - if the first rect is smaller than second rect
	//
	// 2. Rect can't be extended if :-
	// - if the first rect is larger than the second rectangle.

	// Every rect has a boundary where it can start and end
	// Elements can be extended backwards
	max_area := 0
	stack := [][2]int{}
	for i, v := range heights {
		if len(stack) == 0 {
			stack = append(stack, [2]int{i, v})
			continue
		}
		start := i
		for len(stack) > 0 && stack[len(stack)-1][1] > v {
			area := (i - stack[len(stack)-1][0]) * stack[len(stack)-1][1]
			max_area = max(max_area, area)
			start = stack[len(stack)-1][0]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, [2]int{start, v})
	}
	// Stack items not popped
	for _, v := range stack {
		area := (len(heights) - v[0]) * v[1]
		max_area = max(max_area, area)
	}
	return max_area
}
func LargestRectangleAreaAlt(heights []int) int {
	heights = append(heights, 0)
	stack := []int{}
	maxArea := 0

	for i, h := range heights {
		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]] // get last element in stack
			stack = stack[:len(stack)-1]           // remove last element from stack
			width := i                             // current start
			if len(stack) > 0 {                    // until stack is empty
				width -= stack[len(stack)-1] + 1 // last stack
			}
			maxArea = max(maxArea, height*width)
		}
		stack = append(stack, i)
	}

	return maxArea
}
