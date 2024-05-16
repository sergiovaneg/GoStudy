package p0234_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/LeetCode/p0234"
	lists "github.com/sergiovaneg/GO_leetcode/Lists"
)

func TestIsPalindrome(t *testing.T) {
	makeList := lists.MakeSinglyLinkedList[int]
	if !p0234.IsPalindrome(makeList([]int{1, 2, 2, 1})) {
		t.Fatal()
	}
	if p0234.IsPalindrome(makeList([]int{1, 2})) {
		t.Fatal()
	}
	if !p0234.IsPalindrome(makeList([]int{1})) {
		t.Fatal()
	}
}
