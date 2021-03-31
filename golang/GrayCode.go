package main

import "fmt"

/**
这个感觉是个数学题
*/
func grayCode(n int) []int {
	res := make([]int, 1<<n)
	fillGradCode(res, n)
	return res
}

func fillGradCode(res []int, n int) int {
	if n == 1 {
		res[0], res[1] = 0, 1
		return 2
	} else {
		l := fillGradCode(res, n-1)
		for i := 0; i < l; i++ {
			res[l+i] = l + res[l-1-i]
		}
		return 2 * l
	}
}

func main() {
	fmt.Println(grayCode(1))
	fmt.Println(grayCode(2))
	fmt.Println(grayCode(3))
}
