package main

import "fmt"

func findAnagrams(s string, p string) []int {
	var output []int

	if len(p) > len(s) {
		return output
	}

	pMap := make(map[rune]int)
	for _, ch := range p {
		pMap[ch] = pMap[ch] + 1
	}

	anagramLength := len(p)

	for i := 0; i < anagramLength; i++ {
		r := rune(s[i])
		_, ok := pMap[r]
		if ok {
			pMap[r] = pMap[r] - 1
		}
	}

	if isAnagram(pMap) {
		output = []int{0}
	}

	for i := anagramLength; i < len(s); i++ {
		rLeft := rune(s[i-anagramLength])
		rRight := rune(s[i])

		_, ok := pMap[rLeft]
		if ok {
			pMap[rLeft] = pMap[rLeft] + 1
		}

		_, ok = pMap[rRight]
		if ok {
			pMap[rRight] = pMap[rRight] - 1

			if isAnagram(pMap) {
				output = append(output, i-anagramLength+1)
			}
		}
	}

	return output
}

func isAnagram(m map[rune]int) bool {
	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Printf("%v\n", findAnagrams("cbaebabacdabcabcabcabcabc", "abc"))
	fmt.Printf("%v\n", findAnagrams("abab", "ab"))
}
