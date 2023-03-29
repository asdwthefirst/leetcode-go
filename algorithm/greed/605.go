/**
  @author: jiangxi
  @since: 2022/12/27
  @desc: //TODO
**/
package greed

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			if (i-1 < 0 || flowerbed[i-1] == 0) && (i+1 >= len(flowerbed) || flowerbed[i+1] == 0) {
				flowerbed[i] = 1
				n--
			}
		}
		if n <= 0 {
			break
		}
	}
	return n <= 0

}
