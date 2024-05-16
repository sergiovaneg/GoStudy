package p0404

import binarytree "github.com/sergiovaneg/GoStudy/BinaryTree"

type TreeNode = binarytree.TreeNode[int]

func recursiveLeftSum(branch *TreeNode) int {
	if branch.Left == nil && branch.Right == nil {
		return branch.Val
	}

	var val int
	if branch.Left != nil {
		val = recursiveLeftSum(branch.Left)
	}
	if branch.Right != nil {
		if branch.Right.Left != nil || branch.Right.Right != nil {
			val += recursiveLeftSum(branch.Right)
		}
	}
	return val
}

func SumOfLeftLeaves(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 0
	}

	return recursiveLeftSum(root)
}
