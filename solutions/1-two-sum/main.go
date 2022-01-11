package main

import "fmt"

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int, len(nums))

	for i, v := range nums {
		p, ok := indexMap[target-v]
		if ok {
			return []int{i, p}
		} else {
			indexMap[v] = i
		}
	}

	return nil
}

func main() {
	fmt.Printf("%v\n", twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Printf("%v\n", twoSum([]int{3, 2, 4}, 6))
	fmt.Printf("%v\n", twoSum([]int{3, 3}, 6))
}
