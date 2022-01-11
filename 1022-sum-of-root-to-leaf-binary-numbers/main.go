package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return getSum(root, 0)
}

func getSum(child *TreeNode, prevSum int) int {
	if child == nil {
		return 0
	}

	sum := prevSum*2 + child.Val
	if child.Left == nil && child.Right == nil {
		return sum
	}

	return getSum(child.Left, sum) + getSum(child.Right, sum)
}

func main() {
	fmt.Printf("%d\n", sumRootToLeaf(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
	}))
}
