package main

import "fmt"

func addDigits(num int) int {
	if num == 0 {
		return 0
	}

	if res := num % 9; res == 0 {
		return 9
	} else {
		return res
	}
}

func main() {
	fmt.Printf("%d\n", addDigits(0))
	fmt.Printf("%d\n", addDigits(10))
	fmt.Printf("%d\n", addDigits(5))
	fmt.Printf("%d\n", addDigits(44))
}
