/**
  @author: jiangxi
  @since: 2022/12/28
  @desc: //TODO
**/
package binary

//根据奇偶判断那个数在左在右
func singleNonDuplicate(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if (mid-1 < 0 || nums[mid-1] != nums[mid]) && (mid+1 >= len(nums) || nums[mid+1] != nums[mid]) {
			return nums[mid]
		}
		if mid%2 == 1 {
			if nums[mid-1] != nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid+1] != nums[mid] {
				right = mid - 2
			} else {
				left = mid + 2
			}
		}
	}
	return nums[left]

}
