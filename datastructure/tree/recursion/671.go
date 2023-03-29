/**
  @author: jiangxi
  @since: 2022/11/3
  @desc: //TODO
**/
package recursion

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil || root.Left == nil {
		return -1
	}
	//if root.Left.Left == nil && root.Right.Left == nil {
	//	if root.Left.Val == root.Right.Val {
	//		return -1
	//	}
	//	if root.Left.Val > root.Right.Val {
	//		return root.Left.Val
	//	}
	//	return root.Right.Val
	//}
	left, right := root.Left.Val, root.Right.Val
	if root.Left.Val == root.Val {
		left = findSecondMinimumValue(root.Left)
	}
	if root.Right.Val == root.Val {
		right = findSecondMinimumValue(root.Right)
	}

	if left == -1 {
		return right
	}
	if right == -1 {
		return left
	}
	if left < right {
		return left
	}
	return right

	//if left != -1 && right != -1 {
	//	if left < right {
	//		return left
	//	}
	//	return right
	//}
	//if left == -1 {
	//	return right
	//}
	//return left

}
