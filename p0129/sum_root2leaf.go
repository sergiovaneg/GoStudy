package p0129

import binarytree "github.com/sergiovaneg/GO_leetcode/BinaryTree"

type TreeNode = binarytree.TreeNode[int]

func growAndAdd(root *TreeNode, acc *int, num int) {
	num = num*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*acc += num
	} else {
		if root.Left != nil {
			growAndAdd(root.Left, acc, num)
		}
		if root.Right != nil {
			growAndAdd(root.Right, acc, num)
		}
	}
}

func SumNumbers(root *TreeNode) int {
	acc := new(int)
	var num int

	growAndAdd(root, acc, num)

	return *acc
}
