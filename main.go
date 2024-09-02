package main

import (
	"fmt"

	"github.com/abassGarane/leetcode/sliding"
)

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	expected := "BANC"
	res := sliding.MinWindow(s, t)
	if res != expected {
		fmt.Printf("expected %s got %s\n", expected, res)
	} else {
		println("got it")
		println(res)
	}

}
