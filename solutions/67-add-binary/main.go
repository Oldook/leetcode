package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	result := ""

	idxA := len(a) - 1
	idxB := len(b) - 1

	carry := 0
	for idxA > -1 || idxB > -1 {
		// 1 - 49, 0 - 48
		firstBit, secondBit := 48, 48
		if idxA > -1 {
			firstBit = int(a[idxA])
			idxA--
		}

		if idxB > -1 {
			secondBit = int(b[idxB])
			idxB--
		}

		sum := firstBit + secondBit + carry
		switch sum {
		case 96:
			result = string('0') + result
			carry = 0
		case 97:
			result = string('1') + result
			carry = 0
		case 98:
			result = string('0') + result
			carry = 1
		case 99:
			result = string('1') + result
			carry = 1
		}
	}

	if carry == 1 {
		result = string('1') + result
	}

	return result
}

func main() {
	fmt.Printf("%s\n", addBinary("11", "1"))
}
