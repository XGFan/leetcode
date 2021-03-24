package main

func setZeroes(matrix [][]int) {
	xs := make([]int, 0)
	ys := make([]int, 0)
	for x := range matrix {
		for y := range matrix[x] {
			if matrix[x][y] == 0 {
				xs = append(xs, x)
				ys = append(ys, y)
			}
		}
	}
	for _, x := range xs {
		for y := range matrix[x] {
			matrix[x][y] = 0
		}
	}
	for _, y := range ys {
		for x := range matrix {
			matrix[x][y] = 0
		}
	}
}
