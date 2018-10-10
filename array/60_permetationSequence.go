package main

import "fmt"

func getNthEle(n int, used []bool) int {
	if n < 0 {
		return nil
	}

	i := -1
	for n > 0 {
		i++
		if nums[i] == false {
			n--
		}
	}
	return i
}

func factorial(n int) []int {
	if n < 0 {
		return nil
	}

	res := make([]int)
	res[0] = 1
	for i := 1; i <= n; i++ {
		res[i] = res[i-1] * i
	}
	return res
}

func getPermutation(n int, k int) string {
	fact := factorial(n - 1)
	out := make([]byte, 0, n)
	used := make([]bool, n)
	k--
	for i := n - 1; i >= 0; i-- {
		p := k / fact[i]
		v := getNthEle(p+1, used)
		used[v] = true
		out = append(out, '0'+byte(v+1))
		k %= fact[i]
	}
	return string(out)
}
