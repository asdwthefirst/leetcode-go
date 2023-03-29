/**
  @author: jiangxi
  @since: 2022/11/12
  @desc: //TODO
**/
package arraymatrix

func matrixReshape(mat [][]int, r int, c int) [][]int {

	oldR := len(mat)
	oldC := len(mat[0])
	if r*c != oldR*oldC {
		return mat
	}
	result := make([][]int, r)
	for i := range result {
		result[i] = make([]int, c)
	}

	for i := 0; i < r*c; i++ {
		oldRow := i / oldC
		oldColumn := i % oldC
		newRow := i / c
		newColumn := i % c
		result[newRow][newColumn] = mat[oldRow][oldColumn]
	}
	return result

}
