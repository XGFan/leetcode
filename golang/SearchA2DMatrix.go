package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	y := len(matrix[0])
	l := 0
	r := len(matrix)*y - 1
	try := (l + r) / 2
	for l <= r {
		i, j := try/y, try%y
		//fmt.Println(l, try, r)
		if matrix[i][j] < target {
			if l == r {
				return false
			}
			l = try + 1
			try = (l + r + 1) / 2

		} else if matrix[i][j] > target {
			if l == r {
				return false
			}
			r = try - 1
			try = (l + r + 1) / 2
		} else {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	fmt.Println(!searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
	fmt.Println(searchMatrix([][]int{{1}}, 1))
	fmt.Println(searchMatrix([][]int{{1, 3}}, 1))
	fmt.Println(searchMatrix([][]int{{1, 3}}, 3))
	fmt.Println(!searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
	fmt.Println(!searchMatrix([][]int{{1, 60}}, 0))
}
