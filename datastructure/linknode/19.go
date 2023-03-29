/**
  @author: jiangxi
  @since: 2022/11/2
  @desc: //TODO
**/
package linknode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair
	tail := head
	i := 0
	for ; tail != nil && i < n; i++ {
		tail = tail.Next
	}
	if i < n-1 {
		return head
	}
	for tail != nil {
		pre = pre.Next
		tail = tail.Next
	}
	pre.Next = pre.Next.Next
	return hair.Next

}
