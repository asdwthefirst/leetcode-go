/**
  @author: jiangxi
  @since: 2022/11/10
  @desc: //TODO
**/
package BST

var last int

func convertBST(root *TreeNode) *TreeNode {
	last = 0
	convert(root)
	return root

}

func convert(root *TreeNode) {
	if root == nil {
		return
	}
	convert(root.Right)
	root.Val = last + root.Val
	last = root.Val
	convert(root.Left)
}
