/*
*

	@author: jiangxi
	@since: 2022/11/10
	@desc: //TODO

*
*/
package BST

// 挺高级的，我自己想不到反正，感觉上是一种自底向上的搜索方法
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil {
		if right != nil {
			return root
		}
		return right
	}
	if right != nil {
		return right
	}
	return root
}
