package p0234_test

import (
	"testing"

	lists "github.com/sergiovaneg/GO_leetcode/Lists"
	"github.com/sergiovaneg/GO_leetcode/p0234"
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
