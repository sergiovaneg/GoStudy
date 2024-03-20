package p1669_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/p1669"
)

func makeList(numbers []int) *p1669.ListNode {
	list := new(p1669.ListNode)
	current := list
	for _, e := range numbers {
		current.Next = &p1669.ListNode{e, nil}
		current = current.Next
	}
	return list.Next
}

func compareLists(list1 *p1669.ListNode, list2 *p1669.ListNode) bool {
	for list1.Next != nil && list2.Next != nil {
		if list1.Val != list2.Val {
			return false
		}
		list1 = list1.Next
		list2 = list2.Next
	}
	if list1.Val != list2.Val {
		return false
	}
	if list1.Next != nil || list2.Next != nil {
		return false
	}
	return true
}

func TestMergeInBetween(t *testing.T) {
	list1 := makeList([]int{10, 1, 13, 6, 9, 5})
	list2 := makeList([]int{1000000, 1000001, 1000002})
	list3 := makeList([]int{10, 1, 13, 1000000, 1000001, 1000002, 5})
	if !compareLists(p1669.MergeInBetween(list1, 3, 4, list2), list3) {
		t.Fatal()
	}
}
