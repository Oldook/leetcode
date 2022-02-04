package main

import "fmt"

func findMaxLength(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	sumIndexes := map[int]int{0: -1}

	maxLen := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			sum -= 1
		} else {
			sum += 1
		}

		sumIdx, ok := sumIndexes[sum]
		if !ok {
			sumIndexes[sum] = i
		} else {
			curLen := i - sumIdx
			if curLen > maxLen {
				maxLen = curLen
			}
		}
	}

	return maxLen
}

func main() {
	fmt.Printf("%d\n", findMaxLength([]int{1, 0}))
	fmt.Printf("%d\n", findMaxLength([]int{1, 0, 1}))
	fmt.Printf("%d\n", findMaxLength([]int{1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 0}))
	fmt.Printf("%d\n", findMaxLength([]int{0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 1}))
}
