package main

import "fmt"

func findTheDifference(s string, t string) byte {
	m := make(map[rune]int64)
	for _, v := range s {
		m[v]++
	}

	for _, v := range t {
		m[v]--
		if m[v] < 0 {
			return byte(v)
		}
	}

	return 0
}

func main() {
	fmt.Printf("%v\n", findTheDifference("abcd", "abcde"))
}
