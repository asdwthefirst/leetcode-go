/**
  @author: jiangxi
  @since: 2023/1/3
  @desc: //TODO
**/
package fibonacci

func rob(nums []int) int {
	return max198(robRange(nums, 0, len(nums)-2), robRange(nums, 1, len(nums)-1))
}

func robRange(nums []int, start, end int) int {
	pre1, pre2 := 0, 0
	cur := 0
	for i := start; i <= end; i++ {
		cur = max198(nums[i]+pre1, pre2)
		pre1, pre2 = pre2, cur
	}
	return max198(pre2, cur)
}

func max213(a, b int) int {
	if a > b {
		return a
	}
	return b
}
