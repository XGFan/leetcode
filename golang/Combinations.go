package main

import "fmt"

func g2(prefix []int, from, to, n int, result *[][]int) {
	if n == 0 {
		ints := make([]int, len(prefix))
		copy(ints, append(prefix))
		*result = append(*result, ints)
	} else {
		for i := from; i <= to; i++ {
			g2(append(prefix, i), i+1, to, n-1, result)
		}
	}
}

func combine(n int, k int) [][]int {
	res := make([][]int, 0, c(n, k))
	ints := make([]int, 0, k)
	g2(ints, 1, n, k, &res)
	return res
}

func c(n, k int) int {
	c := 1
	x, y := n-k, k
	if x > y {
		x, y = y, x
	}
	for i := n; i > 0; i-- {
		if i > y {
			c *= i
		} else if i <= x {
			c /= i
		}
	}
	return c
}

func main() {
	//fmt.Println(c(4, 2))
	//fmt.Println(c(1, 1))
	//fmt.Println(c(4, 1))
	//fmt.Println(c(4, 3))
	//fmt.Println(c(7, 3))
	//fmt.Println(combine(4, 2))
	fmt.Println(combine(5, 4))
	fmt.Println(combine(7, 4))
}
