/**
  @author: jiangxi
  @since: 2022/11/3
  @desc: //TODO
**/
package recursion

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricTwo(root.Left, root.Right)

}

func isSymmetricTwo(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true

	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return isSymmetricTwo(left.Left, right.Right) && isSymmetricTwo(left.Right, right.Left)
}
