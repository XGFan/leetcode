package main

func sortColors(nums []int) {
	l := 0
	r := len(nums) - 1
	i := l
	for i <= r {
		switch nums[i] {
		case 0:
			nums[i], nums[l] = nums[l], nums[i]
			l++
			i = l
		case 1:
			i++
		case 2:
			nums[i], nums[r] = nums[r], nums[i]
			r--
		}
	}
}

func main() {
	sortColors([]int{2, 0, 2, 1, 1, 0})
	sortColors([]int{2, 0, 1})
	sortColors([]int{0})
	sortColors([]int{1})
	sortColors([]int{2})
	//sortColors([]int{2, 0, 1})
}
