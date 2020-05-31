package main

import "fmt"

func main() {
	res := uniquePaths(2, 2)
	fmt.Println(res)
}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	return uniquePathsDp(dp, m, n)
}

func uniquePathsDp(dp [][]int, m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	if dp[m-1][n-1] > 0 {
		return dp[m-1][n-1]
	} else {
		dp[m-1][n-1] = uniquePathsDp(dp, m-1, n) + uniquePathsDp(dp, m, n-1)
		return dp[m-1][n-1]
	}

}

/*func uniquePaths(m int, n int) int {*/
//if m < 0 || n < 0 {
//return 0
//}

//if m == 1 && n == 1 {
//return 1
//}
//return uniquePaths(m-1, n) + uniquePaths(m, n-1)
/*}*/
