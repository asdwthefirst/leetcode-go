/**
  @author: jiangxi
  @since: 2022/11/2
  @desc: //TODO
**/
package recursion

/*time exceed
var count int

func pathSum(root *TreeNode, targetSum int) int {
	count = 0
	pathSumAll(root, targetSum)
	return count
}

func pathSumAll(root *TreeNode, targetSum int) {
	if root != nil {
		fmt.Println("in2:", root.Val)
		pathSumAll(root.Left, targetSum)
		pathSumWithRoot(root, targetSum)
		pathSumAll(root.Right, targetSum)
	}
}

func pathSumWithRoot(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	fmt.Println("in:", root.Val)
	left := pathSumWithRoot(root.Left, targetSum-root.Val)
	right := pathSumWithRoot(root.Right, targetSum-root.Val)
	if root.Val == targetSum {
		fmt.Println(root.Val)
		count++
		return true
	}
	return left || right
}

*/

//这就没超时，但是不知道为什么
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return pathSumWithRoot(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

func pathSumWithRoot(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	if root.Val == targetSum {
		return 1 + pathSumWithRoot(root.Left, targetSum-root.Val) + pathSumWithRoot(root.Right, targetSum-root.Val)
	}
	return pathSumWithRoot(root.Left, targetSum-root.Val) + pathSumWithRoot(root.Right, targetSum-root.Val)

}
