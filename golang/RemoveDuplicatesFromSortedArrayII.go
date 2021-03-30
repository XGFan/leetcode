package main

import "fmt"

func removeDuplicates(nums []int) int {
	count := 1
	last := nums[0]
	repeated := false
	for i := 1; i < len(nums); i++ {
		if nums[i] != last {
			nums[count] = nums[i]
			last = nums[i]
			count++
			repeated = false
		} else {
			if !repeated {
				nums[count] = nums[i]
				last = nums[i]
				count++
				repeated = true
			}
		}
	}
	return count
}

func main() {
	ints := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates(ints))
	fmt.Println(ints)
	ints = []int{1,1,1,2,2,3}
	fmt.Println(removeDuplicates(ints))
	fmt.Println(ints)
}
