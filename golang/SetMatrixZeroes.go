package main

import "fmt"

func setZeroes(matrix [][]int) {
	xcol := false
	ycol := false
	for x := 0; x < len(matrix); x++ {
		if matrix[x][0] == 0 {
			ycol = true
			break
		}
	}
	for y := 0; y < len(matrix[0]); y++ {
		if matrix[0][y] == 0 {
			xcol = true
			break
		}
	}
	for x := 1; x < len(matrix); x++ {
		for y := 1; y < len(matrix[x]); y++ {
			if matrix[x][y] == 0 {
				matrix[x][0] = 0
				matrix[0][y] = 0
			}
		}
	}
	for x := 1; x < len(matrix); x++ {
		for y := 1; y < len(matrix[x]); y++ {
			if matrix[x][0] == 0 || matrix[0][y] == 0 {
				matrix[x][y] = 0
			}
		}
	}
	if xcol {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
	if ycol {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}

func main() {
	var a [][]int
	a = [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(a)
	fmt.Println(a)
	a = [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(a)
	fmt.Println(a)
	a = [][]int{{1, 0, 3}}
	setZeroes(a)
	fmt.Println(a)
	a = [][]int{{1, 2, 3, 4}, {5, 0, 7, 8}, {0, 10, 11, 12}, {13, 14, 15, 0}}
	setZeroes(a)
	fmt.Println(a)
}
