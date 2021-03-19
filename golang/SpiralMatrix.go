package main

import (
	"fmt"
)

func visit(matrix [][]int, f int, result []int) []int {
	Y := len(matrix[0]) - 2*f
	X := len(matrix) - 2*f
	x, y := f, f
	if (X == 0 || Y == 0) || (X < 0 || Y < 0) {
		return result
	}
	result = append(result, matrix[x][y])
	for i := 1; i < Y; i++ {
		y++
		result = append(result, matrix[x][y])
	}
	if X <= 1 {
		return result
	}
	for i := 1; i < X; i++ {
		x++
		result = append(result, matrix[x][y])
	}
	if Y <= 1 {
		return result
	}
	for i := 1; i < Y; i++ {
		y--
		result = append(result, matrix[x][y])
	}
	for i := 1; i < X-1; i++ {
		x--
		result = append(result, matrix[x][y])
	}
	return visit(matrix, f+1, result)
}

func spiralOrder(matrix [][]int) []int {
	Y := len(matrix[0])
	X := len(matrix)
	amount := Y * X
	ints := make([]int, 0, amount)
	return visit(matrix, 0, ints)
}
func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(spiralOrder([][]int{{1, 2}, {3, 4}}))
	fmt.Println(spiralOrder([][]int{{1, 2}, {3, 4}, {5, 6}}))
}
