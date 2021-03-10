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

func getPermutation2(n int, k int) string {
	ints := make([]int, n)
	total := 1
	t := 0
	for i := 1; i <= n; i++ {
		total *= i
		if k <= total {
			t = k - total/i
			total = i - 1
			break
		}
	}
	for i := 0; i < n; i++ {
		if i < n-total {
			ints[i] = i + 1
		} else {
			ints[n-1-(i-(n-total))] = i + 1
		}
	}
	//fmt.Println(total, t, ints)
	for i := 0; i < t; i++ {
		next(ints)
	}
	builder := strings.Builder{}
	for _, v := range ints {
		builder.WriteByte(byte(v + 48))
	}
	return builder.String()
}
func main() {
	fmt.Println(getPermutation(3, 6))
	fmt.Println(getPermutation2(3, 6))
	fmt.Println(getPermutation(4, 13))
	fmt.Println(getPermutation2(4, 13))
	fmt.Println(getPermutation(9, 132))
	fmt.Println(getPermutation2(9, 132))
	fmt.Println(getPermutation(4, 1))
	fmt.Println(getPermutation2(4, 1))

}
