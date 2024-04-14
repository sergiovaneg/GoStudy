package lists

type SinglyLinkedNode[T any] struct {
	Val  T
	Next *SinglyLinkedNode[T]
}

func MakeSinglyLinkedList[T any](elements []T) *SinglyLinkedNode[T] {
	list := new(SinglyLinkedNode[T])
	current := list
	for _, e := range elements {
		current.Next = &SinglyLinkedNode[T]{Val: e, Next: nil}
		current = current.Next
	}
	return list.Next
}

func CompareSinglyLinkedLists[T comparable](
	list1 *SinglyLinkedNode[T],
	list2 *SinglyLinkedNode[T]) bool {
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
