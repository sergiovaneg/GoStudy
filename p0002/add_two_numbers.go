package p0002

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addAndUpdate(new_temp int, current *ListNode) int {
	if new_temp >= 10 {
		current.Val = new_temp - 10
		return 1
	} else {
		current.Val = new_temp
		return 0
	}
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	current := new(ListNode)
	last := current
	res := current

	tmp := 0
	for (l1 != nil) && (l2 != nil) {
		tmp = addAndUpdate(tmp+l1.Val+l2.Val, current)

		l1 = l1.Next
		l2 = l2.Next

		current.Next = new(ListNode)
		last = current
		current = current.Next
	}

	for l1 != nil {
		tmp = addAndUpdate(tmp+l1.Val, current)

		l1 = l1.Next

		current.Next = new(ListNode)
		last = current
		current = current.Next
	}
	for l2 != nil {
		tmp = addAndUpdate(tmp+l2.Val, current)

		l2 = l2.Next

		current.Next = new(ListNode)
		last = current
		current = current.Next
	}

	if tmp == 0 {
		last.Next = nil
	} else {
		current.Val = 1
	}

	return res
}
