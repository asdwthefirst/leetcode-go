/**
  @author: jiangxi
  @since: 2022/12/27
  @desc: //TODO
**/
package sort

//桶排性能没堆好
func topKFrequent2(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	bucket := make([][]int, len(nums))
	for i := range bucket {
		bucket[i] = []int{}
	}
	for key, value := range occurrences {
		bucket[value-1] = append(bucket[value-1], key)
	}
	result := make([]int, k)

	i, j := 0, len(bucket)-1
	for i < k && j >= 0 {
		for z := len(bucket[j]) - 1; z >= 0 && i < k; z-- {
			result[i] = bucket[j][z]
			i++
		}
		j--
	}
	return result
}
