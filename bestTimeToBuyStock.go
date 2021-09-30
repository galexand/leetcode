package main

import (
	"fmt"
	"math"
)

/**
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
 */
func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}

	var maxProfit int = 0
	var min int = math.MaxInt32
	var max int = 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i] > max {
			max = prices[i]
		}
		if maxProfit < max-min {
			maxProfit = max - min
		}
	}

	return maxProfit
}

func main() {
	fmt.Println("Max profit for [7,1,5,3,6,4] should be 5: ", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println("Max profit for [7,6,4,3,1] should be 0: ", maxProfit([]int{7, 6, 4, 3, 1}))
}
