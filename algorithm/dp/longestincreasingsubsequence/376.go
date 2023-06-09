/*
*

	@author: jiangxi
	@since: 2023/1/4
	@desc: //TODO

*
*/
package longestincreasingsubsequence

//如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为 摆动序列 。
//第一个差（如果存在的话）可能是正数或负数。仅有一个元素或者含两个不等元素的序列也视作摆动序列。
//
//例如， [1, 7, 4, 9, 2, 5] 是一个 摆动序列 ，因为差值 (6, -3, 5, -7, 3) 是正负交替出现的。
//
//相反，[1, 4, 7, 2, 5] 和 [1, 7, 4, 5, 5] 不是摆动序列，第一个序列是因为它的前两个差值都是正数，第二个序列是因为它的最后一个差值为零。
//子序列 可以通过从原始序列中删除一些（也可以不删除）元素来获得，剩下的元素保持其原始顺序。
//
//给你一个整数数组 nums ，返回 nums 中作为 摆动序列 的 最长子序列的长度 。
//
//示例 1：
//
//输入：nums = [1,7,4,9,2,5]
//输出：6
//解释：整个序列均为摆动序列，各元素之间的差值为 (6, -3, 5, -7, 3) 。

//因为一段相邻的相同元素中我们最多只能选择其中的一个，所以我们可以忽略相邻的相同元素。
//现在我们假定序列中任意两个相邻元素都不相同，即要么左侧大于右侧，要么右侧大于左侧。
//对于序列中既非「峰」也非「谷」的元素，我们称其为「过渡元素」。
//如序列 [1,2,3,4][1,2,3,4][1,2,3,4] 中，222 和 333 都是「过渡元素」。

// 我且记住这个专题单调连接使用的是包含i的dp，不单调连接使用的是不一定包含i的dp
// 这里的思路是到此为止的最长序列，而不是包含自己的最长序列，也妙哉,转移方程的设立用的假设带入替换的思想，妙哉
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	up, down := make([]int, n), make([]int, n)
	up[0], down[0] = 1, 1
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			up[i] = up[i-1]
			down[i] = down[i-1]
		}
		if nums[i] > nums[i-1] {
			down[i] = down[i-1]
			up[i] = max376(down[i-1]+1, up[i-1])
		}
		if nums[i] < nums[i-1] {
			up[i] = up[i-1]
			down[i] = max376(up[i-1]+1, down[i-1])
		}
	}
	return max376(up[n-1], down[n-1])
}

func max376(a, b int) int {
	if a > b {
		return a
	}
	return b

}
