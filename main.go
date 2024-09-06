package main

import (
	"fmt"

	"github.com/abassGarane/leetcode/stack"
)

func main() {

	s := "()[]{}"
	res := stack.IsValid(s)
	if res != true {
		fmt.Println("wrong answer dude")
	} else {
		println("right answer dude")
	}
}
