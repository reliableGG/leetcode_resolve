package main

import "fmt"

func main() {
	ca := []int{3, 2, 3}
	ans := majorityElement(ca)
	fmt.Println(ans)
}

func majorityElement(nums []int) int {
	m := map[int]int{}
	for _, n := range nums {
		//if _, exist := m[n]; !exist {
		//    m[n] = 0
		//} else {
		//    m[n]+=1
		//}
		m[n] += 1
	}

	max := 0
	element := 0
	for value, count := range m {
		if count > max {
			max = count
			element = value
		}
	}

	return element
}
