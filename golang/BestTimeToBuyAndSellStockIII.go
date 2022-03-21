package main

func maxProfit(prices []int) int {
	buy1, buy2, sell1, sell2 := -prices[0], -prices[0], 0, 0
	for i := 1; i < len(prices); i++ {
		nbuy1 := max(buy1, -prices[i])
		nsell1 := max(sell1, buy1+prices[i])
		nbuy2 := max(sell1-prices[i], buy2)
		nsell2 := max(sell2, buy2+prices[i])
		buy1, sell1, buy2, sell2 = nbuy1, nsell1, nbuy2, nsell2
	}
	return max(sell2, sell1)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	_ = maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4})
}
