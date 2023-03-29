/**
  @author: jiangxi
  @since: 2022/11/3
  @desc: //TODO
**/
package recursion

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	left, right := minDepth(root.Left), minDepth(root.Right)
	if left == 0 {
		left = right + 1
	}
	if right == 0 {
		right = left + 1
	}
	if left < right {
		return left + 1
	}
	return right + 1

}
