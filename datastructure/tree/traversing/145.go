/**
  @author: jiangxi
  @since: 2022/11/10
  @desc: //TODO
**/
package traversing

//
func postorderTraversal(root *TreeNode) []int {
	var vals []int
	var result []int
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			vals = append(vals, node.Val)
			stack = append(stack, node)
			node = node.Right
		}
		node = stack[len(stack)-1].Left
		stack = stack[:len(stack)-1]
	}
	for i := len(vals) - 1; i >= 0; i-- {
		result = append(result, vals[i])
	}
	return result

}
