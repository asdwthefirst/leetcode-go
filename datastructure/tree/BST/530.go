/**
  @author: jiangxi
  @since: 2022/11/11
  @desc: //TODO
**/
package BST

func getMinimumDifference(root *TreeNode) int {
	var nums []int
	inOrder(root, &nums)
	min := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] < min {
			min = nums[i] - nums[i-1]
		}
	}
	return min
}

func inOrder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, nums)
	*nums = append(*nums, root.Val)
	inOrder(root.Right, nums)

}
