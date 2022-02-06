package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	length := 1
	met := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[length-1] {
			if met < 2 {
				met++
				nums[length] = nums[i]
				length++
			}
		} else {
			met = 1
			nums[length] = nums[i]
			length++
		}
	}

	return length
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Printf("%d, %v\n", removeDuplicates(nums), nums)
	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3, 3, 3, 3, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8}
	fmt.Printf("%d, %v\n", removeDuplicates(nums), nums)
}
