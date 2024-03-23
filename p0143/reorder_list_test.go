package p0143_test

import (
	"testing"

	lists "github.com/sergiovaneg/GO_leetcode/Lists"
	"github.com/sergiovaneg/GO_leetcode/p0143"
)

type ListNode = lists.SinglyLinkedNode

func TestReorderList(t *testing.T) {
	list_1 := lists.MakeSinglyLinkedList([]int{1, 2, 3, 4, 5})
	p0143.ReorderList(list_1)
	if !lists.CompareSinglyLinkedLists(list_1,
		lists.MakeSinglyLinkedList([]int{1, 5, 2, 4, 3})) {
		t.Fatal()
	}

	list_2 := lists.MakeSinglyLinkedList([]int{1, 2, 3, 4})
	p0143.ReorderList(list_2)
	if !lists.CompareSinglyLinkedLists(list_2,
		lists.MakeSinglyLinkedList([]int{1, 4, 2, 3})) {
		t.Fatal()
	}
}
