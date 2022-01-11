package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s := sum(l1, l2, 0)
	return s
}

func sum(l1 *ListNode, l2 *ListNode, mem int) *ListNode {
	if l1 == nil && l2 == nil {
		if mem == 1 {
			return &ListNode{
				Val:  1,
				Next: nil,
			}
		}

		return nil
	}

	val1 := 0
	if l1 != nil {
		val1 = l1.Val
		l1 = l1.Next
	}

	val2 := 0
	if l2 != nil {
		val2 = l2.Val
		l2 = l2.Next
	}

	stepVal := val1 + val2 + mem
	if stepVal > 9 {
		stepVal %= 10
		mem = 1
	} else {
		mem = 0
	}

	return &ListNode{
		Val:  stepVal,
		Next: sum(l1, l2, mem),
	}
}

func main() {
	result := addTwoNumbers(
		&ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
		&ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	)

	for result != nil {
		fmt.Printf("%d", result.Val)

		result = result.Next
	}

	fmt.Printf("\n")
}
