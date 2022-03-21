package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums); {
		needSum := 0 - nums[i]
		for l, r := i+1, len(nums)-1; l < r; {
			if nums[l]+nums[r] < needSum {
				l++
			} else if nums[l]+nums[r] > needSum {
				r--
			} else {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				for l++; l < r && nums[l] == nums[l-1]; l++ {
				}
				for r--; l < r && nums[r] == nums[r+1]; r-- {
				}
			}
		}
		for i++; i < len(nums)-1 && nums[i] == nums[i-1]; i++ {
		}
	}
	return result
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
