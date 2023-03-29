/**
  @author: jiangxi
  @since: 2022/12/15
  @desc: //TODO
**/
package arraymatrix

func isToeplitzMatrix(matrix [][]int) bool {
	row, col := len(matrix), len(matrix[0])
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}

	}
	return true

}
