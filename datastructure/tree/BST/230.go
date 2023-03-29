/**
  @author: jiangxi
  @since: 2022/11/10
  @desc: //TODO
**/
package BST

func kthSmallest(root *TreeNode, k int) int {
	node := root
	stack := []*TreeNode{}
	var in []int
	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		in = append(in, stack[len(stack)-1].Val)
		stack = stack[:len(stack)-1]
		node = node.Right
	}
	return in[k-1]

}
