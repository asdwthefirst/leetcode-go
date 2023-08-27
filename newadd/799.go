/*
*

	@author: jiangxi
	@since: 2023/8/21
	@desc: //TODO

*
*/
package newadd

import "fmt"

// 每个杯子只会收到左右两边的溢出，一个是i-1,j-1,一个是i-1,j,假设我有一个函数in(x,y)能知道x，y的总流入量，实际上有起点0，0是有初始流入量的。
func champagneTower(poured int, query_row int, query_glass int) float64 {
	fmt.Println(poured, query_row, query_glass)
	dp := make([]float64, query_glass+2)
	dp[0] = 0
	dp[1] = float64(poured)
	for i := 1; i <= query_row; i++ {
		for j := query_glass + 1; j >= 1; j-- {
			add := 0.0
			if dp[j] >= 1 {
				add = (dp[j] - 1) / 2
			}
			if dp[j-1] >= 1 {
				add += (dp[j-1] - 1) / 2
			}
			dp[j] = add
		}
	}
	return min(1.0, dp[query_glass+1])

}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
