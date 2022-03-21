package main

import "fmt"

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfitAmount := 0
	for p := 0; p < len(prices); p++ {
		if prices[p] < minPrice {
			minPrice = prices[p]
		} else if prices[p] > minPrice {
			if prices[p]-minPrice > maxProfitAmount {
				maxProfitAmount = prices[p] - minPrice
			}
		}
	}
	return maxProfitAmount
}

func maxSubIncrListSize(nums []int) int {
	results := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		results[i] = 1
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				if results[j]+1 > results[i] {
					results[i] = results[j] + 1
				}
			}
		}
	}
	max := 1
	for _, v := range results {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	fmt.Println(maxSubIncrListSize([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(maxSubIncrListSize([]int{4, 2, 5, 3, 5, 6}))
}
