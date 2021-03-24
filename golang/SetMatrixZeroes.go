package main

func setZeroes(matrix [][]int) {
	ys := make([]int, 0)
	for x := range matrix {
		justClean := false
		for y := range matrix[x] {
			if matrix[x][y] == 0 {
				ys = append(ys, y)
				if !justClean {
					for i := 0; i < y; i++ {
						matrix[x][i] = 0
					}
					justClean = true
				}
			} else {
				if justClean {
					matrix[x][y] = 0
				}
			}
		}
	}
	for _, y := range ys {
		for x := range matrix {
			matrix[x][y] = 0
		}
	}
}
