package lists

type SinglyLinkedNode struct {
	Val  int
	Next *SinglyLinkedNode
}

func MakeSinglyLinkedList(numbers []int) *SinglyLinkedNode {
	list := new(SinglyLinkedNode)
	current := list
	for _, e := range numbers {
		current.Next = &SinglyLinkedNode{Val: e, Next: nil}
		current = current.Next
	}
	return list.Next
}

func CompareSinglyLinkedLists(
	list1 *SinglyLinkedNode,
	list2 *SinglyLinkedNode) bool {
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
