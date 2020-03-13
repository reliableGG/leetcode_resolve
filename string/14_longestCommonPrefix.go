package main

import "fmt"

func main() {
	//ca := []string{"flower", "flow", "flight"}
	ca := []string{"aa", "a"}
	ans := longestCommonPrefix(ca)
	fmt.Println(ans)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	first := strs[0]
	ans := []byte{}
	for i := 0; i < len(first); i++ {
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) <= i || strs[j][i] != first[i] {
				return string(ans)
			}
		}
		ans = append(ans, byte(first[i]))
	}

	return string(ans)
}
