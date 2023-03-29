/**
  @author: jiangxi
  @since: 2022/11/12
  @desc: //TODO
**/
package arraymatrix

//选右上角或左下角，这样大了或小了是有确定的方向的
func searchMatrix(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] < target {
			i++
		} else if matrix[i][j] > target {
			j--
		} else {
			return true
		}
	}
	return false

}
