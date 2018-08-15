package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}
	res, p := 0, x
	for p != 0 {
		res = res*10 + p%10
		p /= 10
	}
	return res
}

func main() {
	fmt.Println(math.MaxInt32)
	res := reverse(1534236469)
	fmt.Println(res)
}
