package solution

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var odd, even *ListNode

	for n := 1; head != nil; head = head.Next {
		if n&1 == 1 {
			if odd == nil {
				odd = head
			} else {
				odd.Next = head
			}
		} else {

		}
	}

	return head
}
