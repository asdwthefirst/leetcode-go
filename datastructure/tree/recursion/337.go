/**
  @author: jiangxi
  @since: 2022/11/3
  @desc: //TODO
**/
package recursion

/*超时
func rob(root *TreeNode) int {
	withRoot := robWithRoot(root)
	withoutRoot := robWithoutRoot(root)
	if withRoot > withoutRoot {
		return withRoot
	}
	return withoutRoot

}

func robWithRoot(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return robWithoutRoot(root.Left) + robWithoutRoot(root.Right) + root.Val
}

func robWithoutRoot(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := 0, 0
	with, without := robWithRoot(root.Left), robWithoutRoot(root.Left)
	if with > without {
		left = with
	} else {
		left = without
	}
	with, without = robWithRoot(root.Right), robWithoutRoot(root.Right)
	if with > without {
		right = with
	} else {
		right = without
	}
	return left + right

}

*/
var numMap = make(map[*TreeNode]int)

func rob(root *TreeNode) int {
	var num int
	if num, ok := numMap[root]; ok {
		return num
	}
	if root == nil {
		return 0
	}
	val1 := root.Val
	val2 := 0
	if root.Left != nil {
		val1 += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		val1 += rob(root.Right.Left) + rob(root.Right.Right)
	}
	val2 += rob(root.Left) + rob(root.Right)
	if val1 > val2 {
		num = val1
	} else {
		num = val2
	}
	numMap[root] = num
	return num
}
