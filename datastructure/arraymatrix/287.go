/*
*

	@author: jiangxi
	@since: 2022/12/8
	@desc: //TODO

*
*/
package arraymatrix

// 给定一个包含 n + 1 个整数的数组 nums ，
// 其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。
// 假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。
// 要求不能修改数组，也不能使用额外的空间。
func findDuplicate(nums []int) int {
	n := len(nums) - 1
	left, right := 1, n
	ans := -1
	for left <= right {
		mid := left + (right-left)/2
		cnt := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			left = mid + 1
		} else {
			right = mid - 1
			ans = mid
		}
	}
	return ans
}
