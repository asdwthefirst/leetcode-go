/*
*

	@author: jiangxi
	@since: 2022/12/30
	@desc: //TODO

*
*/
package backtracking

//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
//示例 1：
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

func permute(nums []int) (ret [][]int) {
	for i := 0; i < len(nums); i++ {
		visited := make([]bool, len(nums))
		backtrack46(nums, visited, i, []int{}, &ret)
	}
	return

}

func backtrack46(nums []int, visited []bool, index int, permuted []int, ret *[][]int) {
	permuted = append(permuted, nums[index])
	visited[index] = true
	if len(permuted) == len(nums) {
		*ret = append(*ret, permuted)
	} else {
		for i := 0; i < len(nums); i++ {
			if !visited[i] {
				backtrack46(nums, visited, i, permuted, ret)
			}
		}
	}
	visited[index] = false
	return
}
