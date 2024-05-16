package p0623

import binarytree "github.com/sergiovaneg/GoStudy/BinaryTree"

type TreeNode = binarytree.TreeNode[int]

func traverseAndInsert(root *TreeNode, val_addr *int, depth int) {
	if depth == 1 {
		root.Left = &TreeNode{Val: *val_addr, Left: root.Left, Right: nil}
		root.Right = &TreeNode{Val: *val_addr, Left: nil, Right: root.Right}
	} else {
		if root.Left != nil {
			traverseAndInsert(root.Left, val_addr, depth-1)
		}
		if root.Right != nil {
			traverseAndInsert(root.Right, val_addr, depth-1)
		}
	}
}

func AddOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{Val: val, Left: root, Right: nil}
	}
	traverseAndInsert(root, &val, depth-1)
	return root
}
