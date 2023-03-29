/**
  @author: jiangxi
  @since: 2022/11/2
  @desc: //TODO
**/
package linknode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	pos1 := 0
	pos2 := 0
	var stack1 []int
	var stack2 []int
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
		pos1++
	}
	pos1--
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
		pos2++
	}
	pos2--
	carry := 0
	hair := &ListNode{}
	for pos1 >= 0 || pos2 >= 0 || carry != 0 {
		num1 := 0
		num2 := 0
		if pos1 >= 0 {
			num1 = stack1[pos1]
		}
		if pos2 >= 0 {
			num2 = stack2[pos2]
		}
		newPos := (carry + num1 + num2) % 10
		carry = (carry + num1 + num2) / 10
		newNode := &ListNode{Val: newPos, Next: hair.Next}
		hair.Next = newNode
		pos1--
		pos2--
	}
	return hair.Next
}
