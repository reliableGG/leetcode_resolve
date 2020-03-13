package main

import "fmt"

//import "sort"

func main() {
	fmt.Println("vim-go")
	//ca := []int{9, 0, 3, 5, 7}
	ca := []int{1, 2, 3}
	//sort.Slice(ca, func(i, j int) bool {
	//	return ca[i] < ca[j]
	//})
	res := permute(ca)
	fmt.Println(res)
}
func permute(nums []int) [][]int {
	tmp := []int{}
	result := [][]int{}
	backtrack(&result, tmp, nums)
	return result
}

func backtrack(list *[][]int, tmp []int, nums []int) {
	if len(tmp) == len(nums) {
		tt := append([]int{}, tmp...)
		*list = append(*list, tt)
	} else {
		for i := 0; i < len(nums); i++ {
			iFound := false
			for _, v := range tmp {
				if v == nums[i] {
					iFound = true
				}
			}
			if iFound {
				continue
			}
			tmp = append(tmp, nums[i])
			backtrack(list, tmp, nums)
			tmp = tmp[:len(tmp)-1]
		}
	}
}
