/**
  @author: jiangxi
  @since: 2022/12/29
  @desc: //TODO
**/
package BFS

//经典最短路径
func ladderLength(beginWord string, endWord string, wordList []string) int {
	n := len(wordList)
	pathLen := make([]int, n+1)
	pathLen[n] = 1
	var queue []int
	queue = append(queue, n)
	for len(queue) != 0 {
		curIdx := queue[0]
		var curWord string
		if curIdx == -1 {
			curWord = beginWord
		} else {
			curWord = wordList[curIdx]
		}
		queue = queue[1:]
		if curIdx != -1 && curWord == endWord {
			return pathLen[curIdx]
		}
		for i := 0; i < n; i++ {
			if pathLen[i] == 0 && canChange(curWord, wordList[i]) {
				pathLen[i] = pathLen[curIdx] + 1
				queue = append(queue, i)
			}
		}
	}
	return 0
}

func canChange(a, b string) bool {
	if len(a) != len(b) || a == b {
		return false
	}
	diff := 1
	for i := range a {
		if a[i:i+1] != b[i:i+1] {
			if diff > 0 {
				diff--
			} else {
				return false
			}
		}
	}
	return true
}
