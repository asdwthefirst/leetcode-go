/*
*

	@author: jiangxi
	@since: 2023/1/4
	@desc: //TODO

*
*/
package __1bag

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] = dp[j-coin] + dp[j]
		}
	}
	return dp[amount]
}

// 0630
func change0630(amount int, coins []int) int {

	//dp[i][j]表示用前i种钱币达到[j]的组合数
	//dp[i][j]=dp[i-1][j]+dp[i-1][j-coin[i]]+dp[i-1][j-*2coin[i]]...
	//我这种想法的j是每次得从0开始算。。。
	if amount == 0 {
		return 1
	}
	dp := make([][]int, len(coins))
	for i, _ := range dp {
		dp[i] = make([]int, amount+1)
	}

	for z := 0; coins[0]*z <= amount; z++ {
		dp[0][coins[0]*z] = 1 //处理边界问题，target=0的情况仔细思考了一下
	}

	//target=0的情况仔细思考了一下，发现就是dp[i][0]=1
	for i := 1; i < len(coins); i++ {
		for j := 0; j <= amount; j++ {
			dp[i][j] += dp[i-1][j] //target=0的情况仔细思考了一下，
			for z := 1; j-coins[i]*z >= 0; z++ {
				dp[i][j] += dp[i-1][j-coins[i]*z]
			}
		}
		// fmt.Println(i,'\n',dp)
	}
	return dp[len(coins)-1][amount]

	//dp[i][j]表示用前i种钱币达到[j]的组合数
	//dp[i][j]=dp[i-1][j-coin[i]],dp[i-1][j-*2coin[i]]...
	//但是dp[i-1][j-2*coin[i]]的情况，和dp[i][j-coin[i]]相同，而且没有被j倍数覆盖的与i-1保持一致
	// if amount==0{
	//     return 1
	// }
	// dp:=make([]int,amount+1)
	// sort.Ints(coins)
	// for i:=0;i<len(coins);i++{
	//     if coins[i]>amount{
	//         break
	//     }
	//     dp[coins[i]]+=1
	//     for j:=coins[i];j<=amount;j++{
	//         dp[j]+=dp[j-coins[i]]
	//     }
	// }

	// return dp[amount]
}
