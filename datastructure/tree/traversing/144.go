/**
  @author: jiangxi
  @since: 2022/11/3
  @desc: //TODO
**/
package traversing

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal1(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	stack := []*TreeNode{root}
	pos := 0
	for pos != -1 {
		curNode := stack[pos]
		result = append(result, curNode.Val)
		pos--
		if curNode.Right != nil {
			pos++
			if len(stack) < pos+1 {
				stack = append(stack, curNode.Right)
			} else {
				stack[pos] = curNode.Right
			}
		}
		if curNode.Left != nil {
			pos++
			if len(stack) < pos+1 {
				stack = append(stack, curNode.Left)
			} else {
				stack[pos] = curNode.Left
			}
		}
	}
	return result
}

func preorderTraversal(root *TreeNode) (vals []int) {
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			vals = append(vals, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return
}
