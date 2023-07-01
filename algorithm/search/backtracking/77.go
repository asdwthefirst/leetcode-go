/*
*

	@author: jiangxi
	@since: 2022/12/30
	@desc: //TODO

*
*/
package backtracking

//给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
//
//你可以按 任何顺序 返回答案。
//
//示例 1：
//
//输入：n = 4, k = 2
//输出：
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//]

func combine(n int, k int) (ret [][]int) {
	for i := 1; i <= n; i++ {
		visited := make([]bool, n)
		backtrack77(i, visited, &ret, []int{}, n, k)
	}
	return
}

func backtrack77(num int, visited []bool, ret *[][]int, nums []int, n, k int) {
	nums = append(nums, num)
	visited[num-1] = true
	if len(nums) == k {
		*ret = append(*ret, append([]int{}, nums...))
		//fmt.Println(nums)
		//fmt.Println(len(*ret))
	} else {
		for i := num + 1; i <= n; i++ { //组合而不是排练，那就从自己开始
			if !visited[i-1] {
				backtrack77(i, visited, ret, nums, n, k)
			}
		}
	}
	visited[num-1] = false
	return
}
