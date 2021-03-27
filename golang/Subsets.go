package main

import "fmt"

func g(res *[][]int, prefix, other []int) {
	for i := range other {
		ints := make([]int, len(prefix)+1)
		ints[copy(ints, prefix)] = other[i]
		*res = append(*res, ints)
		g(res, ints, other[i+1:])
	}
}

func subsets(nums []int) [][]int {
	res := make([][]int, 1, 1<<len(nums))
	g(&res, []int{}, nums)
	return res
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
