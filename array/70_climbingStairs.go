package main

import "fmt"

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	two_step_before := 1
	one_step_before := 2
	all_step := 0
	for i := 2; i < n; i++ {
		all_step = two_step_before + one_step_before
		two_step_before = one_step_before
		one_step_before = all_step
	}
	return all_step
}

func main() {
	fmt.Println(climbStairs(8))
}
