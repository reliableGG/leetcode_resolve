package main

import "fmt"

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)

	for index, value := range nums {
		dict[value] = index
	}

	for i, v := range nums {
		k := target - v
		fmt.Println(k)
		if i2, exsit := dict[k]; i2 != i && exsit {
			return []int{i, i2}
		}
	}
	return nil
}

func main() {
	num := []int{3, 2, 4}
	fmt.Println(twoSum(num, 6))
}
