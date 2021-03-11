package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	table := make([][]int, m)
	for i := range table {
		table[i] = make([]int, n)
	}
	if obstacleGrid[m-1][n-1] == 0 {
		table[m-1][n-1] = 1
	} else {
		return 0
	}
	for i := m + n; i >= 0; i-- {
		for p := 0; p < m; p++ {
			q := i - p
			if q >= 0 && q < n {
				//fmt.Print("[", p, q, "] ")
				cur := obstacleGrid[p][q]
				if cur == 1 {
					//无事发生
				} else {
					if p == 0 {
						if q == 0 {
							return table[p][q]
						} else {
							if obstacleGrid[p][q-1] != 1 {
								table[p][q-1] += table[p][q]
							}
						}
					} else {
						if q == 0 {
							if obstacleGrid[p-1][q] != 1 {
								table[p-1][q] += table[p][q]
							}
						} else {
							if obstacleGrid[p][q-1] != 1 {
								table[p][q-1] += table[p][q]
							}
							if obstacleGrid[p-1][q] != 1 {
								table[p-1][q] += table[p][q]
							}
						}
					}

				}
			}
		}
	}
	return table[0][0]
}

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{{1, 0}}))
	//fmt.Println(uniquePathsWithObstacles([][]int{{1, 1}}))
	//fmt.Println(uniquePathsWithObstacles([][]int{{0, 0}, {1, 1}, {0, 0}}))
	//fmt.Println(uniquePathsWithObstacles([][]int{{0}}))
	//tryTable(4, 3)
}
