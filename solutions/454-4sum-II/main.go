package main

import "fmt"

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count := 0

	n := len(nums1)
	if n == 0 {
		return count
	}

	m := make(map[int]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m[nums1[i]+nums2[j]]++
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			count += m[-nums3[i]-nums4[j]]
		}
	}

	return count
}

func main() {
	fmt.Printf("%v\n", fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}))
	fmt.Printf("%v\n", fourSumCount([]int{0, 0}, []int{0, 0}, []int{0, 0}, []int{0, 0}))
}
