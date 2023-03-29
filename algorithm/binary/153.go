/**
  @author: jiangxi
  @since: 2022/12/28
  @desc: //TODO
**/
package binary

func findMinBad(nums []int) int {
	left, right := 0, len(nums)-1
	if nums[left] < nums[right] {
		return nums[left]
	}

	for left < right {
		mid := left + (right-left)/2
		if mid == left {
			if nums[left] < nums[right] {
				return nums[left]
			} else {
				return nums[right]
			}
		} else if nums[mid] > nums[left] {
			if nums[mid+1] < nums[mid] { //特殊情况
				return nums[mid+1]
			}
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

//因为是找最左边界，所以right=mid，所以不能是mid=（left+right）/2是等于right的，会陷入循环
//因为mid不会等于right，所以跟right比较会少特殊情况
//而且直观考虑，上面和left比较会导致的截出来已经是递增的数组，导致特殊情况，而对特殊情况，上面的判断会失效，而下面的判断依然能正确的判断，但是还是没有找出共性
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}
