/**
  @author: jiangxi
  @since: 2022/11/11
  @desc: //TODO
**/
package stackqueue

func isValid(s string) bool {
	if s == "" {
		return true
	}
	stack := make([]string, len(s))
	pos := 0
	for i := 0; i < len(s); i++ {
		//fmt.Println(s, i, string(s[i]), stack)
		switch string(s[i]) {
		case "(", "[", "{":
			stack[pos] = string(s[i])
			//fmt.Println(stack)
			pos++
		case ")":
			pos--
			if pos < 0 || stack[pos] != "(" {
				return false
			}
		case "]":
			pos--
			if pos < 0 || stack[pos] != "[" {
				return false
			}
		case "}":
			pos--
			if pos < 0 || stack[pos] != "{" {
				return false
			}
		default:
			return false
		}
	}
	if pos == 0 {
		return true
	}
	return false
}
