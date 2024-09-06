package stack

// Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

// Implement the MinStack class:

// MinStack() initializes the stack object.
// void push(int val) pushes the element val onto the stack.
// void pop() removes the element on the top of the stack.
// int top() gets the top element of the stack.
// int getMin() retrieves the minimum element in the stack.
// You must implement a solution with O(1) time complexity for each function.

// Stack -> last in first out
type MinStack struct {
	items []int
	min   []int
}

func Constructor() MinStack {
	return MinStack{
		items: []int{},
		min:   []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.items = append(this.items, val)
	if len(this.min) == 0 {
		this.min = append(this.min, val)
	} else {
		if this.min[len(this.min)-1] >= val {
			this.min = append(this.min, val)
		}
	}
}

func (this *MinStack) Pop() {
	items_top := this.items[len(this.items)-1]
	this.items = this.items[:len(this.items)-1]
	min_top := this.min[len(this.min)-1]
	if min_top == items_top {
		this.min = this.min[:len(this.min)-1]
	}
}

func (this *MinStack) Top() int {
	return this.items[len(this.items)-1]
}

func (this *MinStack) GetMin() int {

	return this.min[len(this.min)-1]
}
