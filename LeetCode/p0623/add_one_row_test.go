package p0623_test

import (
	"testing"

	binarytree "github.com/sergiovaneg/GoStudy/BinaryTree"
	"github.com/sergiovaneg/GoStudy/LeetCode/p0623"
)

type TreeNode = binarytree.TreeNode[int]

func TestAddOneRow(t *testing.T) {
	null := binarytree.NullInt
	var expected, res *TreeNode

	res = p0623.AddOneRow(binarytree.MakeBinaryTree(
		[]int{4, 2, 6, 3, 1, 5}, null), 1, 2)
	expected = binarytree.MakeBinaryTree(
		[]int{4, 1, 1, 2, null, null, 6, 3, 1, 5}, null)
	if !binarytree.CompareBinaryTree(res, expected) {
		t.Fatalf("Expected:\n%v\nGot:\n%v",
			binarytree.PrintBinaryTree(expected),
			binarytree.PrintBinaryTree(res))
	}

	t.Log("\n" + binarytree.PrintBinaryTree(res))

	res = p0623.AddOneRow(binarytree.MakeBinaryTree(
		[]int{4, 2, null, 3, 1}, null), 1, 3)
	expected = binarytree.MakeBinaryTree(
		[]int{4, 2, null, 1, 1, 3, null, null, 1}, null)
	if !binarytree.CompareBinaryTree(res, expected) {
		t.Fatalf("Expected:\n%v\nGot:\n%v",
			binarytree.PrintBinaryTree(expected),
			binarytree.PrintBinaryTree(res))
	}
}
