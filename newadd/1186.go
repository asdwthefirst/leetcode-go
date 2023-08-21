/*
*

	@author: jiangxi
	@since: 2023/7/24
	@desc: //TODO

*
*/

package newadd

//redstar第三题应该是这样的 ，但是我当时没想到。

// 思路，dp1[i]保存，到i为止，已经删除了数的最大值，dp2[i]保存，到i为止，未删除数的最大值
func maximumSum(arr []int) int {
	var res int
	if len(arr) == 0 {
		return -1
	}
	var dp1, dp2 int
	dp2 = arr[0]
	res = arr[0]
	dp1 = arr[0]

	for i := 1; i < len(arr); i++ {
		dp1, dp2 = max(dp1+arr[i], arr[i], dp2), max(dp2+arr[i], arr[i])
		if dp1 > res {
			res = dp1
		}
		if dp2 > res {
			res = dp2
		}
		// fmt.Println(i, dp1, dp2)
	}
	return res
}

func max(nums ...int) int {
	maxN := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxN {
			maxN = nums[i]
		}

	}
	return maxN
}
