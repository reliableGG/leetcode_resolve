package main

import "fmt"

func main() {
	ca := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(ca)
	fmt.Println(ca)
}
func gameOfLife(board [][]int) {
	m := len(board)
	n := len(board[0])

	dx := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	dy := []int{-1, -1, -1, 0, 0, 1, 1, 1}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			live := 0
			for k := 0; k < 8; k++ {
				x := i + dx[k]
				y := j + dy[k]
				if x > -1 && x < m && y > -1 && y < n && (board[x][y] == 1 || board[x][y] == 2) {
					live++
				}
			}
			if board[i][j] == 0 && live == 3 {
				board[i][j] = 3
			} else if board[i][j] == 1 && (live < 2 || live > 3) {
				board[i][j] = 2
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] %= 2
		}
	}
}
