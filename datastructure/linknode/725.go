/**
  @author: jiangxi
  @since: 2022/11/2
  @desc: //TODO
**/
package linknode

func splitListToParts(head *ListNode, k int) []*ListNode {
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}
	quotient, remainder := n/k, n%k
	parts := make([]*ListNode, k)
	curr := head
	for i := 0; i < k && curr != nil; i++ {
		parts[i] = curr
		partSize := quotient
		if i < remainder {
			partSize++
		}
		for i := 0; i < partSize-1; i++ {
			curr = curr.Next
		}
		curr, curr.Next = curr.Next, nil
	}
	return parts
}
