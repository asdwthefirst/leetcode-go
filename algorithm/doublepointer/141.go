/**
  @author: jiangxi
  @since: 2022/12/19
  @desc: //TODO
**/
package doublepointer

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	pa, pb := head, head.Next
	for pb != nil && pb.Next != nil {
		if pa == pb {
			return true
		}
		pa = pa.Next
		pb = pb.Next.Next

	}
	return false
}
