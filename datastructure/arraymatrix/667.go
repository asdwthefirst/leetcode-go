/*
*

	@author: jiangxi
	@since: 2022/12/15
	@desc: //TODO

*
*/
package arraymatrix

func constructArray(n, k int) []int {
	ans := make([]int, 0, n)
	//1,2,3,...,n-k-1,n-k,n,n-k+1,n+1,...
	//1,1,1,...,1,k,k-1,k-2,...,2
	for i := 1; i < n-k; i++ {
		ans = append(ans, i)
	}
	for i, j := n-k, n; i <= j; i++ {
		ans = append(ans, i)
		if i != j {
			ans = append(ans, j)
		}
		j--
	}
	return ans
}
