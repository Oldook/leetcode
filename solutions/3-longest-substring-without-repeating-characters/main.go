package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	uniqueSymbols := make(map[rune]int)
	maxLength := 0
	offset := 0

	for i, v := range s {
		if symbolIdx, ok := uniqueSymbols[v]; ok {
			curLen := len(uniqueSymbols)
			if curLen > maxLength {
				maxLength = curLen
			}

			for j := offset; j <= symbolIdx; j++ {
				delete(uniqueSymbols, rune(s[j]))
			}

			offset = symbolIdx + 1
		}

		uniqueSymbols[v] = i
	}

	if len(uniqueSymbols) > maxLength {
		maxLength = len(uniqueSymbols)
	}

	return maxLength
}

func main() {
	fmt.Printf("%d\n", lengthOfLongestSubstring("aabaab!bb"))
}
