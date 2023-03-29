/**
  @author: jiangxi
  @since: 2022/12/3
  @desc: //TODO
**/
package arraymatrix

/*在[left, right]中符合条件(count(矩阵中小于它的数)==k)的数可能会有多个，
这些数中，最小的那个(设为a)一定在矩阵中，对于任意整数i(i<a) ，count(i)<k，
直到i等于a时，count将第一次等于k。因此二分查找找到第一个使count==k的位置，
也就是c++里lower_bound所求的位置，就一定是所求。*/

/* f（mid）=k，可以比较容易的发现，当k突变为k+1的时候的mid，肯定在矩阵中
（因为k代表的是矩阵中小于等于mid的数，mid慢慢变大，直到mid等于矩阵中的某一个数的时候，
k就会变成k+1，如果存在相同元素，也可能不是1，但是这个不影响结果）。
所以，就是当f（mid）=k时mid的左边界（最小值），当target不在k的序列中的时候，
用找k的左边界的方法找到的结果会是比target大的k中存在的数字，结果也是对的。
*/
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	left := matrix[0][0]
	right := matrix[n-1][n-1]
	for left < right {
		mid := left + (right-left)/2
		if larger378(matrix, n, mid, k) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func larger378(matrix [][]int, n, mid, k int) bool {
	i, j := n-1, 0
	num := 0
	for i >= 0 && j < n {
		if matrix[i][j] <= mid {
			j++
			num += i + 1
		} else {
			i--
		}
	}
	return num >= k
}
