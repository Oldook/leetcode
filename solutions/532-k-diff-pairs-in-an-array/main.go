package main

import (
	"fmt"
	"math"
)

func findPairs(nums []int, k int) int {
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	pairs := 0

	for _, num := range nums {
		left := num - k
		right := num + k

		if v, ok := m1[right]; ok {
			if v == num {
				pairs++
				m1[right] = math.MaxInt
				m2[num] = math.MaxInt
			}

		}

		if v, ok := m2[left]; ok {
			if v == num {
				pairs++
				m2[left] = math.MaxInt
				m1[num] = math.MaxInt
			}

		}

		if _, ok := m1[num]; !ok {
			m1[num] = left
		}

		if _, ok := m2[num]; !ok {
			m2[num] = right
		}
	}

	return pairs
}

func main() {
	fmt.Printf("%d\n", findPairs([]int{3, 1, 4, 1, 5}, 2))
	fmt.Printf("%d\n", findPairs([]int{1, 2, 3, 4, 5}, 1))
	fmt.Printf("%d\n", findPairs([]int{1, 3, 1, 5, 4}, 0))
	fmt.Printf("%d\n", findPairs([]int{1, 2, 4, 4, 3, 3, 0, 9, 2, 3}, 3))
	fmt.Printf("%d\n", findPairs([]int{3, 3, 0, 3}, 3))
}
