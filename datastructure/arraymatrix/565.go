/**
  @author: jiangxi
  @since: 2022/12/15
  @desc: //TODO
**/
package arraymatrix

func arrayNesting(nums []int) int {
	vis := make([]bool, len(nums))
	var ans int
	for i := 0; i < len(nums); i++ {
		cnt := 0
		for !vis[i] {
			vis[i] = true
			i = nums[i]
			cnt++
		}
		if cnt > ans {
			ans = cnt
		}

	}
	return ans
}