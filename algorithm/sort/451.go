/**
  @author: jiangxi
  @since: 2022/12/27
  @desc: //TODO
**/
package sort

func frequencySort(s string) string {
	t := []byte(s) //string to char array
	newStr := []byte{}
	occurrences := make(map[byte]int)
	for _, ch := range t {
		occurrences[ch]++
	}
	bucket := make([][]byte, len(t))
	for i := range bucket {
		bucket[i] = []byte{}
	}
	for key, value := range occurrences {
		bucket[value-1] = append(bucket[value-1], key)
	}
	i, j := 0, len(bucket)-1
	for i < len(t) && j >= 0 {
		for z := len(bucket[j]) - 1; z >= 0; z-- {
			for freq := 0; freq < j+1; freq++ {
				newStr = append(newStr, bucket[j][z])
				i++
			}
		}
		j--
	}
	return string(newStr)

}
