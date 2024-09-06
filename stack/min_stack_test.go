package stack_test

import (
	"testing"

	"github.com/abassGarane/leetcode/stack"
)

func TestConstructor(t *testing.T) {
	min_stack := stack.Constructor()
	if min_stack.GetMin() != 0 {
		t.Fail()
	}
}
