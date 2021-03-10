package main

import (
	"fmt"
	"strings"
)

func getPermutation(n int, k int) string {
	ints := make([]int, n)
	for i := 0; i < len(ints); i++ {
		ints[i] = i + 1
	}
	for i := 1; i < k; i++ {
		next(ints)
	}
	builder := strings.Builder{}
	for _, v := range ints {
		builder.WriteByte(byte(v + 48))
	}
	return builder.String()
}

func next(nums []int) {
	var i = len(nums) - 1
	for i > 0 && nums[i] <= nums[i-1] {
		i--
	}
	p := i - 1
	if i != 0 {
		for i < len(nums) && nums[i] > nums[p] {
			i++
		}
		nums[p], nums[i-1] = nums[i-1], nums[p]
	}
	for i = 1; p+i < len(nums)-i; i++ {
		nums[p+i], nums[len(nums)-i] = nums[len(nums)-i], nums[p+i]
	}
}

func main() {
	fmt.Println(getPermutation(3, 3))
}
