package main

func firstMissingPositive(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 1
	}
	var index = 0
	for index < size {
		value := nums[index]
		orderShouldBe := value - 1
		//正数 && 小于等于长度 && 不在应该在的位置 && 应该在的位置不是现在的值
		if value > 0 && value <= size && index != orderShouldBe && nums[orderShouldBe] != value {
			nums[index], nums[orderShouldBe] = nums[orderShouldBe], nums[index]
		} else {
			index++
		}
	}
	if nums[0] != 1 {
		return 1
	}
	index = 0
	for index < size && nums[index] == index+1 {
		index++
	}
	return index + 1
}
