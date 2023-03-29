/**
  @author: jiangxi
  @since: 2022/12/8
  @desc: //TODO
**/
package arraymatrix

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
