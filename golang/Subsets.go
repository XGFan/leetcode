package main

import "fmt"

func g(res *[][]int, prefix, other []int) {
	for i := range other {
		ints := make([]int, len(prefix))
		copy(ints, prefix)
		ints = append(ints, other[i])
		//fmt.Println("ADD", i, ints)
		*res = append(*res, ints)
		g(res, append(prefix, other[i]), other[i+1:])
	}
}

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{})
	g(&res, []int{}, nums)
	return res
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
