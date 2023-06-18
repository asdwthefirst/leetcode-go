/*
*

	@author: jiangxi
	@since: 2022/11/11
	@desc: //TODO

*
*/
package BST

import "fmt"

var curCount int
var maxCount int
var Modes []int

//var countMap map[int]int

func findMode(root *TreeNode) []int {
	Modes = []int{}
	maxCount = -1
	curCount = 0
	inOrder501(root, &[]int{})
	return Modes
}

func inOrder501(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	inOrder501(root.Left, nums)
	fmt.Println(nums, root.Val)
	fmt.Println(curCount, maxCount)
	if len(*nums) > 0 && root.Val == (*nums)[len(*nums)-1] { //0618这一步的处理忘了，每遍历到一个节点都count
		curCount++
	} else {
		curCount = 1
	}
	if curCount > maxCount {
		Modes = Modes[:0]
		Modes = append(Modes, root.Val)
		maxCount = curCount
	} else if curCount == maxCount {
		Modes = append(Modes, root.Val)
	}

	*nums = append(*nums, root.Val)
	inOrder501(root.Right, nums)

}
