package main

func numTrees(n int) int {
	ints := make([]int, n)
	for i := range ints {
		ints[i] = i + 1
	}
	return numTreesIter(ints)
}

func numTreesIter(ints []int) int {
	if len(ints) == 0 {
		return 1
	}
	res := 0
	for i := range ints {
		ltrees := numTreesIter(ints[:i])
		rtrees := numTreesIter(ints[i+1:])
		res += ltrees * rtrees
	}
	return res
}

func main() {

}
