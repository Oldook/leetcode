package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}

	result := lists[0]

	for i := 1; i < length; i++ {
		result = merge(result, lists[i])
	}

	return result
}

func merge(first *ListNode, second *ListNode) *ListNode {
	node := &ListNode{}
	mergeResult := node

	for first != nil || second != nil {
		if first == nil {
			mergeResult.Next = second
			second = second.Next
			mergeResult = mergeResult.Next
			continue
		}

		if second == nil {
			mergeResult.Next = first
			first = first.Next
			mergeResult = mergeResult.Next
			continue
		}

		if first.Val <= second.Val {
			mergeResult.Next = first
			first = first.Next
		} else {
			mergeResult.Next = second
			second = second.Next
		}

		mergeResult = mergeResult.Next
	}

	return node.Next
}

func main() {
	//fmt.Printf("%v\n", mergeKLists(
	//	[]*ListNode{
	//		{
	//			Val: 1,
	//			Next: &ListNode{
	//				Val: 4,
	//				Next: &ListNode{
	//					Val:  5,
	//					Next: nil,
	//				},
	//			},
	//		},
	//		{
	//			Val: 1,
	//			Next: &ListNode{
	//				Val: 3,
	//				Next: &ListNode{
	//					Val:  4,
	//					Next: nil,
	//				},
	//			},
	//		},
	//		{
	//			Val: 2,
	//			Next: &ListNode{
	//				Val:  6,
	//				Next: nil,
	//			},
	//		},
	//	},
	//))
	fmt.Printf("%v\n", mergeKLists(
		[]*ListNode{
			{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
			{
				Val: 4,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val:  6,
						Next: nil,
					},
				},
			},
		},
	))
}
