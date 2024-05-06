package p2487

import lists "github.com/sergiovaneg/GO_leetcode/Lists"

type ListNode = lists.SinglyLinkedNode[int]

func recRemove(node *ListNode) *ListNode {
	if node.Next == nil {
		return node
	}

	node.Next = recRemove(node.Next)
	if node.Next.Val > node.Val {
		return node.Next
	}

	return node
}

func RemoveNodes(head *ListNode) *ListNode {
	return recRemove(head)
}
