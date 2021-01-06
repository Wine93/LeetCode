package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	/*prev := nil*/

	var prev *ListNode
	slow, fast := head, head

	for {
		for i := 1; i <= n; i++ {
			fast = fast.Next
		}

		if fast == nil {
			break
		}

		prev = slow
		slow = slow.Next
	}

	if prev == nil {
		head = head.Next
	} else {
		prev.Next = prev.Next.Next
	}

	return head
}

/*
给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
*/
