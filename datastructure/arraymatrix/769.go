/**
  @author: jiangxi
  @since: 2022/12/15
  @desc: //TODO
**/
package arraymatrix

func maxChunksToSorted(arr []int) (ans int) {
	max := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if i == max {
			ans++
		}
	}
	return ans
}
