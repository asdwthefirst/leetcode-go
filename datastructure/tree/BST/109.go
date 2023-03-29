/**
  @author: jiangxi
  @since: 2022/11/11
  @desc: //TODO
**/
package BST

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	preMid := preMid(head)
	mid := preMid.Next
	node := &TreeNode{Val: mid.Val}
	preMid.Next = nil
	node.Left = sortedListToBST(head)
	node.Right = sortedListToBST(mid.Next)
	return node
}

func preMid(head *ListNode) *ListNode {
	hair := &ListNode{Next: head}
	slow := hair
	fast := hair
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
