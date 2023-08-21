/*
*

	@author: jiangxi
	@since: 2023/8/21
	@desc: //TODO

*
*/
package newadd

import "fmt"

// 根据csnote启发写的，我自己没想出来也没总结出来。
// dp[i]为第i个丑数，假设dp[i2]*2,dp[i3]*3,dp[i5]*5，dp[i2-k]*2,dp[i3-k]*3,dp[i5-k]*5都已经纳入丑数数组，那下一个丑数一定是三者其一
// 对于n，n/2,n/3,n/5有丑数的话，那它也是丑数，
func nthUglyNumber(n int) int {
	dp := make([]int, n)
	i2, i3, i5 := 0, 0, 0
	dp[0] = 1
	var n2, n3, n5 int
	for i := 1; i < n; i++ {
		n2 = dp[i2] * 2
		n3 = dp[i3] * 3
		n5 = dp[i5] * 5
		minNum := mino49(n2, n3, n5)
		dp[i] = minNum
		if n2 == minNum {
			i2++
		}
		if n3 == minNum {
			i3++
		}
		if n5 == minNum {
			i5++
		}
	}
	//fmt.Print(dp)
	return dp[n-1]

}

func mino49(a ...int) int {
	min := a[0]
	for i := range a {
		if a[i] < min {
			min = a[i]
		}
	}
	return min
}

// 这是我的想法，不对。会耗尽空间
func nthUglyNumber2(n int) int {
	dp := make([]bool, 10)
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp[1], dp[2], dp[3], dp[4], dp[5] = true, true, true, true, true
	cnt := 2
	for i := 3; cnt < n; i++ {
		if i >= len(dp) {
			dp = append(dp, make([]bool, len(dp))...) //我的想法会在分配的时候把空间耗尽
		}
		if !dp[i] {
			if i%2 == 0 {
				dp[i] = dp[i] || dp[i/2]
			}
			if i%3 == 0 {
				dp[i] = dp[i] || dp[i/3]
			}
			if i%5 == 0 {
				dp[i] = dp[i] || dp[i/5]
			}
		}
		if dp[i] {
			fmt.Println(i)
			cnt++
			if cnt == n {
				return i
			}
		}
	}
	fmt.Println(dp)
	return 1

}
