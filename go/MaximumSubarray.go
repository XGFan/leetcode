package main

import "fmt"

//TODO: Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
func maxSubArray(nums []int) int {
	sum := nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		v := nums[i]
		if v >= 0 {
			if sum >= 0 {
				sum += v
			} else {
				sum = v
			}
			max = maxV(max, sum)
		} else {
			if sum >= 0 {
				sum += v
				//no need to update max
			} else {
				sum = v
				max = maxV(max, sum)
			}
		}
	}
	return max
}

func maxV(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{0}))
	fmt.Println(maxSubArray([]int{-2, -1}))
}
