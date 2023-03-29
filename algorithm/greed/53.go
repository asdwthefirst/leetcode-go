/**
  @author: jiangxi
  @since: 2022/12/28
  @desc: //TODO
**/
package greed

//不需要左右边界所以不用考虑更换边界，只需要记录sum
func maxSubArray(nums []int) int {
	Sum := 0
	maxSum := nums[0]
	for i := 0; i < len(nums); i++ {
		if Sum > 0 {
			Sum = Sum + nums[i]
		} else {
			Sum = nums[i]
		}
		if Sum > maxSum {
			maxSum = Sum
		}
	}
	return maxSum

}
