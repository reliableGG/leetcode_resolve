package main

import "fmt"

func main() {
	m := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(m)
	fmt.Println(m)
}

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	row := make([]bool, len(matrix))
	column := make([]bool, len(matrix[0]))
	length := len(matrix[0])

	for ri, r := range matrix {
		for ci, _ := range r {
			if matrix[ri][ci] == 0 {
				row[ri] = true
				column[ci] = true
			}
		}
	}

	for idx, r := range row {
		if r {
			matrix[idx] = make([]int, length)
		}
	}

	for idx, c := range column {
		if c {
			for _, r := range matrix {
				r[idx] = 0
			}
		}
	}

}
