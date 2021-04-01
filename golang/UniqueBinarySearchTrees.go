package main

import "fmt"

func numTrees(n int) int {
	return numTreesIter(n)
}

func numTreesIter(total int) int {
	if total == 0 {
		return 1
	}
	res := 0
	for i := 0; i <= total-1; i++ {
		ltrees := numTreesIter(i)
		rtrees := numTreesIter(total - 1 - i)
		res += ltrees * rtrees
	}
	return res
}

func main() {
	fmt.Println(numTrees(1))
	fmt.Println(numTrees(2))
	fmt.Println(numTrees(3))
}
