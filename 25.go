/**
  @author: jiangxi
  @since: 2022/10/31
  @desc: //TODO
**/
package leetcode_go

import (
	"leetcode-go/datastructure/linknode"
)

func reverseKGroup(head *linknode.ListNode, k int) *linknode.ListNode {
	hair := &linknode.ListNode{Next: head}
	pre := hair

	for pre.Next != nil {
		head = pre.Next
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		nex := tail.Next
		head, tail = reverseInner(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = nex
	}
	return hair.Next
}

func reverseInner(head, tail *linknode.ListNode) (*linknode.ListNode, *linknode.ListNode) {
	if head == tail {
		return head, head
	}
	newHead, newTail := reverseInner(head.Next, tail)
	newTail.Next = head
	head.Next = nil
	return newHead, head
}
