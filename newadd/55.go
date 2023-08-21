/*
*

	@author: jiangxi
	@since: 2023/7/27
	@desc: //TODO

*
*/
package newadd

// 尝试贪心覆盖法。
// 确实更为简洁，但是其实没看出来贪心在哪里。
// 不断扩充可以到达的边界，同时在边界内探索。很简洁。
func canJump(nums []int) bool {
	cover := 0
	for i := 0; i <= cover; i++ {
		if nums[i]+i > cover {
			cover = nums[i] + i
		}
		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}
