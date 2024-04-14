package p1669_test

import (
	"testing"

	lists "github.com/sergiovaneg/GO_leetcode/Lists"
	"github.com/sergiovaneg/GO_leetcode/p1669"
)

func TestMergeInBetween(t *testing.T) {
	makeList := lists.MakeSinglyLinkedList[int]
	compareLists := lists.CompareSinglyLinkedLists[int]

	list1 := makeList([]int{10, 1, 13, 6, 9, 5})
	list2 := makeList([]int{1000000, 1000001, 1000002})
	list3 := makeList([]int{10, 1, 13, 1000000, 1000001, 1000002, 5})
	if !compareLists(p1669.MergeInBetween(list1, 3, 4, list2), list3) {
		t.Fatal()
	}
}
