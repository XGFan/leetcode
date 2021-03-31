package main

import "fmt"

func qs(ints []int, l, r int) {
	if r > l {
		pivot := par(ints, l, r)
		qs(ints, l, pivot-1)
		qs(ints, pivot+1, r)
	}
}
func par(ints []int, start, end int) int {
	l := start
	for p := start; p < end; p++ {
		if ints[p] < ints[end] {
			if l != p {
				ints[p], ints[l] = ints[l], ints[p]
			}
			l++
		}
	}
	ints[l], ints[end] = ints[end], ints[l]
	return l
}

func subsetsWithDup(nums []int) [][]int {
	//创建指定长度的数组，下标0留给空数组
	res := make([][]int, 1, 1<<len(nums))
	qs(nums, 0, len(nums)-1)
	subsetsWithDupIter(&res, []int{}, nums)
	return res
}
func subsetsWithDupIter(res *[][]int, prefix, other []int) {
	t := 0
	for i := range other {
		if t&(1<<(other[i]+10)) == 0 {
			t |= 1 << (other[i] + 10)
			ints := make([]int, len(prefix)+1)
			ints[copy(ints, prefix)] = other[i]
			*res = append(*res, ints)
			subsetsWithDupIter(res, ints, other[i+1:])
		}
	}
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	fmt.Println(subsetsWithDup([]int{1, 4, 4, 4, 4}))
	fmt.Println(subsetsWithDup([]int{4, 1, 4, 4, 4}))
}
