package main

import "fmt"

func main() {
	ca := []int{1, 2, 1, 1, 1}
	ans := incre(ca)
	fmt.Println(ans)
}

func incre(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	//min := nums[0]
	min := 1<<31 - 1
	sec := 1<<31 - 1
	//sec := nums[0]
	for _, v := range nums {
		if v <= min {
			min = v
		} else if v <= sec {
			sec = v
		} else {
			return true
		}
	}
	return false
}
