/**
  @author: jiangxi
  @since: 2022/12/29
  @desc: //TODO
**/
package DFS

var visited547 []bool
var cityNum int

func findCircleNum(isConnected [][]int) int {
	cityNum = len(isConnected)
	visited547 = make([]bool, len(isConnected))
	ans := 0
	for i := 0; i < cityNum; i++ {
		if visited547[i] == false {
			dfs547(isConnected, i)
			ans++

		}
	}
	return ans

}

func dfs547(isConnected [][]int, i int) {
	visited547[i] = true
	for j := 0; j < cityNum; j++ {
		if visited547[j] == false && isConnected[i][j] == 1 {
			dfs547(isConnected, j)
		}
	}
	return
}
