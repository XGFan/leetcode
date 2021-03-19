package main

import (
	"fmt"
)

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	write(result, 0, 0)
	return result
}

func write(matrix [][]int, p, count int) {
	X := len(matrix) - 2*p
	x, y := p, p-1
	if X == 0 {
		return
	}
	if X == 1 {
		matrix[x][y+1] = count + 1
		return
	}
	for i := 0; i < X; i++ {
		y++
		count++
		matrix[x][y] = count
	}

	for i := 1; i < X; i++ {
		x++
		count++
		matrix[x][y] = count
	}
	for i := 1; i < X; i++ {
		y--
		count++
		matrix[x][y] = count
	}
	for i := 1; i < X-1; i++ {
		x--
		count++
		matrix[x][y] = count
	}
	write(matrix, p+1, count)
}

func main() {
	fmt.Println(generateMatrix(1))
	fmt.Println(generateMatrix(2))
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix(4))
}
