package main

import "fmt"

func main() {
	ca := []int{1, 0, 2}
	//ca := []int{1, 2, 3}
	ans := candy(ca)
	fmt.Println(ans)
}

func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	length := len(ratings)
	increment := make([]int, length)

	inc := 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			increment[i] = inc
			inc++
		} else {
			inc = 1
		}
	}
	fmt.Println(increment)

	inc = 1
	for i := length - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			increment[i] = max(inc, increment[i])
			inc++
		} else {
			inc = 1
		}
	}

	sum := length
	for _, v := range increment {
		sum += v
	}
	return sum
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
