package main

import "fmt"

import "sort"

func main() {
	fmt.Println("vim-go")
	ca := []int{9, 0, 3, 5, 7}
	//ca := []int{1, 2, 3}
	//sort.Slice(ca, func(i, j int) bool {
	//	return ca[i] < ca[j]
	//})
	res := subsets(ca)
	fmt.Println(res)
}

/*func subsets(nums []int) [][]int {*/
//result := [][]int{[]int{}}

//for _, v := range nums {
//fmt.Println(v)
//for _, subset := range result {
////fmt.Println(subset)
//tmp := append(subset, v)
//result = append(result, tmp)
//}
//fmt.Println("*****")
//fmt.Println(result)
//}
//return result
/*}*/

func subsets(nums []int) [][]int {
	result := [][]int{}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	tmp := make([]int, 0)
	backtrack(&result, tmp, nums, 0)
	return result
}

func backtrack(list *[][]int, tmp []int, nums []int, start int) {
	fmt.Println(tmp)
	tt := append([]int{}, tmp...)
	*list = append(*list, tt)
	for i := start; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		backtrack(list, tmp, nums, i+1)
		tmp = tmp[:len(tmp)-1]
	}
}

func subsets(nums []int) [][]int {
	subsets := [][]int{[]int{}}
	for _, num := range nums {
		for _, subset := range subsets {
			newSubset := append([]int{}, subset...)
			newSubset = append(newSubset, num)
			subsets = append(subsets, newSubset)
		}
	}
	return subsets
}
