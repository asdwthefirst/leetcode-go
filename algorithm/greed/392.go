/**
  @author: jiangxi
  @since: 2022/12/27
  @desc: //TODO
**/
package greed

//官解双指针，我看贪心的算法感觉也就是双指针，，，
func isSubsequence(s string, t string) bool {
	n, m := len(s), len(t)
	i, j := 0, 0
	for i < n && j < m {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == n
}
