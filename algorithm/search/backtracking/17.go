/**
  @author: jiangxi
  @since: 2022/12/30
  @desc: //TODO
**/
package backtracking

var letterMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}
var combinations []string

func letterCombinations(digits string) []string {
	combinations = []string{}

	if len(digits) == 0 {
		return combinations
	}
	backtrack17(digits, 0, "")
	return combinations

}

func backtrack17(digits string, index int, combination string) {
	if len(combination) == len(digits) {
		combinations = append(combinations, combination)
		return
	}
	letters := letterMap[digits[index:index+1]]
	for i := range letters {
		backtrack17(digits, index+1, combination+letters[i:i+1])
	}
	return
}
