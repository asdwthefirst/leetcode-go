/**
  @author: jiangxi
  @since: 2022/10/31
  @desc: //TODO
**/
package linknode

func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	hair := &ListNode{}
	for head != nil {
		nex := head.Next
		head.Next = hair.Next
		hair.Next = head
		head = nex
	}
	return hair.Next
}
