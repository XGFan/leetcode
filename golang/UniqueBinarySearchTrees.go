package main

import "fmt"

func numTrees(n int) int {
	cache := make([]int, n+1)
	cache[0] = 1
	return numTreesIter(cache, n)
}

func numTreesIter(cache []int, total int) int {
	if cache[total] != 0 {
		return cache[total]
	}
	res := 0
	for i := 0; i <= total-1; i++ {
		ltrees := numTreesIter(cache, i)
		rtrees := numTreesIter(cache, total-1-i)
		res += ltrees * rtrees
	}
	cache[total] = res
	return res
}
func main() {
	fmt.Println(numTrees(1))
	fmt.Println(numTrees(2))
	fmt.Println(numTrees(3))
}
