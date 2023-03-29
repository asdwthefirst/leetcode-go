/**
  @author: jiangxi
  @since: 2022/11/11
  @desc: //TODO
**/
package hashset

func findLHS(nums []int) int {
	countMap := make(map[int]int)
	for _, num := range nums {
		if _, ok := countMap[num]; ok {
			countMap[num] += 1
			//fmt.Println(num, countMap[num])
		} else {
			countMap[num] = 1
			//fmt.Println(num, countMap[num])
		}
	}
	max := 0
	for k, _ := range countMap {
		if left, ok := countMap[k-1]; ok {
			if left+countMap[k] > max {
				max = left + countMap[k]
			}
		}
	}
	return max
}
