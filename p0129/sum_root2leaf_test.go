package p0129_test

import (
	"testing"

	binarytree "github.com/sergiovaneg/GO_leetcode/BinaryTree"
	"github.com/sergiovaneg/GO_leetcode/p0129"
)

func TestSumNumbers(t *testing.T) {
	const NullInt = -int(^uint(0)>>1) + 1
	var res int

	MakeBinaryTree := binarytree.MakeBinaryTree[int]

	res = p0129.SumNumbers(MakeBinaryTree([]int{1, 2, 3}, NullInt))
	if res != 25 {
		t.Fatalf("Expected 25; got %v", res)
	}

	res = p0129.SumNumbers(MakeBinaryTree([]int{4, 9, 0, 5, 1}, NullInt))
	if res != 1026 {
		t.Fatalf("Expected 1026; got %v", res)
	}
}
