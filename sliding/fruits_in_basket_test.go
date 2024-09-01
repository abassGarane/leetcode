package sliding_test

import (
	"testing"

	"github.com/abassGarane/leetcode/sliding"
)

func TestTotalFruit(t *testing.T) {

	given := []int{1, 2, 1}
	expected := 3
	res := sliding.TotalFruit(given)
	if res != expected {
		t.Errorf("given %v expected %d but got %d", given, expected, res)
	}
}
