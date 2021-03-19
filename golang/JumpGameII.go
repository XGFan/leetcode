package main

import "fmt"

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	var limit, next, steps int
	for i := 0; i <= limit; i++ { //从左到limit
		if i+nums[i] > next { //更新最大的next
			next = i + nums[i]
		}
		if i == limit { //如果已经移动到了limit
			steps++
			limit = next              //从0-limit之间选出最远的next
			if limit >= len(nums)-1 { //如果current-limit之间的最远距离已经达到目标
				break
			}
		}
	}
	return steps
}
func jump2(nums []int) int {
	var steps, last, curr int
	for i := 0; i < len(nums); i++ {
		// position beyond last reach
		if i > last {
			last = curr
			steps++
		}
		// furthest distance can be reach
		if curr < i+nums[i] {
			curr = i + nums[i]
		}
	}
	return steps
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func jump3(nums []int) int {
	return jumpIter(nums, 0)
}

func findMin(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i+nums[i] >= len(nums)-1 {
			return i
		}
	}
	return len(nums) - 1
}
func jumpIter(nums []int, count int) int {
	if len(nums) == 1 {
		return count
	}
	return jumpIter(nums[:findMin(nums)+1], count+1)
}

func main() {
	fmt.Println(jump([]int{2, 3, 0, 1, 4}))
	fmt.Println(jump([]int{0}))
	fmt.Println(jump([]int{1, 2}))
	fmt.Println(jump([]int{1, 2, 3}))
}
