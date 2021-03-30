package main

func search(nums []int, target int) bool {
	if nums[0] < target {
		for i := 1; i < len(nums); i++ {
			if nums[i] == target {
				return true
			} else if nums[i] > target {
				return false
			}
		}
		return false
	} else if nums[0] > target {
		for i := len(nums) - 1; i > 0; i-- {
			if nums[i] == target {
				return true
			} else if nums[i] < target {
				return false
			}
		}
		return false
	} else {
		return true
	}

}
