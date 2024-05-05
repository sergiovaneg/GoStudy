package p0237

import lists "github.com/sergiovaneg/GO_leetcode/Lists"

type ListNode = lists.SinglyLinkedNode[int]

func DeleteNode(node *ListNode) {
	aux := node.Next
	if aux == nil {
		node = nil
		return
	}

	node.Val = aux.Val
	node.Next = aux.Next
	aux = nil
}
