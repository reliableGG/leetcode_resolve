package main

import "fmt"

func fact(n int) []int {
	if n < 0 {
		return nil
	}
	var out []int
	out = append(out, 1)
	for i := 1; i <= n; i++ {
		out = append(out, out[len(out)-1]*i)
	}
	return out
}

func getNth(used []bool, n int) int {
	//fmt.Println(used, n)
	var i int = -1
	for n > 0 {
		i++
		if used[i] == false {
			n--
		}
	}
	return i
}

func getPermutation(n int, k int) string {
	fact := fact(n - 1)
	fmt.Println(fact)
	out := make([]byte, 0, n)
	used := make([]bool, n)
	k--
	for i := n - 1; i >= 0; i-- {
		p := k / fact[i]
		v := getNth(used, p+1)
		used[v] = true
		out = append(out, '0'+byte(v+1))
		k %= fact[i]
	}
	return string(out)
}

func main() {
	fmt.Println(getPermutation(4, 13))
}
