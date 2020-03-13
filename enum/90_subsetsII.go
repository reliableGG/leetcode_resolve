package main

import "fmt"

//import "sort"

func main() {
	fmt.Println("vim-go")
	ca := []int{1, 2, 3}
	res := subsetsWithDump(ca)
	fmt.Println(res)
	fmt.Println(len(res))
}

func subsetsWithDump(nums []int) [][]int {
	result := [][]int{{}}

	for _, v := range nums {
		for _, subset := range result {
			result = append(result, append(subset, v))
		}
	}
	return result
}
