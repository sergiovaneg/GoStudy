package p0404_test

import (
	"testing"

	binarytree "github.com/sergiovaneg/GO_leetcode/BinaryTree"
	"github.com/sergiovaneg/GO_leetcode/p0404"
)

type TreeNode = binarytree.TreeNode[int]

func TestSumOfLeftLeaves(t *testing.T) {
	const NullInt = -int(^uint(0)>>1) + 1

	var res int
	var root *TreeNode

	root = binarytree.MakeBinaryTree([]int{0, 2, 4, 1, NullInt, 3, -1, 5, 1, NullInt, 6, NullInt, 8}, NullInt)
	res = p0404.SumOfLeftLeaves(root)
	if res != 5 {
		t.Fatalf("Expected 5; got %v", res)
	}

	root = binarytree.MakeBinaryTree([]int{0, -4, -3, NullInt, -1, 8, NullInt, NullInt, 3, NullInt, -9, -2, NullInt, 4}, NullInt)
	res = p0404.SumOfLeftLeaves(root)
	if res != 2 {
		t.Fatalf("Expected 2; got %v", res)
	}

	root = binarytree.MakeBinaryTree([]int{1}, NullInt)
	res = p0404.SumOfLeftLeaves(root)
	if res != 0 {
		t.Fatalf("Expected 0; got %v", res)
	}
}
