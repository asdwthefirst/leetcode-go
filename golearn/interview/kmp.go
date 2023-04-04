package interview

import "fmt"

// 在这个实现中，我们首先计算了模式串的前缀表（next 数组），
// 然后使用它进行匹配。在匹配过程中，我们维护了一个指针 j，
// 它指向模式串中已经匹配的前缀的下一个字符。如果当前字符匹配失败，
// 则我们将 j 移动到前缀表中对应的位置，以尽可能地利用已经匹配的信息。
// 如果 j 移动到了模式串的末尾，则表示匹配成功，返回匹配的起始位置。
// 如果匹配失败，则返回 -1。
func KMPSearch(text, pattern string) int {
	n, m := len(text), len(pattern)
	if m == 0 {
		return 0
	}

	// 计算前缀表
	next := make([]int, m)
	j := 0
	for i := 1; i < m; i++ {
		for j > 0 && pattern[j] != pattern[i] {
			j = next[j-1]
		}
		if pattern[j] == pattern[i] {
			j++
		}
		next[i] = j
	}

	// 使用前缀表进行匹配
	j = 0
	for i := 0; i < n; i++ {
		for j > 0 && pattern[j] != text[i] {
			j = next[j-1]
		}
		if pattern[j] == text[i] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}

	return -1
}

func main() {
	text := "ABCABCDABABCDABCDABDE"
	pattern := "ABCDABD"
	pos := KMPSearch(text, pattern)
	fmt.Println(pos) // Output: 11
}
