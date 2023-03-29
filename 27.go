/**
  @author: jiangxi
  @since: 2022/10/31
  @desc: //TODO
**/
package leetcode_go

func removeElement(nums []int, val int) int {
	p, q := 0, 0
	if len(nums) == 0 {
		return 0
	}
	for q < len(nums) {
		for q < len(nums) && nums[q] == val {
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
