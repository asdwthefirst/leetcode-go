/*
*

	@author: jiangxi
	@since: 2022/12/18
	@desc: //TODO

*
*/
package graph

//求拓扑，分为未搜索，搜索中，已完成
//实际上就是出现了找环，有环就是no

var (
	valid207                        bool
	edges                           [][]int
	vis                             []int
	UNSEARCHED, SEARCHING, SEARCHED = 0, 1, 2
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	valid207 = true
	edges = make([][]int, numCourses)
	vis = make([]int, numCourses)
	for _, e := range prerequisites {
		edges[e[1]] = append(edges[e[1]], e[0])
	}
	for i := 0; i < numCourses && valid207; i++ {
		if vis[i] == UNSEARCHED {
			dfs207(i)
		}
	}
	return valid207
}

func dfs207(u int) {
	vis[u] = 1
	for _, v := range edges[u] {
		if vis[v] == UNSEARCHED {
			dfs207(v)
		} else if vis[v] == SEARCHING {
			valid207 = false
			return
		}
	}
	vis[u] = 2
	return
}
