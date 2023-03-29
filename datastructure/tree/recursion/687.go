/**
  @author: jiangxi
  @since: 2022/11/3
  @desc: //TODO
**/
package recursion

var max int

func longestUnivaluePath(root *TreeNode) int {
	max = 0
	dfs(root)
	return max

}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	leftPath, rightPath := 0, 0

	if root.Left != nil && root.Left.Val == root.Val {
		leftPath = left + 1
	}
	if root.Right != nil && root.Right.Val == root.Val {
		rightPath = right + 1
	}
	if leftPath+rightPath > max {
		max = leftPath + rightPath
	}
	if leftPath > rightPath {
		return leftPath
	}
	return rightPath

}
