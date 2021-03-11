package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	if obstacleGrid[m-1][n-1] == 0 {
		obstacleGrid[m-1][n-1] = -1
	} else {
		return 0
	}
	for i := m + n; i >= 0; i-- {
		for p := 0; p < m; p++ {
			q := i - p
			if q >= 0 && q < n {
				cur := obstacleGrid[p][q]
				if cur == 1 {
					obstacleGrid[0][0] = 0
				} else {
					if p != 0 {
						if obstacleGrid[p-1][q] != 1 {
							obstacleGrid[p-1][q] += cur
						}
					}
					if q != 0 {
						if obstacleGrid[p][q-1] != 1 {
							obstacleGrid[p][q-1] += cur
						}
					}
				}
			}
		}
	}
	return -obstacleGrid[0][0]
}

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{{1, 0}}))
	//fmt.Println(uniquePathsWithObstacles([][]int{{1, 1}}))
	//fmt.Println(uniquePathsWithObstacles([][]int{{0, 0}, {1, 1}, {0, 0}}))
	//fmt.Println(uniquePathsWithObstacles([][]int{{0}}))
	//tryTable(4, 3)
}
