/**
  @author: jiangxi
  @since: 2022/12/28
  @desc: //TODO
**/
package divideandconquer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	return generateTreesBetween(1, n)

}

func generateTreesBetween(left, right int) []*TreeNode {
	ret := []*TreeNode{}
	if left > right {
		return ret
	}
	if left == right {
		ret = append(ret, &TreeNode{Val: left})
	}
	rightTrees := generateTreesBetween(left+1, right)
	for _, rightTree := range rightTrees {
		node := &TreeNode{
			Val:   left,
			Right: rightTree,
		}
		ret = append(ret, node)

	}
	leftTrees := generateTreesBetween(left, right-1)
	for _, leftTree := range leftTrees {
		node := &TreeNode{
			Val:  right,
			Left: leftTree,
		}
		ret = append(ret, node)

	}
	for i := left + 1; i < right; i++ {
		leftTrees := generateTreesBetween(left, i-1)
		rightTrees := generateTreesBetween(i+1, right)
		for _, leftTree := range leftTrees {
			for _, rightTree := range rightTrees {
				node := &TreeNode{
					Val: i,
				}
				node.Left = leftTree
				node.Right = rightTree
				ret = append(ret, node)
			}
		}
	}
	return ret

}