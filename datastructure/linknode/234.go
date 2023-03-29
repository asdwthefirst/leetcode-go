/**
  @author: jiangxi
  @since: 2022/11/2
  @desc: //TODO
**/
package linknode

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow := head
	fast := head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	l2 := reverseList234(slow.Next)
	slow.Next = nil
	return isEqual(head, l2)

}

func reverseList234(head *ListNode) *ListNode {
	newHair := &ListNode{}
	for head != nil {
		nex := head.Next
		head.Next = newHair.Next
		newHair.Next = head
		head = nex
	}
	return newHair.Next
}

func isEqual(l1, l2 *ListNode) bool {
	if l1 == nil || l2 == nil {
		if l1 == nil && l2 == nil {
			return true
		} else {
			return false
		}
	}
	for l1 != nil && l2 != nil {
		if l1 == nil || l2 == nil || l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}
