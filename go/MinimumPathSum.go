package main

import "fmt"

func minV(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	for i := m + n; i >= 0; i-- {
		for p := 0; p < m; p++ {
			q := i - p
			if q >= 0 && q < n {
				//fmt.Printf("[%d,%d] from %d to ", p, q, grid[p][q])
				if p < m-1 && q < n-1 {
					grid[p][q] += minV(grid[p+1][q], grid[p][q+1])
				} else {
					if p < m-1 {
						grid[p][q] += grid[p+1][q]
					}
					if q < n-1 {
						grid[p][q] += grid[p][q+1]
					}
				}
				//fmt.Println(grid[p][q])
			}
		}
	}
	return grid[0][0]
}

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
