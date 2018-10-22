package main

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			swapELe(matrix, i, j, n-j-1, n-i-1)
		}
	}
	for j := 0; j < n; j++ {
		for i := 0; i < n/2; i++ {
			swapELe(matrix, i, j, n-i-1, j)
		}
	}
}

func swapELe(m [][]int, i, j, p, q int) {
	m[i][j], m[p][q] = m[p][q], m[i][j]
}

func main() {
	/*
	 *m := [][]int{{1, 2, 3},
	 *    {4, 5, 6},
	 *    {7, 8, 9}}
	 */
	//m := [][]int{{1}}
	m1 := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(m1)
	fmt.Println(m1)
}
