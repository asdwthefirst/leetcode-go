/**
  @author: jiangxi
  @since: 2022/12/30
  @desc: //TODO
**/
package backtracking

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) (ret []string) {
	if root == nil {
		return
	}
	backtrack257(root, &ret, "")
	return

}

func backtrack257(root *TreeNode, ret *[]string, path string) {
	path += "->" + fmt.Sprint(root.Val)
	if root.Left == nil && root.Right == nil {
		*ret = append(*ret, path)
	}
	if root.Left != nil {
		backtrack257(root.Left, ret, path)
	}
	if root.Right != nil {
		backtrack257(root.Right, ret, path)
	}
}
