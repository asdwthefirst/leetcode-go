/**
  @author: jiangxi
  @since: 2022/11/10
  @desc: //TODO
**/
package BST

//我怎么这么笨，分在左，在右，和夹在中间，中间用排除法就好了，没必要区分哪个大哪个小
func lowestCommonAncestor235(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	//var big, small int
	//if p.Val >= q.Val {
	//	big = p.Val
	//	small = q.Val
	//} else {
	//	small = q.Val
	//	big = p.Val
	//}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
