/*
*

	@author: jiangxi
	@since: 2023/7/23
	@desc: //TODO

*
*/
package dp

import "fmt"

//https://www.acmcoder.com/ojques.html?id=5fe5ad711ea3bb5d796b0b31

var mod = 1000000007

func cutBloomsMain() {
	var n int
	var blooms []int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		blooms = append(blooms, num)
	}
	fmt.Print(cutCnt(blooms, n))

}

// dp[i]表示前i个有多少种cut法，首先它前面的位置单独cut一下是一种剪法，
// 如果它前面的位置不cut，
func cutCnt(blooms []int, n int) int {
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1]
		for j := i - 1; j >= 0 && i-j <= 9; j-- {
			if isOk(blooms, j, i) {
				if j == 0 {
					dp[i] += 1
				} else {
					dp[i] = (dp[i] + dp[j-1]) % mod
				}
			}
		}
	}
	// fmt.Println(dp)
	return dp[n-1]
}

func isOk(blooms []int, left, right int) bool {
	colorMap := make(map[int]bool)
	for i := left; i <= right; i++ {
		if _, ok := colorMap[blooms[i]]; ok {
			return false
		}
		colorMap[blooms[i]] = true
	}
	return true

}
