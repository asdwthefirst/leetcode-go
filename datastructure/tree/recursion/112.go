/**
  @author: jiangxi
  @since: 2022/11/2
  @desc: //TODO
**/
package recursion

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == targetSum {
			return true
		}
		return false
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
