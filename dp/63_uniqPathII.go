package main

import "fmt"

func main() {
	res := uniquePathsWithObstacles()
	fmt.Println(res)
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	if obstacleGrid[0][0] != 0 || obstacleGrid[m-1][n-1] != 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	if obstacleGrid[0][0] == 0 {
		dp[0][0] = 1
	}
	return uniquePathsDp(obstacleGrid, dp, m-1, n-1)
}

func uniquePathsDp(obstacleGrid [][]int, dp [][]int, x int, y int) int {
	if x < 0 || y < 0 {
		return 0
	}

	if obstacleGrid[x][y] != 0 {
		return 0
	}

	if x == 0 && y == 0 {
		return dp[0][0]
	}

	if dp[x][y] > 0 {
		return dp[x][y]
	} else {
		dp[x][y] = uniquePathsDp(obstacleGrid, dp, x-1, y) + uniquePathsDp(obstacleGrid, dp, x, y-1)
		return dp[x][y]
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
