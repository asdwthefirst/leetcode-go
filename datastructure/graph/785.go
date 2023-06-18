/*
*

	@author: jiangxi
	@since: 2022/12/16
	@desc: //TODO

*
*/
package graph

var (
	UNCOLORED, RED, GREEN = 0, 1, 2
	color                 []int
	valid                 bool
)

func isBipartite(graph [][]int) bool {
	color = make([]int, len(graph))
	valid = true
	for i := 0; i < len(graph) && valid; i++ {
		if color[i] == UNCOLORED { //先判定是否color，我0618做的时候没有考虑好
			dfs(graph, RED, i)
		}
	}
	return valid

}

func dfs(graph [][]int, c int, u int) {
	color[u] = c
	cNei := RED
	if c == RED {
		cNei = GREEN
	}
	for _, v := range graph[u] {
		if color[v] == UNCOLORED {
			dfs(graph, cNei, v)
			if !valid {
				return
			}
		} else if color[v] != cNei {
			valid = false
			return
		}
	}
	return

}
