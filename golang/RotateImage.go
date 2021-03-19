package main

import "fmt"

func rotate(matrix [][]int) {
	var x, y int
	n := len(matrix) - 1
	for {
		//fmt.Print(".")
		matrix[x][y], matrix[y][n-x], matrix[n-x][n-y], matrix[n-y][x] = matrix[n-y][x], matrix[x][y], matrix[y][n-x], matrix[n-x][n-y]
		if y < n-1-x {
			y++
		} else {
			//fmt.Println()
			x++
			y = x
			if x >= (n+1)/2 {
				break
			}
		}
	}
}

func main() {
	matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(matrix)
	fmt.Printf("%v \n", matrix)
	m2 := [][]int{{2, 29, 20, 26, 16, 28}, {12, 27, 9, 25, 13, 21}, {32, 33, 32, 2, 28, 14}, {13, 14, 32, 27, 22, 26}, {33, 1, 20, 7, 21, 7}, {4, 24, 1, 6, 32, 34}}
	rotate(m2)
	fmt.Printf("%v \n", m2)
}
