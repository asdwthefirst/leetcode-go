/**
  @author: jiangxi
  @since: 2022/12/28
  @desc: //TODO
**/
package binary

func searchRange(nums []int, target int) (ret []int) {
	left, right := 0, len(nums)-1
	//找第一个
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if len(nums) == 0 || nums[left] != target {
		return []int{-1, -1}
	}
	ret = append(ret, left)

	right = len(nums) - 1
	//找最后一个
	for left < right {
		mid := left + (right-left)/2 + 1
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid
		}
	}
	ret = append(ret, left)
	return

}
