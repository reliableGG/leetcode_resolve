package main

import "fmt"

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	ans := canCompleteCircuit(gas, cost)
	fmt.Println(ans)
}

func canCompleteCircuit(gas []int, cost []int) int {

	sum := 0
	total := 0
	j := 0

	for idx, _ := range gas {
		sum += gas[idx] - cost[idx]
		total += gas[idx] - cost[idx]
		if sum < 0 {
			j = idx + 1
			sum = 0
		}
	}

	if total > 0 {
		return j
	} else {
		return -1
	}

}
