/*
*

	@author: jiangxi
	@since: 2023/8/27
	@desc: //TODO

*
*/
package newadd

func majorityElement(nums []int) int {
	var candidate, count int
	for i := range nums {
		if count == 0 {
			candidate = nums[i]
		}
		if nums[i] == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
