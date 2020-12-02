package main

import "fmt"

func trap(height []int) int {
	var trap, i, j, t, maxIndex int
	for i = 0; i < len(height); i = j {
		t = 0
		for j = i + 1; j < len(height); j++ {
			if height[j] < height[i] {
				t += height[i] - height[j]
			} else {
				maxIndex = j
				trap += t
				break
			}
		}
	}

	for i = len(height) - 1; i > maxIndex; i = j {
		t = 0
		for j = i - 1; j >= maxIndex; j-- {
			if height[j] < height[i] {
				t += height[i] - height[j]
			} else {
				trap += t
				break
			}
		}
	}
	return trap
}

func trap1(height []int) int {
	trap := 0
	maxIndex := 0
	for index, value := range height {
		if value > trap {
			trap = value
			maxIndex = index
		}
	}
	trap = 0
	i := 0
	for i < maxIndex {
		value := height[i]
		j := i + 1
		t := 0
		for j <= maxIndex {
			if height[j] < value {
				t += value - height[j]
				j++
			} else {
				trap += t
				break
			}
		}
		i = j
	}
	i = len(height) - 1
	for i > maxIndex {
		value := height[i]
		j := i - 1
		t := 0
		for j >= maxIndex {
			if height[j] < value {
				t += value - height[j]
				j--
			} else {
				trap += t
				break
			}
		}
		i = j
	}
	return trap
}

func trap0(height []int) int {
	max := 0
	for _, value := range height {
		if value > max {
			max = value
		}
	}
	trap := 0
	for i := 1; i <= max; i++ {
		shouldCount := false
		t := 0
		for _, value := range height {
			if value >= i {
				if shouldCount {
					trap += t
					t = 0
				} else {
					shouldCount = true
				}
			} else {
				if shouldCount {
					t++
				}
			}
		}
	}
	return trap
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{4, 2, 3}))
}

//
