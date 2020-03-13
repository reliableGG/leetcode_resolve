package main

import "fmt"

func main() {
	a := []int{2, 7, 11, 5}
	res := twoSum(a, 9)
	fmt.Println(res)
}
func twoSum(nums []int, target int) []int {
	result := []int{}
	if len(nums) == 0 {
		return result
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				//result[0], result[1] = i, j
				result = append(result, i, j)
				return result
			}
		}

	}
	return result

}
