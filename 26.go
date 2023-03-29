/**
  @author: jiangxi
  @since: 2022/10/31
  @desc: //TODO
**/
package leetcode_go

func removeDuplicates(nums []int) int {
	p, q := 1, 1
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	for q < len(nums) {
		for q < len(nums) && nums[q] == nums[q-1] {
			q = q + 1
		}
		if q == len(nums) {
			break
		}
		nums[p] = nums[q]
		p = p + 1
		q = q + 1
	}

	return p
}
