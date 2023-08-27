/*
*

	@author: jiangxi
	@since: 2023/8/27
	@desc: //TODO

*
*/
package newadd

//使用异或思想

func singleNumber(nums []int) int {
	ans := nums[0]
	if len(nums) > 1 {
		for i := 1; i < len(nums); i++ {
			ans = ans ^ nums[i]
		}
	}
	return ans
}
