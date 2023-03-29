/**
  @author: jiangxi
  @since: 2023/1/1
  @desc: //TODO
**/
package fibonacci

func rob198(nums []int) int {
	pre1, pre2 := 0, 0
	cur := 0
	for i := 0; i < len(nums); i++ {
		cur = max198(nums[i]+pre1, pre2)
		pre1, pre2 = pre2, cur

	}
	return max198(pre2, cur)
}

func max198(a, b int) int {
	if a > b {
		return a
	}
	return b
}
