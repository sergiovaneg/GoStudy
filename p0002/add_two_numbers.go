package p0002

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addAndUpdate(new_tmp int, current *ListNode) int {
	if new_tmp >= 10 {
		current.Next = &ListNode{new_tmp - 10, nil}
		return 1
	} else {
		current.Next = &ListNode{new_tmp, nil}
		return 0
	}
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	current := new(ListNode)
	res := current

	tmp := 0
	for (l1 != nil) && (l2 != nil) {
		tmp = addAndUpdate(tmp+l1.Val+l2.Val, current)

		l1 = l1.Next
		l2 = l2.Next

		current = current.Next
	}
	for l1 != nil {
		tmp = addAndUpdate(tmp+l1.Val, current)
		l1 = l1.Next
		current = current.Next
	}
	for l2 != nil {
		tmp = addAndUpdate(tmp+l2.Val, current)
		l2 = l2.Next
		current = current.Next
	}
	if tmp == 1 {
		current.Next = &ListNode{1, nil}
	}

	return res.Next
}
