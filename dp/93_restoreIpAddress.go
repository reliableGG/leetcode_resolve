package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "25525511135"
	res := restoreIpAddresses(input)
	fmt.Println(res)
}

func restoreIpAddresses(s string) []string {
	result := []string{}
	ip := []string{}
	dfs(s, ip, &result, 0)
	return result
}

func dfs(s string, ip []string, result *[]string, start int) {
	fmt.Println(s)
	if len(ip) == 4 && start == len(s) {
		*result = append(*result, strings.Join(ip, "."))
		return
	}

	if len(ip) == 4 || len(s)-start > 3*(4-len(ip)) || len(s)-start < 4-len(ip) {
		return
	}

	num := 0
	for i := start; i < start+3 && i < len(s); i++ {
		num = num*10 + int(s[i]-byte(0))
		if num < 0 || num > 255 {
			continue
		}

		ip = append(ip, s[start:i+1])
		dfs(s, ip, result, i+1)
		ip = ip[:len(ip)-1]
	}

}
