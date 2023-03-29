/**
  @author: jiangxi
  @since: 2022/11/12
  @desc: //TODO
**/
package arraymatrix

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	curLen := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			curLen++
		} else {
			if curLen > max {
				max = curLen
			}
			curLen = 0
		}
	}
	if curLen > max {
		max = curLen
	}
	return max

}
