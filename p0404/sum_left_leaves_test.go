package p0404_test

import (
	"testing"

	binarytree "github.com/sergiovaneg/GO_leetcode/BinaryTree"
	"github.com/sergiovaneg/GO_leetcode/p0404"
)

type TreeNode = binarytree.TreeNode[int]

func TestSumOfLeftLeaves(t *testing.T) {
	var res int
	var root *TreeNode

	root = binarytree.MakeBinaryTree([]int{0, 2, 4, 1, binarytree.NullInt, 3, -1, 5, 1, binarytree.NullInt, 6, binarytree.NullInt, 8}, binarytree.NullInt)
	res = p0404.SumOfLeftLeaves(root)
	if res != 5 {
		t.Fatalf("Expected 5; got %v", res)
	}

	root = binarytree.MakeBinaryTree([]int{0, -4, -3, binarytree.NullInt, -1, 8, binarytree.NullInt, binarytree.NullInt, 3, binarytree.NullInt, -9, -2, binarytree.NullInt, 4}, binarytree.NullInt)
	res = p0404.SumOfLeftLeaves(root)
	if res != 2 {
		t.Fatalf("Expected 2; got %v", res)
	}

	root = binarytree.MakeBinaryTree([]int{1}, binarytree.NullInt)
	res = p0404.SumOfLeftLeaves(root)
	if res != 0 {
		t.Fatalf("Expected 0; got %v", res)
	}
}
