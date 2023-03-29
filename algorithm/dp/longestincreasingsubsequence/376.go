/**
  @author: jiangxi
  @since: 2023/1/4
  @desc: //TODO
**/
package longestincreasingsubsequence

//因为一段相邻的相同元素中我们最多只能选择其中的一个，所以我们可以忽略相邻的相同元素。
//现在我们假定序列中任意两个相邻元素都不相同，即要么左侧大于右侧，要么右侧大于左侧。
//对于序列中既非「峰」也非「谷」的元素，我们称其为「过渡元素」。
//如序列 [1,2,3,4][1,2,3,4][1,2,3,4] 中，222 和 333 都是「过渡元素」。

//我且记住这个专题单调连接使用的是包含i的dp，不单调连接使用的是不一定包含i的dp
//这里的思路是到此为止的最长序列，而不是包含自己的最长序列，也妙哉,转移方程的设立用的假设带入替换的思想，妙哉
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
