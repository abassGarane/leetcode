package main

import (
	"fmt"

	"github.com/abassGarane/leetcode/sliding"
)

func main() {
	s := "10101"
	k := 1

	res := sliding.CountKConstraintSubstrings(s, k)
	if res != 12 {
		fmt.Printf("expected 12 got %d\n", res)
	} else {
		println("got it")
		println(res)
	}

}
