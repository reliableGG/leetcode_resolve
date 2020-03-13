package main

import "fmt"

func main() {
}

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}
	for idx, v := range nums {
		if _, exist := m[v]; !exist {
			m[v] = idx
		} else {
			if idx-m[v] <= k {
				return true
			} else {
				m[v] = idx
			}
		}
	}
	return false
}
