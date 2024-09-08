package stack

import (
	"sort"
)

func CarFleet(target int, position []int, speed []int) int {
	f_stack := []float64{}
	zipped := zip(position, speed)

	sort.SliceStable(zipped, func(i, j int) bool {
		return zipped[i][0] > zipped[j][0]
	})
	for _, v := range zipped {
		speed := float64(target-v[0]) / float64(v[1])
		f_stack = append(f_stack, speed)
		len_s := len(f_stack)
		if len(f_stack) >= 2 && f_stack[len_s-1] <= f_stack[len_s-2] {
			f_stack = f_stack[:len_s-1]
		}
	}
	return len(f_stack)
}

func zip(position, speed []int) [][2]int {
	if len(position) != len(speed) || len(position) == 0 || len(speed) == 0 {
		return [][2]int{}
	}
	zipped := [][2]int{}
	for i := 0; i < len(position); i++ {
		zipped = append(zipped, [2]int{position[i], speed[i]})
	}
	return zipped
}

func CarFleetAlt(target int, position []int, speed []int) int {
	stack := make([]float32, target+1)
	var current float32 = 0
	res := 0

	for i, pos := range position {
		stack[pos] = float32(target-pos) / float32(speed[i])
	}

	for i := target; i >= 0; i-- {
		if stack[i] > current {
			current = stack[i]
			res++
		}
	}

	return res
}
