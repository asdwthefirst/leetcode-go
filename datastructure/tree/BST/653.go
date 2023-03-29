/**
  @author: jiangxi
  @since: 2022/11/11
  @desc: //TODO
**/
package BST

func findTarget(root *TreeNode, k int) bool {
	var nums []int
	inOrder653(root, &nums)
	numMap := make(map[int]bool)
	for _, num := range nums {
		if _, ok := numMap[k-num]; ok {
			return true
		}
		numMap[num] = true
	}
	return false

}

func inOrder653(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	inOrder653(root.Left, nums)
	*nums = append(*nums, root.Val)
	inOrder653(root.Right, nums)

}
