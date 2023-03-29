/**
  @author: jiangxi
  @since: 2022/12/19
  @desc: //TODO
**/
package graph

func findRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1)
	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			//要一级一级往上找,知道找到总点等于自己点真总点
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(u, v int) bool {
		x, y := find(u), find(v)
		if x == y {
			return false
		}
		parent[x] = y
		return true
	}
	for _, e := range edges {
		if !union(e[0], e[1]) {
			return e
		}

	}
	return nil
}
