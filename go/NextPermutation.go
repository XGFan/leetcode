package main

func nextPermutation(nums []int) {
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
	nextPermutation([]int{1, 2, 3})
}
