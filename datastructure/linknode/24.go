/**
  @author: jiangxi
  @since: 2022/10/31
  @desc: //TODO
**/
package linknode

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tail := swapPairs1(head.Next.Next)
	newHead := head.Next
	newHead.Next = head
	head.Next = tail
	return newHead

}

func swapPairs(head *ListNode) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair
	for pre.Next != nil && pre.Next.Next != nil {
		l1 := pre.Next
		l2 := pre.Next.Next
		nex := pre.Next.Next.Next
		pre.Next = l2
		l2.Next = l1
		l1.Next = nex
		pre = l1
	}
	return hair.Next
}
