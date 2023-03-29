/**
  @author: jiangxi
  @since: 2022/12/27
  @desc: //TODO
**/
package greed

import "sort"

//好难的思想，先放高个子，因为k值只会收到比自己高的影响
func reconstructQueue(people [][]int) (ans [][]int) {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0] || people[i][0] == people[j][0] && people[i][1] < people[j][1]
	})
	for _, person := range people {
		idx := person[1]
		ans = append(ans[:idx], append([][]int{person}, ans[idx:]...)...) //插入
	}
	return
}
