package main

import "fmt"

func uniquePaths(m int, n int) int {
	if m > n {
		m, n = n, m
	}
	c := 1
	//Find the Number of Possibilities 9 choose 3
	//aCb = a!/b!(a-b)!
	for i := m + n - 2; i > 0; i-- {
		if i >= n {
			c *= i
		} else if i < m {
			c /= i
		}
	}
	return c
}

func main() {
	fmt.Println(uniquePaths(3, 2))
	fmt.Println(uniquePaths(3, 7))
	fmt.Println(uniquePaths(3, 3))
	fmt.Println(uniquePaths(51, 9))
}
