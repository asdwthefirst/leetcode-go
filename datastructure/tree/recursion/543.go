/**
  @author: jiangxi
  @since: 2022/11/2
  @desc: //TODO
**/
package recursion

func diameterOfBinaryTree(root *TreeNode) int {
	result := 0

	var maxDepth110 func(root *TreeNode) int
	maxDepth110 = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := maxDepth110(root.Left)
		right := maxDepth110(root.Right)
		if left+right > result {
			result = left + right
		}
		if left > right {
			return left + 1
		}
		return right + 1
	}
	maxDepth110(root)
	return result
}
