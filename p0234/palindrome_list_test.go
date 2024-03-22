package p0234_test

import (
	"testing"

	lists "github.com/sergiovaneg/GO_leetcode/Lists"
	"github.com/sergiovaneg/GO_leetcode/p0234"
)

func TestIsPalindrome(t *testing.T) {
	if !p0234.IsPalindrome(lists.MakeSinglyLinkedList([]int{1, 2, 2, 1})) {
		t.Fatal()
	}
	if p0234.IsPalindrome(lists.MakeSinglyLinkedList([]int{1, 2})) {
		t.Fatal()
	}
	if !p0234.IsPalindrome(lists.MakeSinglyLinkedList([]int{1})) {
		t.Fatal()
	}
}
