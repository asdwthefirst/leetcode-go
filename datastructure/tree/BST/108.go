/**
  @author: jiangxi
  @since: 2022/11/11
  @desc: //TODO
**/
package BST

func sortedArrayToBST(nums []int) *TreeNode {
	return ToBST(nums, 0, len(nums)-1)

}

func ToBST(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	node := &TreeNode{Val: nums[mid]}
	node.Left = ToBST(nums, left, mid-1)
	node.Right = ToBST(nums, mid+1, right)
	return node
}
