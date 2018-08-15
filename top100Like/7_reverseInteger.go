package main

func reverse(x int) int {
	res, p := 0, x
	for p != 0 {
		res = res*10 + p%10
		p /= 10
	}
	return res
}
