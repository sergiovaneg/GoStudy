package binarytree

import "slices"

type TreeNode[T any] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func countNulls[T comparable](elements []T, null T) int {
	var acc int
	for _, e := range elements {
		if e == null {
			acc++
		}
	}
	return acc
}

func MakeBinaryTree[T comparable](elements []T, null T) *TreeNode[T] {
	l := len(elements)
	if l == 0 || elements[0] == null {
		return nil
	}

	root := new(TreeNode[T])
	root.Val = elements[0]

	block_size := 1
	var left_skips, right_skips int
	left := make([]T, 0, l)
	right := make([]T, 0, l)

	for idx := 1; idx < l; {
		limit := min(idx+block_size-left_skips, l)
		left = append(left, elements[idx:limit]...)
		left_skips += countNulls(elements[idx:limit], null)
		idx = limit

		if idx == l {
			break
		}

		limit = min(idx+block_size-right_skips, l)
		right = append(right, elements[idx:limit]...)
		right_skips += countNulls(elements[idx:limit], null)
		idx = limit

		block_size *= 2
		left_skips *= 2
		right_skips *= 2
	}

	// Memory optimization
	elements = nil
	left = slices.Clip(left)
	right = slices.Clip(right)

	root.Left = MakeBinaryTree(left, null)
	root.Right = MakeBinaryTree(right, null)

	return root
}
