/*
*

	@author: jiangxi
	@since: 2023/1/5
	@desc: //TODO

*
*/
package stringedit

import "fmt"

// 其实跟583一样，但我也是真的没有意识到
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1) //预留一位不用的0，避免转移的时候判断有没有越界。。。
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 1; i <= len(word1); i++ {
		dp[i][0] = i //我自己写错了写成1，属于秀逗
	}
	for j := 1; j <= len(word2); j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min72(min72(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
		}
	}
	fmt.Println(dp)
	return dp[len(word1)][len(word2)]
}

func min72(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// if wordi==wordj,dp[i][j]=dp[i-1][j-1]
// if !=,dp[i][j]=min(dp[i-1][j]+1,dp[i][j-1]+1,dp[i-1][j-1]+1)，分别对应，删除，删除，更换
func minDistance0701(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for j := range dp {
		dp[j] = make([]int, len(word2)+1)
	}
	for i := range dp {
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}
	// fmt.Println(dp)
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min720701(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}
	// fmt.Println(dp)
	return dp[len(word1)][len(word2)]

}

func min720701(nums ...int) int {
	min := nums[0]
	for i := range nums {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}
