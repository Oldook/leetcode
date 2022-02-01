package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	result := 0
	minPrice := prices[0]

	for _, v := range prices[1:] {
		if v < minPrice {
			minPrice = v
		}

		diff := v - minPrice
		if diff > result {
			result = diff
		}
	}

	return result
}

func main() {
	fmt.Printf("%v\n", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Printf("%v\n", maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Printf("%v\n", maxProfit([]int{}))
	fmt.Printf("%v\n", maxProfit([]int{1}))
}
