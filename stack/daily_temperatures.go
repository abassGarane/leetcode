package stack

func DailyTemperatures(temperatures []int) []int {
	t_len := len(temperatures)
	warm_temps := make([]int, t_len)
	t_stack := make([][2]int, 0)
	for i, t := range temperatures {
		for len(t_stack) > 0 && t > t_stack[len(t_stack)-1][0] {
			diff := i - t_stack[len(t_stack)-1][1]
			warm_temps[t_stack[len(t_stack)-1][1]] = diff
			t_stack = t_stack[:len(t_stack)-1]
		}
		t_stack = append(t_stack, [2]int{t, i})
	}
	return warm_temps
}
func DailyTemperaturesAlt(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := []int{}
	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			res[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}
