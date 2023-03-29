/**
  @author: jiangxi
  @since: 2022/12/27
  @desc: //TODO
**/
package greed

//真的有点绕，但反正就是有限制的改数，所以就改完使结果尽量不用再多改，看改完是不是能满足。
func checkPossibility(nums []int) bool {
	cnt := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if cnt == 0 {
				cnt++
				if i == 1 {
					nums[i-1] = nums[i]
				} else {
					if nums[i] > nums[i-2] {
						nums[i-1] = nums[i]
					} else {
						nums[i] = nums[i-1]
					}
				}
			} else {
				return false
			}
		}
	}
	return true
}
