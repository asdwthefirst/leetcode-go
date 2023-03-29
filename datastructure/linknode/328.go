/**
  @author: jiangxi
  @since: 2022/11/2
  @desc: //TODO
**/
package linknode

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre1 := head
	pre2 := head.Next
	l1, l2 := pre1, pre2
	for l1.Next != nil || l2.Next != nil {
		if l1.Next != nil {
			l1.Next = l1.Next.Next
		}
		if l2.Next != nil {
			l2.Next = l2.Next.Next
		}
		if l1.Next != nil {
			l1 = l1.Next
		}
		if l2.Next != nil {
			l2 = l2.Next
		}
	}
	l1.Next = pre2
	return pre1

}
