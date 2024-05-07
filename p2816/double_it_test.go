package p2816_test

import (
	"testing"

	lists "github.com/sergiovaneg/GO_leetcode/Lists"
	"github.com/sergiovaneg/GO_leetcode/p2816"
)

type ListNode = lists.SinglyLinkedNode[int]

func TestDoubleIt(t *testing.T) {
	var res, exp *ListNode

	res = p2816.DoubleIt(lists.MakeSinglyLinkedList([]int{1, 8, 9}))
	exp = lists.MakeSinglyLinkedList([]int{3, 7, 8})
	if !lists.CompareSinglyLinkedLists(res, exp) {
		t.Fatalf("Expected %v; got %v", exp, res)
	}

	res = p2816.DoubleIt(lists.MakeSinglyLinkedList([]int{9, 9, 9}))
	exp = lists.MakeSinglyLinkedList([]int{1, 9, 9, 8})
	if !lists.CompareSinglyLinkedLists(res, exp) {
		t.Fatalf("Expected %v; got %v", exp, res)
	}
}
