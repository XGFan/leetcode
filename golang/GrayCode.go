package main

import "fmt"

/**
这个感觉是个数学题
*/
func grayCode(n int) []int {
	if n == 1 {
		return []int{0, 1}
	} else {
		prev := grayCode(n - 1)
		ints := make([]int, len(prev)*2)
		for i := range ints {
			if i < len(prev) {
				ints[i] = prev[i]
			} else {
				ints[i] = 1<<(n-1) + prev[len(ints)-i-1]
			}
		}
		return ints
	}
}

func main() {
	fmt.Println(grayCode(1))
	fmt.Println(grayCode(2))
	fmt.Println(grayCode(3))
}
