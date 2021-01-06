package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	last := dummy.Next

	for a, b, add := 0, 0, 0; l1 != nil && l2 != nil; {
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		} else {
			a = 0
		}

		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		} else {
			b = 0
		}

		sum := a + b + add

		last = new(ListNode)
		last.Val = sum % 10
		last = last.Next

		add = sum / 10
	}

	return dummy.Next
}

/*
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/
