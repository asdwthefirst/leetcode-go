/**
  @author: jiangxi
  @since: 2022/11/2
  @desc: //TODO
**/
package recursion

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
