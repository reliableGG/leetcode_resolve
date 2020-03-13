package main

import "fmt"

func main() {
	cas := []int{1, 2, 3, 3, 5}
	ans := searchRange(cas, 3)
	fmt.Println(ans)
}

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	n := len(nums)
	i, j := 0, n-1
	for i < j {
		mid := (i + j) / 2
		if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid
		}
	}
	if nums[i] != target {
		return res
	} else {
		res[0] = i
	}

	i, j = 0, n-1
	for i < j {
		mid := (i+j)/2 + 1
		if nums[mid] > target {
			j = mid - 1
		} else {
			i = mid
		}
	}
	res[1] = j
	return res

}
