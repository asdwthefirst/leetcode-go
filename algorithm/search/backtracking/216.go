/**
  @author: jiangxi
  @since: 2022/12/30
  @desc: //TODO
**/
package backtracking

func combinationSum3(k int, n int) (ret [][]int) {
	visited := make([]bool, 9)
	for i := 1; i <= 9; i++ {
		backtrack216(k, n, i, []int{}, visited, &ret)
	}
	return

}

func backtrack216(k, n int, num int, nums []int, visited []bool, ret *[][]int) {
	visited[num-1] = true
	nums = append(nums, num)
	n -= num
	k--
	if k == 0 {
		if n == 0 {
			*ret = append(*ret, append([]int{}, nums...))
		}
	} else {
		for i := num + 1; i <= 9; i++ {
			if !visited[i-1] && i <= n {
				backtrack216(k, n, i, nums, visited, ret)
			}
		}
	}
	visited[num-1] = false
}
