/**
  @author: jiangxi
  @since: 2022/11/10
  @desc: //TODO
**/
package traversing

func inorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	node := root
	var result []int
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		result = append(result, node.Val)
		stack = stack[:len(stack)-1]
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return result

}
