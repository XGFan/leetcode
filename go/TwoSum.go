package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		t := m[target-num]
		if t != 0 {
			return []int{i, t - 1}
		} else {
			m[num] = i + 1
		}
	}
	return nil
}
