/**
  @author: jiangxi
  @since: 2022/11/3
  @desc: //TODO
**/
package recursion

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if root == nil {
		return false
	}

	return isSubtreeWithRoot(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)

}

func isSubtreeWithRoot(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if root == nil || subRoot == nil {
		return false
	}
	if root.Val != subRoot.Val {
		return false
	}
	return isSubtreeWithRoot(root.Left, subRoot.Left) && isSubtreeWithRoot(root.Right, subRoot.Right)
}
