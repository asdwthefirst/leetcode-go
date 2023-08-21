/*
*

	@author: jiangxi
	@since: 2022/12/19
	@desc: //TODO

*
*/
package graph

//并查集

func findRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1)
	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			//祖先没有祖先
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(u, v int) bool {
		x, y := find(u), find(v)
		if x == y {
			return false
		}
		parent[x] = y //人为连接，这样做，已经包含的边连接的点只会有一个共同的祖先。
		return true
	}
	for _, e := range edges {
		if !union(e[0], e[1]) {
			return e
		}

	}
	return nil
}
