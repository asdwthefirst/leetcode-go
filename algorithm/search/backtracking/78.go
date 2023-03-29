/**
  @author: jiangxi
  @since: 2022/12/30
  @desc: //TODO
**/
package backtracking

func subsets(nums []int) (ret [][]int) {
	visited := make([]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		backtrack78(nums, i, visited, &ret, []int{})
	}
	ret = append(ret, []int{})
	return

}

func backtrack78(nums []int, index int, visited []bool, ret *[][]int, choosed []int) {
	visited[index] = true
	choosed = append(choosed, nums[index])
	*ret = append(*ret, append([]int{}, choosed...))
	for i := index + 1; i < len(nums); i++ {
		if !visited[i] {
			backtrack78(nums, i, visited, ret, choosed)
		}
	}
	visited[index] = false
	return
}
