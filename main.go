package main

import (
	"log"

	binarysearch "github.com/abassGarane/leetcode/binary_search"
)

func main() {
	num1 := []int{1, 3}
	num2 := []int{2}
	medium := 2.00
	res := binarysearch.FindMedianSortedArrays(num1, num2)
	if res != medium {
		log.Printf("wanted %f got %f\n", medium, res)
	} else {
		log.Printf("medium found at %f\n", res)
	}
}
