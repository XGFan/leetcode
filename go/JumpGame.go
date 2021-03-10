package main

import "fmt"

func canJump(nums []int) bool {
	target := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= target {
			target = i
		}
	}
	return target == 0
}

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}), true)
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}), false)
}
