/**
  @author: jiangxi
  @since: 2022/12/27
  @desc: //TODO
**/
package greed

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) (ans int) {
	sort.Ints(g)
	sort.Ints(s)
	m, n := len(g), len(s)
	for i, j := 0, 0; i < m && j < n; i++ {
		for j < n && g[i] > s[j] {
			j++
		}
		if j < n {
			fmt.Println(i, j)
			ans++
			j++
		}
	}
	return ans
}
